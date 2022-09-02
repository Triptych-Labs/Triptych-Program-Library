use anchor_lang::prelude::*;
use anchor_spl::associated_token::{self, Create};
use anchor_spl::token::{self, Burn, Mint, MintTo, Transfer};
use constants::*;
use errors::QuestError;
use helper_fns::*;
use ix_accounts::*;
use std::result::Result;
use structs::*;

declare_id!("9iMuz8Lf27R9Y2jQhWM1wrSVtPB4Tt5wqkh1opjMTK11");

mod constants;
mod errors;
pub mod helper_fns;
mod ix_accounts;
pub mod state;
pub mod structs;

#[program]
pub mod questing {

    use solana_program::program::invoke_signed;

    use super::*;

    pub fn enable_quests(ctx: Context<EnableQuests>) -> Result<(), Error> {
        let quests = &mut ctx.accounts.quests;
        quests.oracle = ctx.accounts.oracle.key();
        quests.quests = 0;

        Ok(())
    }

    pub fn register_quests_staking_reward(
        ctx: Context<InitializeQuestsStakingReward>,
        quests_bump: u8,
        name: String,
        symbol: String,
        uri: String,
    ) -> Result<(), Error> {
        let mpl_metadata_program = &ctx.accounts.mpl_metadata_program;
        let metadata_account = &ctx.accounts.metadata_account;
        let reward_mint = &ctx.accounts.reward_mint;
        let quests = &ctx.accounts.quests;
        let oracle = &ctx.accounts.oracle;
        let system_program = &ctx.accounts.system_program;
        let rent = &ctx.accounts.rent;

        let oracle_key = oracle.key();
        let quests_bump_bytes = quests_bump.to_le_bytes();
        let quests_authority_seeds = &[
            QUEST_ORACLE_SEED.as_ref(),
            oracle_key.as_ref(),
            quests_bump_bytes.as_ref(),
        ];
        let quests_authority = &[&quests_authority_seeds[..]];

        invoke_signed(
            &mpl_token_metadata::instruction::create_metadata_accounts_v2(
                mpl_metadata_program.key(),
                metadata_account.key(),
                reward_mint.key(),
                quests.key(),
                oracle.key(),
                quests.key(),
                name,
                symbol,
                uri,
                None,
                0,
                true,
                true,
                None,
                None,
            ),
            &[
                metadata_account.to_account_info(),
                reward_mint.to_account_info(),
                quests.to_account_info(),
                oracle.to_account_info(),
                system_program.to_account_info(),
                rent.to_account_info(),
            ],
            quests_authority,
        )?;

        Ok(())
    }

    pub fn register_quests_reward(
        ctx: Context<InitializeQuestsRewardToken>,
        quests_bump: u8,
        name: String,
        symbol: String,
        uri: String,
    ) -> Result<(), Error> {
        let mpl_metadata_program = &ctx.accounts.mpl_metadata_program;
        let metadata_account = &ctx.accounts.metadata_account;
        let oracle = &ctx.accounts.oracle;
        let quests = &mut ctx.accounts.quests;

        let system_program = &ctx.accounts.system_program;
        let rent = &ctx.accounts.rent;
        let reward_mint = &ctx.accounts.reward_mint;

        let oracle_key = oracle.key();
        let quests_bump_bytes = quests_bump.to_le_bytes();
        let quests_authority_seeds = &[
            QUEST_ORACLE_SEED.as_ref(),
            oracle_key.as_ref(),
            quests_bump_bytes.as_ref(),
        ];
        let quests_authority = &[&quests_authority_seeds[..]];

        invoke_signed(
            &mpl_token_metadata::instruction::create_metadata_accounts_v2(
                mpl_metadata_program.key(),
                metadata_account.key(),
                reward_mint.key(),
                quests.key(),
                oracle.key(),
                quests.key(),
                name,
                symbol,
                uri,
                None,
                0,
                true,
                true,
                None,
                None,
            ),
            &[
                metadata_account.to_account_info(),
                reward_mint.to_account_info(),
                quests.to_account_info(),
                oracle.to_account_info(),
                system_program.to_account_info(),
                rent.to_account_info(),
            ],
            quests_authority,
        )?;

        Ok(())
    }

    pub fn register_quest_reward(
        ctx: Context<InitializeQuestRewardToken>,
        quest_bump: u8,
        _quest_index: u64,
        reward: Reward,
        name: String,
        symbol: String,
        uri: String,
    ) -> Result<(), Error> {
        let mpl_metadata_program = &ctx.accounts.mpl_metadata_program;
        let metadata_account = &ctx.accounts.metadata_account;
        let oracle = &ctx.accounts.oracle;
        let quest = &mut ctx.accounts.quest;
        let system_program = &ctx.accounts.system_program;
        let rent = &ctx.accounts.rent;
        let reward_mint = &ctx.accounts.reward_mint;

        let oracle_key = oracle.key();
        let quest_bump_bytes = quest_bump.to_le_bytes();
        let quest_index_bytes = quest.index.to_le_bytes();
        let quest_authority_seeds = &[
            QUEST_ORACLE_SEED.as_ref(),
            oracle_key.as_ref(),
            quest_index_bytes.as_ref(),
            quest_bump_bytes.as_ref(),
        ];
        let quest_authority = &[&quest_authority_seeds[..]];

        // assert valid reward threshold sum before appending to rewards
        assert_valid_sum_rng_thresholds(quest.rewards.clone(), reward.threshold).unwrap();
        quest.rewards.push(reward);

        invoke_signed(
            &mpl_token_metadata::instruction::create_metadata_accounts_v2(
                mpl_metadata_program.key(),
                metadata_account.key(),
                reward_mint.key(),
                quest.key(),
                oracle.key(),
                quest.key(),
                name,
                symbol,
                uri,
                None,
                0,
                true,
                true,
                None,
                None,
            ),
            &[
                metadata_account.to_account_info(),
                reward_mint.to_account_info(),
                quest.to_account_info(),
                oracle.to_account_info(),
                system_program.to_account_info(),
                rent.to_account_info(),
            ],
            quest_authority,
        )?;

        Ok(())
    }

    pub fn create_quest(
        ctx: Context<CreateQuest>,
        quest_index: u64,
        name: String,
        duration: i64,
        wl_candy_machines: Vec<Pubkey>,
        tender: Option<Tender>,
        tender_splits: Option<Vec<Split>>,
        xp: u64,
        required_level: Option<u64>,
        enabled: bool,
        staking_config: Option<StakingConfig>,
        pairs_config: Option<PairsConfig>,
        rewards: Vec<Reward>,
    ) -> Result<(), Error> {
        let quest = &mut ctx.accounts.quest;
        let quests = &mut ctx.accounts.quests;

        if quests.quests != quest_index {
            return Err(QuestError::UnexpectedQuestingState.into());
        }

        quest.enabled = enabled;
        quest.index = quest_index;
        quest.name = name;
        quest.duration = duration;
        quest.oracle = ctx.accounts.oracle.key();
        quest.wl_candy_machines = wl_candy_machines;
        quest.rewards = rewards;

        if tender.is_some() && tender_splits.is_some() {
            quest.tender = tender;
            quest.tender_splits = tender_splits;
        }
        quest.xp = xp;

        if required_level.is_some() {
            quest.required_level = required_level.unwrap();
            quest.required_xp = get_sum_xp_from_level(quest.required_level).unwrap();
        } else {
            quest.required_level = 0;
            quest.required_xp = 0;
        }

        if staking_config.is_some() {
            quest.staking_config = staking_config;
        }
        if pairs_config.is_some() {
            let pairs_config_unwrapped = pairs_config.unwrap();
            if (pairs_config_unwrapped.left + pairs_config_unwrapped.right) > 5 {
                return Err(QuestError::RatioTooBig.into());
            }
            quest.pairs_config = Some(pairs_config_unwrapped);
        }

        quests.quests += 1;

        Ok(())
    }

    pub fn register_quest_recorder(ctx: Context<RegisterQuestRecorder>) -> Result<(), Error> {
        let quest_recorder = &mut ctx.accounts.quest_recorder;
        quest_recorder.proposals = 0;
        quest_recorder.staked = Vec::new();
        quest_recorder.quest = ctx.accounts.quest.key();
        quest_recorder.initializer = ctx.accounts.initializer.key();
        quest_recorder.oracle = ctx.accounts.quest.oracle;

        Ok(())
    }

    pub fn propose_quest_record(
        ctx: Context<ProposeQuestRecord>,
        depositing_left: Vec<Pubkey>,
        depositing_right: Vec<Pubkey>,
    ) -> Result<(), Error> {
        let quest = &mut ctx.accounts.quest;
        let quest_recorder = &mut ctx.accounts.quest_recorder;
        let quest_proposal = &mut ctx.accounts.quest_proposal;

        let quest_pairs_config = quest.pairs_config.clone().unwrap();

        if (depositing_left.len() + depositing_right.len())
            != (quest_pairs_config.left + quest_pairs_config.right) as usize
        {
            return Err(QuestError::RatioTooBig.into());
        }

        if depositing_left.len() > quest_pairs_config.left as usize {
            return Err(QuestError::RatioTooBig.into());
        }

        if depositing_right.len() > quest_pairs_config.right as usize {
            return Err(QuestError::RatioTooBig.into());
        }

        quest_proposal.index = quest_recorder.proposals;
        quest_proposal.fulfilled = false;
        quest_proposal.started = false;
        quest_proposal.finished = false;
        quest_proposal.withdrawn = false;
        quest_proposal.depositing_left = depositing_left;
        quest_proposal.depositing_right = depositing_right;
        quest_proposal.record_left = Vec::new();
        quest_proposal.record_right = Vec::new();

        quest_recorder.proposals += 1;

        Ok(())
    }

    pub fn enter_quest(
        ctx: Context<EnterQuest>,
        _quest_proposal_index: u64,
        _quest_proposal_bump: u8,
        side_enum: String,
    ) -> Result<(), Error> {
        let token_program = &ctx.accounts.token_program;
        let initializer = &mut ctx.accounts.initializer;
        let quest = &mut ctx.accounts.quest;
        let pixelballz_token_account = &mut ctx.accounts.pixelballz_token_account;
        let quest_proposal = &mut ctx.accounts.quest_proposal;
        let quest_pairs_config = quest.pairs_config.clone().unwrap();

        if quest_proposal.started {
            return Err(QuestError::ProposalStarted.into());
        }

        quest_proposal.started = false;

        if pixelballz_token_account.amount != 1 {
            msg!("empty token account");
            return Err(QuestError::SuspiciousTransaction.into());
        }

        match side_enum.as_str() {
            "left" => {
                if !quest_proposal
                    .depositing_left
                    .iter()
                    .any(|&depositing| depositing == pixelballz_token_account.mint)
                {
                    return Err(QuestError::SuspiciousTransaction.into());
                }
                let validity = assert_valid_token(
                    &ctx.accounts.pixelballz_metadata.to_account_info(),
                    &pixelballz_token_account.mint,
                    quest_pairs_config.left_creators.to_vec(),
                )?;
                if !validity {
                    msg!("invalid nft");
                    return Err(QuestError::SuspiciousTransaction.into());
                }

                quest_proposal.record_left.push(true);
            }
            "right" => {
                if !quest_proposal
                    .depositing_right
                    .iter()
                    .any(|&depositing| depositing == pixelballz_token_account.mint)
                {
                    return Err(QuestError::SuspiciousTransaction.into());
                }
                let validity = assert_valid_token(
                    &ctx.accounts.pixelballz_metadata.to_account_info(),
                    &pixelballz_token_account.mint,
                    quest_pairs_config.right_creators.to_vec(),
                )?;
                if !validity {
                    msg!("invalid nft");
                    return Err(QuestError::SuspiciousTransaction.into());
                }
                quest_proposal.record_right.push(true);
            }
            _ => {
                return Err(QuestError::SuspiciousTransaction.into());
            }
        }

        if (quest_proposal.record_left.len() + quest_proposal.record_right.len())
            == (quest_pairs_config.left + quest_pairs_config.right) as usize
        {
            quest_proposal.fulfilled = true;
        }

        token::set_authority(
            CpiContext::new(
                token_program.to_account_info(),
                token::SetAuthority {
                    current_authority: initializer.to_account_info(),
                    account_or_mint: pixelballz_token_account.to_account_info(),
                },
            ),
            spl_token::instruction::AuthorityType::AccountOwner,
            Some(quest.key()),
        )?;

        Ok(())
    }

    pub fn start_quest<'info>(
        ctx: Context<'_, '_, '_, 'info, StartQuest<'info>>,
        _quest_proposal_index: u64,
        _quest_proposal_bump: u8,
        _quest_recorder_bump: u8,
    ) -> Result<(), Error> {
        let now = Clock::get()?.unix_timestamp;
        let quest = &mut ctx.accounts.quest;
        let quest_proposal = &mut ctx.accounts.quest_proposal;
        let quest_account = &mut ctx.accounts.quest_acc;
        let token_program = &ctx.accounts.token_program;
        let initializer = &ctx.accounts.initializer;
        let quest_recorder = &mut ctx.accounts.quest_recorder;

        for deposit in quest_proposal.depositing_left.clone().into_iter() {
            quest_recorder.staked.push(deposit);
        }

        for deposit in quest_proposal.depositing_right.clone().into_iter() {
            quest_recorder.staked.push(deposit);
        }

        if quest_proposal.fulfilled {
            quest_proposal.started = true;
        } else {
            return Err(QuestError::SuspiciousTransaction.into());
        }

        quest_account.quest = quest.key();
        quest_account.index = quest.index;
        quest_account.start_time = now;
        quest_account.end_time = now + quest.duration;
        // quest_account.deposit_token_mint = deposit_token_account.mint;
        quest_account.initializer = initializer.key();
        quest_account.last_claim = now;

        // here we protect restarting an already completed `quest_account`.
        // this is only ever set in `end_quest` ix to designate a safe re-init
        // of `quest_account`. basically if `Some(quest_account.completed) == false`,
        // then reset its completion. inversely, if the complete field is true then
        // throw an err preventing a future unsafe re-init of `quest_account`.
        //
        // this effectively allows quest re-entry for those who withdrew earlier
        // than the calculated end date.

        /*
        if quest_account.completed != None {
            if quest_account.completed == Some(true) {
                return Err(QuestError::InvalidCompletion.into());
            }
        }
        */
        quest_account.completed = Some(false);

        // TODO verify the creator address from nft metadata account
        let mut remaining_accounts = 0;
        let ctx_remaining_accounts = ctx.remaining_accounts.to_vec();

        if quest.tender.is_some() {
            let tender = quest.tender.clone().unwrap();
            let tender_splits = quest.tender_splits.clone().unwrap();

            if tender_splits.len() == 0 {
                return Err(QuestError::SuspiciousTransaction.into());
            }
            let tender_token_account = &ctx_remaining_accounts[remaining_accounts];
            remaining_accounts += 1;

            let mut i = 0;
            while i < tender_splits.len() {
                let tender_split = tender_splits[i].clone();

                if tender_split.op_code == 0 {
                    let tendersplit_mint = &ctx_remaining_accounts[remaining_accounts];
                    remaining_accounts += 1;

                    msg!(
                        "{:?} {:?}",
                        tendersplit_mint.key(),
                        tender_token_account.key()
                    );
                    token::burn(
                        CpiContext::new(
                            token_program.to_account_info(),
                            Burn {
                                mint: tendersplit_mint.to_account_info(),
                                from: tender_token_account.to_account_info(),
                                authority: initializer.to_account_info(),
                            },
                        ),
                        (tender.amount as f64 * ((tender_split.share as u64) as f64 / 100.0))
                            as u64,
                    )?;
                }

                if tender_split.op_code == 1 {
                    let tendersplit_token_account = &ctx_remaining_accounts[remaining_accounts];
                    remaining_accounts += 1;

                    token::transfer(
                        CpiContext::new(
                            token_program.to_account_info(),
                            Transfer {
                                from: tender_token_account.to_account_info(),
                                to: tendersplit_token_account.to_account_info(),
                                authority: initializer.to_account_info(),
                            },
                        ),
                        (tender.amount as f64 * ((tender_split.share as u64) as f64 / 100.0))
                            as u64,
                    )?;
                }

                i += 1;
            }
        }

        Ok(())
    }

    pub fn flush_quest_record<'info>(
        ctx: Context<'_, '_, '_, 'info, FlushQuestRecord<'info>>,
        _quest_proposal_index: u64,
        _quest_proposal_bump: u8,
        quest_bump: u8,
    ) -> Result<(), Error> {
        let token_program = &ctx.accounts.token_program;
        let pixelballz_token_account = &mut ctx.accounts.pixelballz_token_account;
        let pixelballz_mint = &mut ctx.accounts.pixelballz_mint;
        let pixelballz_mint_key = pixelballz_mint.key();
        let quest = &mut ctx.accounts.quest;
        let quest_proposal = &mut ctx.accounts.quest_proposal;

        if quest_proposal.started {
            if !quest_proposal.finished {
                return Err(QuestError::ProposalNotFinished.into());
            }
        }

        let mut is_valid = false;
        if quest_proposal
            .depositing_left
            .iter()
            .any(|&deposit| deposit == pixelballz_mint_key)
        {
            // quest_proposal.record_left.swap_remove(0);
            is_valid = true;
        }
        if quest_proposal
            .depositing_right
            .iter()
            .any(|&deposit| deposit == pixelballz_mint_key)
        {
            // quest_proposal.record_right.swap_remove(0);
            is_valid = true;
        }
        if !is_valid {
            return Err(QuestError::SuspiciousTransaction.into());
        }

        if quest_proposal.record_left.len() == 0 && quest_proposal.record_right.len() == 0 {
            quest_proposal.withdrawn = true;
        }

        let quest_index_bytes = quest.index.to_le_bytes();
        let quest_bump_bytes = quest_bump.to_le_bytes();

        let seeds_with_bump = &[
            QUEST_ORACLE_SEED.as_ref(),
            quest.oracle.as_ref(),
            quest_index_bytes.as_ref(),
            quest_bump_bytes.as_ref(),
        ];
        let quest_authority = &[&seeds_with_bump[..]];

        // Transfer NFT quested back
        token::set_authority(
            CpiContext::new_with_signer(
                token_program.to_account_info(),
                token::SetAuthority {
                    current_authority: quest.to_account_info(),
                    account_or_mint: pixelballz_token_account.to_account_info(),
                },
                quest_authority,
            ),
            spl_token::instruction::AuthorityType::AccountOwner,
            Some(ctx.accounts.initializer.key()),
        )?;

        Ok(())
    }

    pub fn claim_quest_staking_reward<'info>(
        ctx: Context<'_, '_, '_, 'info, ClaimQuestStakingReward<'info>>,
        quests_bump: u8,
    ) -> Result<(), Error> {
        let now = Clock::get()?.unix_timestamp;
        let quest = &mut ctx.accounts.quest;
        let quest_account = &mut ctx.accounts.quest_acc;
        let initializer = &ctx.accounts.initializer;
        let reward_mint = &ctx.accounts.reward_mint;
        let reward_token_account = &ctx.accounts.reward_token_account;

        let staking_config = quest.staking_config.clone().unwrap();
        if staking_config.mint_address != reward_mint.key() {
            return Err(QuestError::InvalidRewardMint.into());
        }

        if reward_token_account.data.borrow().len() == 0 {
            let _may_fail = associated_token::create(CpiContext::new(
                ctx.accounts.associated_token_program.to_account_info(),
                Create {
                    payer: initializer.to_account_info(),
                    associated_token: reward_token_account.to_account_info(),
                    authority: initializer.to_account_info(),
                    mint: reward_mint.to_account_info(),
                    system_program: ctx.accounts.system_program.to_account_info(),
                    token_program: ctx.accounts.token_program.to_account_info(),
                    rent: ctx.accounts.rent.to_account_info(),
                },
            ));
        }

        let oracle_key = quest.oracle;
        let quests_bump_bytes = quests_bump.to_le_bytes();
        let reward_authority_seeds = &[
            QUEST_ORACLE_SEED.as_ref(),
            oracle_key.as_ref(),
            quests_bump_bytes.as_ref(),
        ];
        let reward_authority = &[&reward_authority_seeds[..]];

        if now > quest_account.end_time && quest_account.last_claim == quest_account.end_time {
            return Ok(());
        }

        let duration = now - quest_account.last_claim;
        quest_account.last_claim = now;
        // TODO REIMPLEMENT
        /*
        if now >= quest_account.end_time {
            duration = quest_account.end_time - quest_account.last_claim;
            quest_account.last_claim = quest_account.end_time;
        } else {
            duration = now - quest_account.last_claim;
            quest_account.last_claim = now;
        }
        */

        let beta = duration as f64 / staking_config.yield_per_time as f64;
        let alpha = (staking_config.yield_per as f64 * beta as f64) as u64;

        msg!("{} {} {}", duration, beta, alpha);

        // Mint reward token amount to initializer
        token::mint_to(
            CpiContext::new_with_signer(
                ctx.accounts.token_program.to_account_info(),
                MintTo {
                    mint: reward_mint.to_account_info(),
                    to: reward_token_account.to_account_info(),
                    authority: ctx.accounts.quests.to_account_info(),
                },
                reward_authority,
            ),
            alpha,
        )?;

        Ok(())
    }

    pub fn end_quest<'info>(
        ctx: Context<'_, '_, '_, 'info, EndQuest<'info>>,
        _quest_proposal_index: u64,
        _quest_proposal_bump: u8,
        _quest_recorder_bump: u8,
        quests_bump: u8,
        quest_bump: u8,
    ) -> Result<(), Error> {
        let now = Clock::get()?.unix_timestamp;
        let oracle = &ctx.accounts.oracle;
        let quest = &mut ctx.accounts.quest;
        let quests = &mut ctx.accounts.quests;
        let quest_account = &mut ctx.accounts.quest_acc;

        let quest_proposal = &mut ctx.accounts.quest_proposal;
        let quest_recorder = &mut ctx.accounts.quest_recorder;

        /*
        for deposit in quest_proposal.depositing_left.clone().into_iter() {
            quest_recorder.staked = quest_recorder
                .staked
                .clone()
                .into_iter()
                .filter(|&stake| stake != deposit)
                .collect();
        }

        for deposit in quest_proposal.depositing_right.clone().into_iter() {
            quest_recorder.staked = quest_recorder
                .staked
                .clone()
                .into_iter()
                .filter(|&stake| stake != deposit)
                .collect();
        }
        */

        if quest_account.end_time > now {
            msg!("ending early {} {}", quest_account.end_time, now);
            return Err(QuestError::SuspiciousTransaction.into());
        }

        quest_account.completed = Some(true);
        quest_proposal.finished = true;

        let mut remaining_accounts: usize = 0;

        if quest.rewards.len() > 0 {
            // RNG a number off recent blockhash
            let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();
            let most_recent = &recent_slot_hash[12..25];
            // nominate for r/shittyprogramming 2022 meme of the year pls
            let percentile: u8 = (u64::from_le_bytes([
                most_recent[0],
                most_recent[1],
                most_recent[2],
                most_recent[3],
                most_recent[4],
                most_recent[5],
                most_recent[6],
                most_recent[7],
            ]) % 100) as u8;

            let (reward_index, reward) =
                convict_percentile_against_thresholds(percentile, quest.rewards.clone()).unwrap();
            if reward_index != -1 {
                // select correct bump from
                let reward_mint = &ctx.remaining_accounts.to_vec()
                    [remaining_accounts..remaining_accounts + quest.rewards.len()]
                    [reward_index as usize];
                remaining_accounts += quest.rewards.len();

                let reward_token_account = &ctx.remaining_accounts.to_vec()
                    [remaining_accounts..remaining_accounts + quest.rewards.len()]
                    [reward_index as usize];
                remaining_accounts += quest.rewards.len();

                if reward_token_account.data.borrow().len() == 0 {
                    let _may_fail = associated_token::create(CpiContext::new(
                        ctx.accounts.associated_token_program.to_account_info(),
                        Create {
                            payer: ctx.accounts.initializer.to_account_info(),
                            associated_token: reward_token_account.to_account_info(),
                            authority: ctx.accounts.initializer.to_account_info(),
                            mint: reward_mint.to_account_info(),
                            system_program: ctx.accounts.system_program.to_account_info(),
                            token_program: ctx.accounts.token_program.to_account_info(),
                            rent: ctx.accounts.rent.to_account_info(),
                        },
                    ));
                }

                let oracle_key = oracle.key();
                let reward_authority;
                let reward_authority_signer;
                let quest_index_bytes = quest.index.to_le_bytes();
                let quest_bump_bytes = quest_bump.to_le_bytes();
                let oracle_ref = QUEST_ORACLE_SEED.as_ref();
                let oracle_key_ref = oracle_key.as_ref();
                let quests_bump_bytes = quests_bump.to_le_bytes();
                let quests_authority_seeds =
                    &[oracle_ref, oracle_key_ref, quests_bump_bytes.as_ref()];
                let quest_authority_seeds = &[
                    oracle_ref,
                    oracle_key_ref,
                    quest_index_bytes.as_ref(),
                    quest_bump_bytes.as_ref(),
                ];
                match reward.authority_enum {
                    0 => {
                        let seeds = &quests_authority_seeds[..];
                        reward_authority_signer = [seeds];
                        reward_authority = quests.to_account_info();
                    }
                    1 => {
                        let seeds = &quest_authority_seeds[..];
                        reward_authority_signer = [seeds];
                        reward_authority = quest.to_account_info();
                    }
                    _ => return Err(QuestError::SuspiciousTransaction.into()),
                }

                let mut reward_record = reward.clone();

                if reward_record.cap >= reward_record.counter + reward.amount {
                    // Mint reward token amount to initializer
                    token::mint_to(
                        CpiContext::new_with_signer(
                            ctx.accounts.token_program.to_account_info(),
                            MintTo {
                                mint: reward_mint.to_account_info(),
                                to: reward_token_account.to_account_info(),
                                authority: reward_authority,
                            },
                            &reward_authority_signer,
                        ),
                        reward.amount,
                    )?;
                    msg!("minted reward");
                    reward_record.counter += reward.amount;
                    quest.rewards[reward_index as usize] = reward_record;
                } else {
                    msg!("ran out of reward. exceeded cap");
                }
            } else {
                msg!("no reward");
                remaining_accounts += quest.rewards.len() * 2;
            }
        }
        if quests.rewards.len() > 0 {
            // RNG a number off recent blockhash
            let recent_slot_hash = &ctx.accounts.slot_hashes.data.borrow();
            let most_recent = &recent_slot_hash[12..25];
            // nominate for r/shittyprogramming 2022 meme of the year pls
            let percentile: u8 = (u64::from_le_bytes([
                most_recent[0],
                most_recent[1],
                most_recent[2],
                most_recent[3],
                most_recent[4],
                most_recent[5],
                most_recent[6],
                most_recent[7],
            ]) % 100) as u8;

            let (reward_index, reward) =
                convict_percentile_against_thresholds(percentile, quests.rewards.clone()).unwrap();

            // select correct bump from
            let reward_mint = &ctx.remaining_accounts.to_vec()
                [remaining_accounts..remaining_accounts + quests.rewards.len()]
                [reward_index as usize];
            remaining_accounts += quests.rewards.len();

            let reward_token_account = &ctx.remaining_accounts.to_vec()
                [remaining_accounts..remaining_accounts + quests.rewards.len()]
                [reward_index as usize];
            // remaining_accounts += quest.rewards.len();

            if reward_token_account.data.borrow().len() == 0 {
                let _may_fail = associated_token::create(CpiContext::new(
                    ctx.accounts.associated_token_program.to_account_info(),
                    Create {
                        payer: ctx.accounts.initializer.to_account_info(),
                        associated_token: reward_token_account.to_account_info(),
                        authority: ctx.accounts.initializer.to_account_info(),
                        mint: reward_mint.to_account_info(),
                        system_program: ctx.accounts.system_program.to_account_info(),
                        token_program: ctx.accounts.token_program.to_account_info(),
                        rent: ctx.accounts.rent.to_account_info(),
                    },
                ));
            }

            let oracle_key = oracle.key();
            let quests_bump_bytes = quests_bump.to_le_bytes();
            let reward_authority_seeds = &[
                QUEST_ORACLE_SEED.as_ref(),
                oracle_key.as_ref(),
                quests_bump_bytes.as_ref(),
            ];
            let reward_authority = &[&reward_authority_seeds[..]];

            // Mint reward token amount to initializer
            token::mint_to(
                CpiContext::new_with_signer(
                    ctx.accounts.token_program.to_account_info(),
                    MintTo {
                        mint: reward_mint.to_account_info(),
                        to: reward_token_account.to_account_info(),
                        authority: quests.to_account_info(),
                    },
                    reward_authority,
                ),
                reward.amount,
            )?;
        }

        Ok(())
    }
}
