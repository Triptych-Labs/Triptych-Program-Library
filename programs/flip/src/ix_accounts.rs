use crate::constants::*;
use crate::program::Flipper;
use crate::state::*;
use anchor_lang::prelude::*;
use escrow;
use escrow::program::Escrow;

#[derive(Accounts)]
pub struct CreateFlip<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        init,
        seeds = [HOUSE.as_ref(), oracle.key().as_ref()],
        bump,
        payer = oracle,
        space = House::LEN
    )]
    pub house: Box<Account<'info, House>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(house_bump: u8)]
pub struct WithdrawHouse<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        mut,
        seeds = [HOUSE.as_ref(), oracle.key().as_ref()],
        bump = house_bump,
    )]
    pub house: Box<Account<'info, House>>,
}

#[derive(Accounts)]
#[instruction(house_bump: u8, flip_bump: u8, escrow_bump: u8, daily_epoch: u64)]
pub struct NewFlip<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    /// CHECK: checked in ix
    pub fees: AccountInfo<'info>,
    #[account(
        mut,
        seeds = [escrow::constants::ESCROW.as_ref(), initializer.key().as_ref()],
        bump = escrow_bump,
        seeds::program = escrow_program.key(),
    )]
    pub escrow: Box<Account<'info, escrow::state::Escrow>>,
    pub escrow_program: Program<'info, Escrow>,
    #[account(mut)]
    /// CHECK: am lazy
    pub oracle: AccountInfo<'info>,
    #[account(
        mut,
        seeds = [HOUSE.as_ref(), oracle.key().as_ref()],
        bump = house_bump,
    )]
    pub house: Box<Account<'info, House>>,
    #[account(
        init_if_needed,
        seeds = [FLIP.as_ref(), oracle.key().as_ref(), daily_epoch.to_le_bytes().as_ref()],
        bump,
        payer = initializer,
        space = Flip::LEN,
    )]
    pub flip: Box<Account<'info, Flip>>,
    pub system_program: Program<'info, System>,
    pub flipper_program: Program<'info, Flipper>,
    /// CHECK: am lazy
    pub slot_hashes: UncheckedAccount<'info>,
}
