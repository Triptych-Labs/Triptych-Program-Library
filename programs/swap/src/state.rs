use crate::structs::*;
use anchor_lang::prelude::*;

#[account]
pub struct SwapRecorder {
    pub proposals: u64,
}

impl SwapRecorder {
    pub const LEN: usize = 8 + 8;
}

#[account]
pub struct SwapProposal {
    pub enabled: bool,
    pub swaps: u64,
    pub swapped: u64,
    pub index: u64,
    pub per: u64,
    pub exchange: u64,
    pub mint_decimals: u8,
    pub oracle: Pubkey,
    pub from_mint: Pubkey,
    pub to_mint: Pubkey,
    pub pool: Pubkey,
    pub splits: Vec<Split>,
}

impl SwapProposal {
    pub const LEN: usize = 8
        + 2
        + 8
        + 8
        + 8
        + 8
        + 8
        + 1
        + 32
        + 32
        + 32
        + 32
        + (4 + (5 * std::mem::size_of::<Split>()));
}
