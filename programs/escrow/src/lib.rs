use anchor_lang::prelude::*;
use errors::FlipError;
use ix_accounts::*;
use std::result::Result;

declare_id!("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ");

pub mod constants;
mod errors;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod escrow {
    use solana_program::{program::invoke, system_instruction::transfer};

    use crate::constants::AUTHORIZED_PROGRAMS;

    use super::*;

    pub fn initialize_escrow(ctx: Context<InitializeEscrow>) -> Result<(), Error> {
        let escrow = &mut ctx.accounts.escrow;
        escrow.initializer = ctx.accounts.initializer.key();
        escrow.available_balance = 0;

        Ok(())
    }

    pub fn deposit_escrow(
        ctx: Context<DepositEscrow>,
        _escrow_bump: u8,
        amount: u64,
    ) -> Result<(), Error> {
        let initializer = &mut ctx.accounts.initializer;
        let escrow = &mut ctx.accounts.escrow;
        let system_program = &mut ctx.accounts.system_program;
        escrow.initializer = initializer.key();
        escrow.available_balance += amount;

        invoke(
            &transfer(&initializer.key(), &escrow.key(), amount),
            &[
                initializer.to_account_info(),
                escrow.to_account_info(),
                system_program.to_account_info(),
            ],
        )?;

        Ok(())
    }

    pub fn install_escrow(
        ctx: Context<InstallEscrow>,
        _escrow_bump: u8,
        amount: u64,
    ) -> Result<(), Error> {
        let initializer = &mut ctx.accounts.initializer;
        let installer = &mut ctx.accounts.installer;
        let escrow = &mut ctx.accounts.escrow;

        // let installer_info = installer.to_account_info();

        escrow.initializer = initializer.key();
        escrow.available_balance += amount;
        /*
        **installer_info.try_borrow_mut_lamports()? =
            installer_info.lamports().checked_sub(amount).unwrap();
        */

        Ok(())
    }

    pub fn drain_escrow(
        ctx: Context<DrainEscrow>,
        _escrow_bump: u8,
        amount: u64,
    ) -> Result<(), Error> {
        let initializer = &mut ctx.accounts.initializer;
        let collector = &mut ctx.accounts.collector;
        let escrow = &mut ctx.accounts.escrow;
        let caller_program = &mut ctx.accounts.caller_program;

        let caller_program_id = caller_program.key().to_string();

        let is_trusted_program = AUTHORIZED_PROGRAMS
            .iter()
            .any(|&authorized_program| authorized_program == caller_program_id.as_str());

        if !is_trusted_program {
            if caller_program_id.as_str() != "11111111111111111111111111111111" {
                return Err(FlipError::SuspiciousTransaction.into());
            }
            if initializer.key().to_string() != collector.key().to_string() {
                return Err(FlipError::SuspiciousTransaction.into());
            }
        }

        let collector_info = collector.to_account_info();
        let escrow_info = escrow.to_account_info();

        **collector_info.try_borrow_mut_lamports()? =
            collector_info.lamports().checked_add(amount).unwrap();
        **escrow_info.try_borrow_mut_lamports()? =
            escrow_info.lamports().checked_sub(amount).unwrap();

        escrow.available_balance -= amount;

        Ok(())
    }
}
