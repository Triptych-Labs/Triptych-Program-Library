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
pub struct Flip {
    pub initialized: Option<bool>,
    pub oracle: Pubkey,
    pub daily_epoch: u64,
    pub heads: [u64; 3],
    pub tails: [u64; 3],
}

impl Flip {
    pub const LEN: usize = 8 + (1 + 1) + 32 + 64 + (2 * (4 * (8 * 3)));
}

#[account]
pub struct Escrow {
    pub initializer: Pubkey,
    pub available_balance: u64,
}

impl Escrow {
    pub const LEN: usize = 8 + 32 + 8;
}
