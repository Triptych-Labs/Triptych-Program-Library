use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct PairsConfig {
    pub left: u8,
    pub left_creators: [Pubkey; 5],
    pub right: u8,
    pub right_creators: [Pubkey; 5],
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct StakingConfig {
    pub mint_address: Pubkey,
    pub yield_per: u64,
    pub yield_per_time: u64,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Reward {
    pub mint_address: Pubkey,
    pub threshold: u8,
    // decimals are hardcoded to 1 in ./ix_accounts.rs:InitializeRewardToken
    pub amount: u64,
    pub authority_enum: u8, // 0 - quests_authortiy, 1 - quest_authority
    pub cap: u64,
    pub counter: u64,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Tender {
    pub mint_address: Pubkey,
    pub amount: u64,
}

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Split {
    pub token_address: Pubkey,
    pub op_code: u8, // 0 - burn, 1 - transfer to `token_address`
    pub share: u8,
}
