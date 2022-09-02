use crate::structs::*;
use anchor_lang::prelude::*;

#[account]
pub struct QuestAccount {
    pub quest: Pubkey,
    pub index: u64,
    pub start_time: i64,
    pub end_time: i64,
    pub initializer: Pubkey,
    pub completed: Option<bool>,
    pub last_claim: i64,
}

impl QuestAccount {
    pub const LEN: usize = 8 + 32 + 8 + 8 + 8 + 32 + 2 + 8;
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
pub struct Quest {
    pub enabled: bool,
    pub index: u64,
    pub name: String,
    pub duration: i64,
    pub oracle: Pubkey,
    pub required_level: u64,
    pub required_xp: u64,
    pub wl_candy_machines: Vec<Pubkey>,
    pub rewards: Vec<Reward>,
    pub tender: Option<Tender>,
    pub tender_splits: Option<Vec<Split>>,
    pub xp: u64,
    pub staking_config: Option<StakingConfig>,
    pub pairs_config: Option<PairsConfig>,
}

impl Quest {
    pub fn space(rewards: usize) -> usize {
        8 + 2
            + 8
            + (4 + 32)
            + 8
            + 32
            + 8
            + 8
            + (4 + (10 * 32))
            + (4 + (rewards * std::mem::size_of::<Reward>()))
            + 32
            + (std::mem::size_of::<Tender>())
            + (4 + (10 * std::mem::size_of::<Split>()))
            + 8
            + (std::mem::size_of::<StakingConfig>())
            + (std::mem::size_of::<PairsConfig>())
    }
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

#[account]
pub struct QuestRecorder {
    pub proposals: u64,
    pub quest: Pubkey,
    pub initializer: Pubkey,
    pub oracle: Pubkey,
    pub staked: Vec<Pubkey>,
}

impl QuestRecorder {
    pub fn space(staked: usize) -> usize {
        8 + 8 + (4 + (staked * 32)) + 32 + 32 + 32
    }
}

#[account]
pub struct QuestProposal {
    pub index: u64,
    pub fulfilled: bool,
    pub started: bool,
    pub finished: bool,
    pub withdrawn: bool,
    pub depositing_left: Vec<Pubkey>,
    pub depositing_right: Vec<Pubkey>,
    pub record_left: Vec<bool>,
    pub record_right: Vec<bool>,
}

impl QuestProposal {
    pub fn space(left: u8, right: u8) -> usize {
        8 + 8
            + 2
            + 2
            + 2
            + 2
            + ((4 + (32 * (left as usize))) + (4 + (32 * (right as usize))))
            + ((4 + (2 * (left as usize))) + (4 + (2 * (right as usize))))
    }
}
