use crate::constants::*;
use crate::state::*;
use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::{Mint, Token, TokenAccount};

#[derive(Accounts)]
pub struct RegisterSwapRecorder<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        init,
        seeds = [SWAP_RECORDER.as_ref(), oracle.key().as_ref()],
        bump,
        payer = oracle,
        space = SwapRecorder::LEN
    )]
    pub swap_recorder: Box<Account<'info, SwapRecorder>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct ProposeSwap<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        init_if_needed,
        seeds = [SWAP_PDA_SEED.as_ref(), to_mint.key().as_ref(), swap_recorder.key().as_ref()],
        bump,
        payer = oracle,
        token::mint = to_mint,
        token::authority = swap_recorder
    )]
    pub swap_pool: Account<'info, TokenAccount>,
    #[account(
        init,
        seeds = [oracle.key().as_ref(), swap_recorder.proposals.to_le_bytes().as_ref()],
        bump,
        payer = oracle,
        space = SwapProposal::LEN
    )]
    pub swap: Box<Account<'info, SwapProposal>>,
    #[account(mut)]
    pub from_mint: Account<'info, Mint>,
    #[account(mut)]
    pub to_mint: Account<'info, Mint>,
    #[account(mut)]
    pub swap_recorder: Box<Account<'info, SwapRecorder>>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
}

#[derive(Accounts)]
#[instruction(swap_bump: u8)]
pub struct UpdateSwap<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        mut,
        seeds = [oracle.key().as_ref(), swap_recorder.proposals.to_le_bytes().as_ref()],
        bump = swap_bump,
    )]
    pub swap: Box<Account<'info, SwapProposal>>,
    #[account(mut)]
    pub swap_recorder: Box<Account<'info, SwapRecorder>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(swap_bump: u8, swap_index: u64, swap_recorder_bump: u8)]
pub struct InvokeSwap<'info> {
    /// CHECK: checked later
    pub oracle: UncheckedAccount<'info>,
    #[account(
        mut,
        seeds = [oracle.key().as_ref(), swap_index.to_le_bytes().as_ref()],
        bump = swap_bump,
    )]
    pub swap: Box<Account<'info, SwapProposal>>,
    #[account(mut)]
    pub swap_pool: Box<Account<'info, TokenAccount>>,
    #[account(
        mut,
        seeds = [SWAP_RECORDER.as_ref(), oracle.key().as_ref()],
        bump = swap_recorder_bump,
    )]
    pub swap_recorder: Box<Account<'info, SwapRecorder>>,
    #[account(mut)]
    pub from_mint: Account<'info, Mint>,
    #[account(mut)]
    pub from_token_account: Account<'info, TokenAccount>,
    #[account(mut)]
    pub to_mint: Account<'info, Mint>,
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    /// CHECK: checked as ata in ix
    pub to_token_account: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
    pub associated_token_program: Program<'info, AssociatedToken>,
}
