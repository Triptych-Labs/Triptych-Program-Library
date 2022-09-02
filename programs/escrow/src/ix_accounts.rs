use crate::constants::*;
use crate::state::*;
use anchor_lang::prelude::*;

#[derive(Accounts)]
pub struct InitializeEscrow<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        init,
        seeds = [ESCROW.as_ref(), initializer.key().as_ref()],
        bump,
        payer = initializer,
        space = Escrow::LEN
    )]
    pub escrow: Box<Account<'info, Escrow>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(escrow_bump: u8)]
pub struct DepositEscrow<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        mut,
        seeds = [ESCROW.as_ref(), initializer.key().as_ref()],
        bump = escrow_bump,
    )]
    pub escrow: Box<Account<'info, Escrow>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(escrow_bump: u8)]
pub struct InstallEscrow<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    pub installer: Signer<'info>,
    #[account(
        mut,
        seeds = [ESCROW.as_ref(), initializer.key().as_ref()],
        bump = escrow_bump,
    )]
    pub escrow: Box<Account<'info, Escrow>>,
}

#[derive(Accounts)]
#[instruction(escrow_bump: u8)]
pub struct DrainEscrow<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    pub collector: Signer<'info>,
    #[account(
        mut,
        seeds = [ESCROW.as_ref(), initializer.key().as_ref()],
        bump = escrow_bump,
    )]
    pub escrow: Box<Account<'info, Escrow>>,
    /// CHECK: am lazy
    pub caller_program: UncheckedAccount<'info>,
}
