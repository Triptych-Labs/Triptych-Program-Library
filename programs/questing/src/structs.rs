use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Reward {
    pub mint_address: Pubkey,
    pub threshold: u8,
    // decimals are hardcoded to 1 in ./ix_accounts.rs:InitializeRewardToken
    pub amount: u64,
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
