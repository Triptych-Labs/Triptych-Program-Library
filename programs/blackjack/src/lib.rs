use anchor_lang::prelude::*;
use constants::*;
use errors::FlipError;
use helper_fns::*;
use ix_accounts::*;
use std::result::Result;
use std::str::FromStr;

declare_id!("4D3g6DHPDiE3gD9G2yfRnEehkQKJwfh5VkExDgeHvhxr");

mod constants;
mod errors;
mod helper_fns;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod blackjack {
    use escrow::cpi::accounts::{DrainEscrow, InstallEscrow};
    use solana_program::{program::invoke, system_instruction::transfer};

    use super::*;

    pub fn create_blackjack(ctx: Context<CreateBlackjack>) -> Result<(), Error> {
        let house = &mut ctx.accounts.house;
        let oracle = &ctx.accounts.oracle;

        house.oracle = oracle.key();
        house.payed_out = 0;
        house.collected = 0;

        Ok(())
    }

    pub fn register_player(ctx: Context<RegisterPlayer>) -> Result<(), Error> {
        let games = &mut ctx.accounts.games;
        let initializer = &ctx.accounts.initializer;

        games.initializer = initializer.key();
        games.games = 0;

        Ok(())
    }

    pub fn start_game(
        ctx: Context<NewGame>,
        house_bump: u8,
        escrow_bump: u8,
        daily_epoch: u64,
        wallet: String,
        amount: u64,
    ) -> Result<(), Error> {
        // this fn does the following:
        //     place bet
        //     compute 3 card values

        let house = &mut ctx.accounts.house;
        let stats = &mut ctx.accounts.stats;
        let game = &mut ctx.accounts.game;
        let games = &mut ctx.accounts.games;
        let escrow = &mut ctx.accounts.escrow;
        let initializer = &mut ctx.accounts.initializer;
        let escrow_program = &mut ctx.accounts.escrow_program;
        let system_program = &mut ctx.accounts.system_program;
        let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();

        let house_bump_bytes = house_bump.to_le_bytes();

        // msg!("is initialized {:?}", flip.initialized);
        if stats.initialized == None {
            msg!("initializing");
            stats.initialized = Some(true);
        }

        game.index = games.games;
        game.bet_amount = amount;
        game.initialized = Some(true);
        game.daily_epoch = daily_epoch;
        game.player = initializer.key();
        game.hands = [[0; 5], [0; 5]];
        game.terminated = false;

        games.games += 1;

        // accept bet
        match wallet.as_str() {
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
                            caller_program: ctx.accounts.blackjack_program.to_account_info(),
                        },
                        &[&house_seeds[..]],
                    ),
                    escrow_bump,
                    amount,
                )
                .unwrap();
            }
            _ => return Err(FlipError::SuspiciousTransaction.into()),
        }

        // (number_hash % 2)≈rng is equal to (wagered) selection
        for i in 0..3 {
            // access some bytes every (i * 32) starting with +13 offset.
            let most_recent = &recent_slot_hash[((i * 32) + 13)..((i * 32) + 25)];
            let rng: u8 = (u64::from_le_bytes([
                most_recent[0],
                most_recent[1],
                most_recent[2],
                most_recent[3],
                most_recent[4],
                most_recent[5],
                most_recent[6],
                most_recent[7],
            ]) % 10) as u8
                + 1; // 0-9 becomes 1-10
            msg!("RNG: {}", rng);
            game.hands[i / 2][i % 2] = rng;
        }

        Ok(())
    }

    pub fn player_turn(
        ctx: Context<AdvanceGame>,
        house_bump: u8,
        _stats_bump: u8,
        escrow_bump: u8,
        _daily_epoch: u64,
        _game_bump: u8,
        _game_index: u64,
        operator: String,
    ) -> Result<(), Error> {
        let game = &mut ctx.accounts.game;
        let house = &mut ctx.accounts.house;
        let escrow = &mut ctx.accounts.escrow;
        let initializer = &mut ctx.accounts.initializer;
        let escrow_program = &mut ctx.accounts.escrow_program;
        let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();

        if game.terminated {
            return Err(FlipError::SuspiciousTransaction.into());
        }

        match operator.as_str() {
            "hit" => {
                let busted = attenuate_player(&mut game.hands[0], &recent_slot_hash);
                game.player_busted = busted;
            }
            "stand" => {
                let busted = attenuate_dealer(&mut game.hands[1], &recent_slot_hash);
                game.dealer_busted = busted;
            }
            "surrender" => {
                declare_dealer_win(game)?;
                return Ok(());
            }
            _ => return Err(FlipError::SuspiciousTransaction.into()),
        }

        if operator.as_str() == "hit" {
            if game.player_busted {
                declare_dealer_win(game)?;
            }
            return Ok(());
        }
        if operator.as_str() == "surrender" {
            declare_dealer_win(game)?;
            return Ok(());
        }
        if !game.player_busted && game.dealer_busted {
            let initializer_info = initializer.to_account_info();
            let house_info = house.to_account_info();
            let house_bump_bytes = house_bump.to_le_bytes();
            let escrow_info = escrow.to_account_info();
            let escrow_program_info = escrow_program.to_account_info();
            declare_player_win(
                initializer_info,
                house_info,
                house_bump_bytes,
                escrow_info,
                escrow_bump,
                escrow_program_info,
                house.oracle,
                game,
            )?;

            return Ok(());
        }
        if game.player_busted && !game.dealer_busted {
            declare_dealer_win(game)?;
            return Ok(());
        }
        if game.player_busted == false && game.player_busted == game.dealer_busted {
            let player_hand_sum: u8 = game.hands[0].iter().sum();
            let dealer_hand_sum: u8 = game.hands[1].iter().sum();
            if player_hand_sum > dealer_hand_sum {
                let initializer_info = initializer.to_account_info();
                let house_info = house.to_account_info();
                let house_bump_bytes = house_bump.to_le_bytes();
                let escrow_info = escrow.to_account_info();
                let escrow_program_info = escrow_program.to_account_info();
                declare_player_win(
                    initializer_info,
                    house_info,
                    house_bump_bytes,
                    escrow_info,
                    escrow_bump,
                    escrow_program_info,
                    house.oracle,
                    game,
                )?;
                return Ok(());
            } else {
                declare_dealer_win(game)?;
                return Ok(());
            }
        }

        Ok(())
    }
}
