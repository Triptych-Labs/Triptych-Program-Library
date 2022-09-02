use anchor_lang::prelude::*;
// use crate::structs::*;

#[account]
pub struct House {
    pub oracle: Pubkey,
    pub payed_out: u64,
    pub collected: u64,
}

impl House {
    pub const LEN: usize = 8 + 32 + 32 + 8 + 8;
}

#[account]
pub struct Games {
    pub initializer: Pubkey,
    pub games: u64,
}

impl Games {
    pub const LEN: usize = 8 + 32 + 8;
}

#[account]
pub struct Game {
    pub index: u64,
    pub initialized: Option<bool>,
    pub player: Pubkey,
    pub bet_amount: u64,
    pub daily_epoch: u64,
    // upto 5 cards per hand
    pub hands: [[u8; 5]; 2],
    pub player_busted: bool,
    pub dealer_busted: bool,
    pub terminated: bool,
}

impl Game {
    pub const LEN: usize = 8 + (1 + 1) + 32 + 64 + (1 * (4 * (8 * 4))) + 10;
}

#[account]
pub struct Stats {
    pub initialized: Option<bool>,
    pub oracle: Pubkey,
    pub daily_epoch: u64,
    // games: [ wins, losses, played_volume, won_volume ]
    pub games: [u64; 4],
}

impl Stats {
    pub const LEN: usize = 8 + (1 + 1) + 32 + 64 + (1 * (4 * (8 * 4)));
}
