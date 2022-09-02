use crate::constants::*;
use crate::state::*;
use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::{Mint, Token, TokenAccount};

#[derive(Accounts)]
pub struct RegisterQuestRecorder<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        init,
        seeds = [QUEST_RECORDER.as_ref(), quest.key().as_ref(), initializer.key().as_ref()],
        bump,
        payer = initializer,
        space = QuestRecorder::space(0)
    )]
    pub quest_recorder: Box<Account<'info, QuestRecorder>>,
    pub quest: Box<Account<'info, Quest>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct ProposeQuestRecord<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        init,
        seeds = [quest.key().as_ref(), initializer.key().as_ref(), quest_recorder.proposals.to_le_bytes().as_ref()],
        bump,
        payer = initializer,
        space = QuestProposal::space(quest.pairs_config.clone().unwrap().left, quest.pairs_config.clone().unwrap().right)
    )]
    pub quest_proposal: Box<Account<'info, QuestProposal>>,
    pub quest: Box<Account<'info, Quest>>,
    #[account(mut)]
    pub quest_recorder: Box<Account<'info, QuestRecorder>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(quest_proposal_index: u64, quest_proposal_bump: u8)]
pub struct EnterQuest<'info> {
    /// CHECK: checked in ix
    pub pixelballz_metadata: UncheckedAccount<'info>,
    #[account(mut)]
    pub pixelballz_token_account: Box<Account<'info, TokenAccount>>,
    #[account(mut)]
    pub quest: Box<Account<'info, Quest>>,
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        mut,
        seeds = [quest.key().as_ref(), initializer.key().as_ref(), quest_proposal_index.to_le_bytes().as_ref()],
        bump = quest_proposal_bump,
    )]
    pub quest_proposal: Box<Account<'info, QuestProposal>>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
}

#[derive(Accounts)]
#[instruction(quest_proposal_index: u64, quest_proposal_bump: u8)]
pub struct FlushQuestRecord<'info> {
    #[account(
        mut,
        seeds = [quest.key().as_ref(), initializer.key().as_ref(), quest_proposal_index.to_le_bytes().as_ref()],
        bump = quest_proposal_bump,
    )]
    pub quest_proposal: Box<Account<'info, QuestProposal>>,
    #[account(mut)]
    pub quest: Account<'info, Quest>,
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    pub pixelballz_mint: Box<Account<'info, Mint>>,
    #[account(mut)]
    pub pixelballz_token_account: Box<Account<'info, TokenAccount>>,
    pub token_program: Program<'info, Token>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(quest_proposal_index: u64, quest_proposal_bump: u8, quest_recorder_bump: u8 )]
pub struct StartQuest<'info> {
    #[account(mut)]
    pub quest: Box<Account<'info, Quest>>,
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(
        mut,
        seeds = [quest.key().as_ref(), initializer.key().as_ref(), quest_proposal_index.to_le_bytes().as_ref()],
        bump = quest_proposal_bump,
    )]
    pub quest_proposal: Box<Account<'info, QuestProposal>>,
    #[account(
        init,
        seeds = [QUEST_PDA_SEED.as_ref(), initializer.key().as_ref(), quest_proposal.key().as_ref(), quest.key().as_ref()],
        bump,
        payer = initializer,
        space = QuestAccount::LEN
    )]
    pub quest_acc: Account<'info, QuestAccount>,
    #[account(
        mut,
        seeds = [QUEST_RECORDER.as_ref(), quest.key().as_ref(), initializer.key().as_ref()],
        bump = quest_recorder_bump,
        realloc = QuestRecorder::space(quest_recorder.staked.len() + quest_proposal.depositing_left.len() + quest_proposal.depositing_right.len()),
        realloc::payer = initializer,
        realloc::zero = false,
    )]
    pub quest_recorder: Box<Account<'info, QuestRecorder>>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
}

#[derive(Accounts)]
#[instruction(quest_proposal_index: u64, quest_proposal_bump: u8, quest_recorder_bump: u8)]
pub struct EndQuest<'info> {
    #[account(
        mut,
        seeds = [QUEST_RECORDER.as_ref(), quest.key().as_ref(), initializer.key().as_ref()],
        bump = quest_recorder_bump,
    )]
    pub quest_recorder: Box<Account<'info, QuestRecorder>>,
    #[account(mut)]
    pub quest_acc: Box<Account<'info, QuestAccount>>,
    #[account(
        mut,
        seeds = [quest.key().as_ref(), initializer.key().as_ref(), quest_proposal_index.to_le_bytes().as_ref()],
        bump = quest_proposal_bump,
    )]
    pub quest_proposal: Box<Account<'info, QuestProposal>>,
    #[account(mut)]
    pub quest: Account<'info, Quest>,
    #[account(mut)]
    pub quests: Account<'info, Quests>,
    #[account(mut)]
    pub initializer: Signer<'info>,
    pub token_program: Program<'info, Token>,
    pub system_program: Program<'info, System>,
    pub rent: Sysvar<'info, Rent>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    /// CHECK: am lazy
    pub slot_hashes: UncheckedAccount<'info>,
    #[account(mut)]
    /// CHECK: am lazy
    pub oracle: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(quest_index: u64)]
pub struct CreateQuest<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        init,
        seeds = [QUEST_ORACLE_SEED.as_ref(), oracle.key().as_ref(), &quest_index.to_le_bytes()],
        bump,
        payer = oracle,
        space = Quest::space(0)
    )]
    pub quest: Box<Account<'info, Quest>>,
    #[account(mut)]
    pub quests: Box<Account<'info, Quests>>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
#[instruction(quests_bump: u8)]
pub struct InitializeQuestsRewardToken<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        mut,
        seeds = [QUEST_ORACLE_SEED.as_ref(), oracle.key().as_ref()],
        bump = quests_bump
    )]
    pub quests: Box<Account<'info, Quests>>,
    #[account(
        init,
        payer = oracle,
        mint::decimals = 1,
        mint::authority = quests,
    )]
    pub reward_mint: Account<'info, Mint>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
    /// CHECK: checked in cpi
    pub mpl_metadata_program: UncheckedAccount<'info>,
    /// CHECK: checked in cpi
    #[account(mut)]
    pub metadata_account: UncheckedAccount<'info>,
}

#[derive(Accounts)]
#[instruction(quest_bump: u8, quest_index: u64)]
pub struct InitializeQuestRewardToken<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        mut,
        seeds = [QUEST_ORACLE_SEED.as_ref(), oracle.key().as_ref(), &quest_index.to_le_bytes()],
        bump = quest_bump,
        realloc = Quest::space(quest.rewards.len() + 1),
        realloc::payer = oracle,
        realloc::zero = false,
    )]
    pub quest: Box<Account<'info, Quest>>,
    #[account(
        init,
        payer = oracle,
        mint::decimals = 1,
        mint::authority = quest,
    )]
    pub reward_mint: Account<'info, Mint>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
    /// CHECK: checked in cpi
    pub mpl_metadata_program: UncheckedAccount<'info>,
    /// CHECK: checked in cpi
    #[account(mut)]
    pub metadata_account: UncheckedAccount<'info>,
}

#[derive(Accounts)]
pub struct InitializeQuestsStakingReward<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(has_one = oracle)]
    pub quests: Box<Account<'info, Quests>>,
    #[account(
        init,
        payer = oracle,
        mint::decimals = 1,
        mint::authority = quests,
    )]
    pub reward_mint: Account<'info, Mint>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    /// CHECK: checked in cpi
    pub mpl_metadata_program: UncheckedAccount<'info>,
    /// CHECK: checked in cpi
    #[account(mut)]
    pub metadata_account: UncheckedAccount<'info>,
    pub rent: Sysvar<'info, Rent>,
}

#[derive(Accounts)]
pub struct ClaimQuestStakingReward<'info> {
    #[account(mut)]
    pub initializer: Signer<'info>,
    #[account(mut)]
    pub quests: Box<Account<'info, Quests>>,
    pub quest: Box<Account<'info, Quest>>,
    #[account(mut, has_one = initializer)]
    pub quest_acc: Box<Account<'info, QuestAccount>>,
    #[account(mut)]
    pub reward_mint: Account<'info, Mint>,
    /// CHECK: will later init if needed in logic
    #[account(mut)]
    pub reward_token_account: UncheckedAccount<'info>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub rent: Sysvar<'info, Rent>,
    pub associated_token_program: Program<'info, AssociatedToken>,
}

#[derive(Accounts)]
pub struct EnableQuests<'info> {
    #[account(mut)]
    pub oracle: Signer<'info>,
    #[account(
        init,
        seeds = [QUEST_ORACLE_SEED.as_ref(), oracle.key().as_ref()],
        bump,
        payer = oracle,
        space = Quests::LEN
    )]
    pub quests: Box<Account<'info, Quests>>,
    pub system_program: Program<'info, System>,
}
