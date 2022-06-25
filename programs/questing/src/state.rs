use crate::structs::*;
use anchor_lang::prelude::*;

#[account]
pub struct QuestAccount {
    pub index: u64,
    pub start_time: i64,
    pub end_time: i64,
    pub deposit_token_mint: Pubkey,
    pub initializer: Pubkey,
    pub completed: Option<bool>,
}

impl QuestAccount {
    pub const LEN: usize = 8 + 8 + 8 + 8 + 32 + 32 + 2;
}

#[account]
pub struct Quests {
    pub quests: u64,
    pub oracle: Pubkey,
    pub rewards: Vec<Reward>,
}
impl Quests {
    pub const LEN: usize = 8 + 8 + 32 + (4 + (10 * (std::mem::size_of::<Reward>() + 32)));
}

#[account]
pub struct RewardToken {
    pub mint_address: Pubkey,
    pub threshold: u8,
    // decimals are hardcoded to 1 in ./ix_accounts.rs:InitializeRewardToken
    pub amount: u64,
}

impl RewardToken {
    pub const LEN: usize = 8 + 32 + 1 + 8;
}

#[account]
pub struct Quest {
    pub enabled: bool,
    pub index: u64,
    pub name: String,
    pub duration: i64,
    pub oracle: Pubkey,
    pub required_level: u64,
    pub required_xp: u64,
    pub wl_candy_machines: Vec<Pubkey>,
    pub entitlement: Option<Reward>,
    pub tender: Option<Tender>,
    pub tender_splits: Option<Vec<Split>>,
    pub xp: u64,
}

impl Quest {
    pub const LEN: usize = 8
        + 2
        + 8
        + 32
        + 8
        + 32
        + 8
        + 8
        + (4 + (10 * 32))
        + std::mem::size_of::<Reward>()
        + 32
        + (std::mem::size_of::<Tender>())
        + (4 + (10 * std::mem::size_of::<Split>()))
        + 8;
}

#[account]
pub struct Questor {
    pub initializer: Pubkey,
    pub xp: u64,
}

impl Questor {
    pub const LEN: usize = 8 + 32 + 8;
}

#[account]
pub struct Questee {
    pub owner: Pubkey,
    pub pixelballz_mint: Pubkey,
    pub quests: u64,
    pub xp: u64,
}

impl Questee {
    pub const LEN: usize = 8 + 32 + 32 + 8 + 8;
}

#[account]
pub struct QuestQuesteeEndReceipt {
    pub owner: Pubkey,
    pub pixelballz_mint: Pubkey,
    pub reward_mint: Pubkey,
    pub amount: u64,
}

impl QuestQuesteeEndReceipt {
    pub const LEN: usize = 8 + 32 + 32 + 32 + 8;
}
