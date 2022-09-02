use anchor_lang::prelude::*;
use constants::*;
use errors::FlipError;
use ix_accounts::*;
use std::result::Result;
use std::str::FromStr;

declare_id!("DbceHs8B185bS1tgnXgeKUNFFEgrXw5DaZG8R1cxJYGf");

mod constants;
mod errors;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod flipper {
    use escrow::cpi::accounts::{DrainEscrow, InstallEscrow};
    use solana_program::{program::invoke, system_instruction::transfer};

    use super::*;

    pub fn create_flip(ctx: Context<CreateFlip>) -> Result<(), Error> {
        let house = &mut ctx.accounts.house;
        let oracle = &ctx.accounts.oracle;

        house.oracle = oracle.key();
        house.payed_out = 0;
        house.collected = 0;

        Ok(())
    }

    pub fn withdraw_house(
        ctx: Context<WithdrawHouse>,
        _house_bump: u8,
        amount: u64,
    ) -> Result<(), Error> {
        let house = &mut ctx.accounts.house;
        let oracle = &mut ctx.accounts.oracle;

        let house_info = house.to_account_info();
        let oracle_info = oracle.to_account_info();

        **house_info.try_borrow_mut_lamports()? =
            house_info.lamports().checked_sub(amount).unwrap();

        **oracle_info.try_borrow_mut_lamports()? =
            oracle_info.lamports().checked_add(amount).unwrap();

        Ok(())
    }

    pub fn new_flip(
        ctx: Context<NewFlip>,
        house_bump: u8,
        _flip_bump: u8,
        escrow_bump: u8,
        daily_epoch: u64,
        operator: String,
        amount: u64,
        selection: u8,
    ) -> Result<(), Error> {
        // only permit 0,1 as selection
        if selection > 1 {
            return Err(FlipError::SuspiciousTransaction.into());
        }

        /*
           TODO: VALIDATE OPERATOR STRING
        */

        let fees = &mut ctx.accounts.fees;
        let house = &mut ctx.accounts.house;
        let flip = &mut ctx.accounts.flip;
        let escrow = &mut ctx.accounts.escrow;
        let initializer = &mut ctx.accounts.initializer;
        let escrow_program = &mut ctx.accounts.escrow_program;
        let system_program = &mut ctx.accounts.system_program;
        let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();
        let most_recent = &recent_slot_hash[13..25];
        let rng: u8 = (u64::from_le_bytes([
            most_recent[0],
            most_recent[1],
            most_recent[2],
            most_recent[3],
            most_recent[4],
            most_recent[5],
            most_recent[6],
            most_recent[7],
        ]) % 100) as u8;

        let escrow_info = escrow.to_account_info();
        let house_info = house.to_account_info();
        let house_bump_bytes = house_bump.to_le_bytes();

        msg!("is initialized {:?}", flip.initialized);
        if flip.initialized == None {
            msg!("initializing");

            /*
             transfer 2519520 lamports into escrow as means of refunding
             the gas spent to create the daily flip account
            */
            let refund: u64 = 2519520;
            /*
            **escrow_info.try_borrow_mut_lamports()? =
                escrow_info.lamports().checked_add(refund).unwrap();
            **house_info.try_borrow_mut_lamports()? =
                house_info.lamports().checked_sub(refund).unwrap();
            */

            flip.initialized = Some(true);
            flip.oracle = ctx.accounts.oracle.key();
            flip.daily_epoch = daily_epoch;
            flip.heads = [0, 0, 0];
            flip.tails = [0, 0, 0];
        }

        match operator.as_str() {
            "escrow" => {
                let house_seeds = &[
                    HOUSE.as_ref(),
                    house.oracle.as_ref(),
                    house_bump_bytes.as_ref(),
                ];
                escrow::cpi::drain_escrow(
                    CpiContext::new_with_signer(
                        escrow_program.to_account_info(),
                        DrainEscrow {
                            initializer: initializer.to_account_info(),
                            collector: house.to_account_info(),
                            escrow: escrow.to_account_info(),
                            caller_program: ctx.accounts.flipper_program.to_account_info(),
                        },
                        &[&house_seeds[..]],
                    ),
                    escrow_bump,
                    amount,
                )
                .unwrap();
            }
            "wallet" => {
                invoke(
                    &transfer(&initializer.key(), &house.key(), amount),
                    &[
                        initializer.to_account_info(),
                        house.to_account_info(),
                        system_program.to_account_info(),
                    ],
                )?;
            }
            _ => return Err(FlipError::SuspiciousTransaction.into()),
        }

        /*
            Apply fee transfer
        */
        let fees_address =
            Pubkey::from_str("G8UyAzcphHUE4ZCtH8YQCFATehWEhLN53Pf1aWi2SPCM").unwrap();
        if fees_address != fees.key() {
            return Err(FlipError::SuspiciousTransaction.into());
        }

        let fee = (amount as f64 * 0.0) as u64;

        invoke(
            &transfer(&initializer.key(), &fees.key(), fee),
            &[
                initializer.to_account_info(),
                fees.to_account_info(),
                system_program.to_account_info(),
            ],
        )?;

        let mut lower: u8 = 0;
        let mut upper: u8 = 0;
        if selection == 0 {
            flip.heads[0] += 1;
            lower = 0;
            upper = 49 - 5;
        }
        if selection == 1 {
            flip.tails[0] += 1;
            lower = 50 + 5;
            upper = 99;
        }

        // (number_hash % 2)≈rng is equal to (wagered) selection
        msg!("RNG: {}", rng);
        match lower <= rng && rng <= upper {
            true => {
                msg!("---DECISION: {}", true);
                house.payed_out += amount;

                if selection == 0 {
                    flip.heads[1] += amount;
                    flip.heads[2] += 1;
                }
                if selection == 1 {
                    flip.tails[1] += amount;
                    flip.tails[2] += 1;
                }

                **house_info.try_borrow_mut_lamports()? =
                    house_info.lamports().checked_sub(amount * 2).unwrap();

                **escrow_info.try_borrow_mut_lamports()? =
                    escrow_info.lamports().checked_add(amount * 2).unwrap();

                let house_seeds = &[
                    HOUSE.as_ref(),
                    house.oracle.as_ref(),
                    house_bump_bytes.as_ref(),
                ];

                escrow::cpi::install_escrow(
                    CpiContext::new_with_signer(
                        escrow_program.to_account_info(),
                        InstallEscrow {
                            initializer: initializer.to_account_info(),
                            installer: house.to_account_info(),
                            escrow: escrow.to_account_info(),
                        },
                        &[&house_seeds[..]],
                    ),
                    escrow_bump,
                    amount * 2,
                )
                .unwrap();
            }
            false => {
                msg!("---DECISION: {}", false);
                house.collected += amount;
            }
        }

        Ok(())
    }
}
