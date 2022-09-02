use anchor_lang::prelude::*;
use anchor_spl::associated_token::{self, Create};
use anchor_spl::token::{self, Burn, Transfer};
use constants::*;
use errors::SwapError;
use ix_accounts::*;
use std::result::Result;
use structs::*;

declare_id!("EocXeJ7KvMZD9k1mVWVHP6KWpRhsLP666wtahSBEMawr");

mod constants;
mod errors;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod swapper {
    use spl_token::{amount_to_ui_amount, ui_amount_to_amount};

    use super::*;

    pub fn register_swap_recorder(ctx: Context<RegisterSwapRecorder>) -> Result<(), Error> {
        let swap_recorder = &mut ctx.accounts.swap_recorder;
        swap_recorder.proposals = 0;

        Ok(())
    }

    pub fn propose_swap_record(
        ctx: Context<ProposeSwap>,
        per: u64,
        exchange: u64,
        splits: Vec<Split>,
    ) -> Result<(), Error> {
        let swap = &mut ctx.accounts.swap;
        let swaps = &mut ctx.accounts.swap_recorder;

        let oracle = &ctx.accounts.oracle;
        let to_mint = &ctx.accounts.to_mint;
        let from_mint = &ctx.accounts.from_mint;
        let pool = &ctx.accounts.swap_pool;

        swap.index = swaps.proposals;
        swap.per = per;
        swap.exchange = exchange;
        swap.mint_decimals = to_mint.decimals;
        swap.oracle = oracle.key();
        swap.from_mint = from_mint.key();
        swap.to_mint = to_mint.key();
        swap.pool = pool.key();
        swap.splits = splits;

        swaps.proposals += 1;

        Ok(())
    }

    pub fn update_swap_record(
        ctx: Context<UpdateSwap>,
        enabled: bool,
        per: u64,
        exchange: u64,
        splits: Vec<Split>,
    ) -> Result<(), Error> {
        let swap = &mut ctx.accounts.swap;
        swap.enabled = enabled;
        swap.per = per;
        swap.exchange = exchange;
        swap.splits = splits;

        Ok(())
    }

    pub fn invoke_swap(
        ctx: Context<InvokeSwap>,
        _swap_bump: u8,
        _swap_index: u64,
        swap_recorder_bump: u8,
        amount: u64,
    ) -> Result<(), Error> {
        let oracle = &ctx.accounts.oracle;
        let to_token_account = &ctx.accounts.to_token_account;
        let from_mint = &ctx.accounts.from_mint;
        let from_token_account = &ctx.accounts.from_token_account;
        let to_mint = &ctx.accounts.to_mint;
        let swap = &mut ctx.accounts.swap;
        let pool = &ctx.accounts.swap_pool;
        let rent = &ctx.accounts.rent;
        let initializer = &ctx.accounts.initializer;

        if to_token_account.data.borrow().len() == 0 {
            let _may_fail = associated_token::create(CpiContext::new(
                ctx.accounts.associated_token_program.to_account_info(),
                Create {
                    payer: initializer.to_account_info(),
                    associated_token: to_token_account.to_account_info(),
                    authority: initializer.to_account_info(),
                    mint: to_mint.to_account_info(),
                    system_program: ctx.accounts.system_program.to_account_info(),
                    token_program: ctx.accounts.token_program.to_account_info(),
                    rent: ctx.accounts.rent.to_account_info(),
                },
            ));
        }

        let oracle_key = oracle.key();
        let swap_recorder_bump_bytes = swap_recorder_bump.to_le_bytes();
        let swap_recorder_authority_seeds = &[
            SWAP_RECORDER.as_ref(),
            oracle_key.as_ref(),
            swap_recorder_bump_bytes.as_ref(),
        ];
        let signers = &[&swap_recorder_authority_seeds[..]];

        let alpha = swap.exchange as f64 / swap.per as f64;
        let from_amount = amount; // * ui_amount_to_amount(alpha, from_mint.decimals);
        let to_amount = ui_amount_to_amount(
            amount_to_ui_amount(amount, from_mint.decimals) * alpha,
            swap.mint_decimals,
        );
        msg!("{} {} {} {}", amount, alpha, from_amount, to_amount,);

        token::burn(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                Burn {
                    mint: from_mint.to_account_info(),
                    from: from_token_account.to_account_info(),
                    authority: initializer.to_account_info(),
                },
            ),
            from_amount,
        )?;

        token::transfer(
            CpiContext::new_with_signer(
                ctx.accounts.token_program.to_account_info(),
                Transfer {
                    from: pool.to_account_info(),
                    to: to_token_account.to_account_info(),
                    authority: ctx.accounts.swap_recorder.to_account_info(),
                },
                signers,
            ),
            to_amount,
        )?;

        swap.swaps += 1;
        swap.swapped += to_amount;

        Ok(())
    }
}
