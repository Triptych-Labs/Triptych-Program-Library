use crate::constants::*;
use crate::program::Blackjack;
use crate::state::*;
use anchor_lang::prelude::*;
use escrow;
use escrow::program::Escrow;

#[derive(Accounts)]
pub struct CreateBlackjack<'info> {
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
pub struct RegisterPlayer<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        init,
        seeds = [GAME.as_ref(), initializer.key().as_ref()],
        bump,
        payer = initializer,
        space = Games::LEN
    )]
    pub games: Box<Account<'info, Games>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(house_bump: u8, escrow_bump: u8, daily_epoch: u64)]
pub struct NewGame<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut, has_one = initializer)]
    pub games: Box<Account<'info, Games>>,
    #[account(
        init,
        seeds = [GAME.as_ref(), initializer.key().as_ref(), games.games.to_le_bytes().as_ref()],
        bump,
        payer = initializer,
        space = Game::LEN
    )]
    pub game: Box<Account<'info, Game>>,
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
        space = Stats::LEN,
    )]
    pub stats: Box<Account<'info, Stats>>,
    pub system_program: Program<'info, System>,
    pub blackjack_program: Program<'info, Blackjack>,
    /// CHECK: am lazy
    pub slot_hashes: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(house_bump: u8, stats_bump: u8, escrow_bump: u8, daily_epoch: u64, game_bump: u8, game_index: u64)]
pub struct AdvanceGame<'info> {
    #[account(
        mut,
        seeds = [GAME.as_ref(), oracle.key().as_ref(), game_index.to_le_bytes().as_ref()],
        bump = game_bump,
    )]
    pub game: Box<Account<'info, Game>>,
    #[account(mut)]
    pub initializer: Signer<'info>,
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
        mut,
        seeds = [FLIP.as_ref(), oracle.key().as_ref(), daily_epoch.to_le_bytes().as_ref()],
        bump = stats_bump,
    )]
    pub stats: Box<Account<'info, Stats>>,
    pub system_program: Program<'info, System>,
    pub blackjack_program: Program<'info, Blackjack>,
    /// CHECK: am lazy
    pub slot_hashes: UncheckedAccount<'info>,
}
