use anchor_lang::prelude::*;

#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct Split {
    pub token_address: Pubkey,
    pub op_code: u8, // 0 - burn, 1 - transfer to `token_address`
    pub share: u8,
}
