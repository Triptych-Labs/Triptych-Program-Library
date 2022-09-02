use crate::errors::*;
use crate::structs::*;
use anchor_lang::prelude::*;
use mpl_token_metadata::state::Metadata;
use std::result::Result;
use std::str::FromStr;

pub fn assert_valid_sum_rng_thresholds(
    rewards: Vec<Reward>,
    additional_rng_threshold: u8,
) -> Result<(), Error> {
    let mut sum: u8 = additional_rng_threshold;
    let mut i = 0;
    while i < rewards.len() {
        if sum > 100 {
            msg!("sum rng theshold must not exceed 100!");
            return Err(QuestError::IndexGreaterThanLength.into());
        }
        sum += rewards[i].threshold;
        i += 1;
    }

    Ok(())
}

pub fn get_sum_xp_from_level(level: u64) -> Result<u64, Error> {
    if level == 0 {
        return Ok(level);
    }
    let base_xp: u64 = 50;
    return Ok((u64::pow(2, level as u32)) * base_xp);
}

pub fn convict_percentile_against_thresholds(
    percentile: u8,
    thresholds: Vec<Reward>,
) -> Result<(i64, Reward), Error> {
    let mut accumulator: u8 = 0;
    let mut i: i64 = 0;
    let thresholds_length = thresholds.len() as i64;
    while i < thresholds_length {
        let pct = thresholds[i as usize].clone();
        let previous_accumulator = accumulator;
        accumulator += pct.threshold;
        if previous_accumulator <= percentile && percentile <= accumulator {
            return Ok((i, pct));
        }
        i += 1;
    }

    Ok((
        -1,
        Reward {
            mint_address: thresholds[0].mint_address,
            threshold: 0,
            amount: 0,
            authority_enum: 0,
            cap: 0,
            counter: 0,
        },
    ))
}

pub fn assert_valid_metadata(
    metadata: &AccountInfo,
    mint: &Pubkey,
) -> Result<Metadata, ProgramError> {
    let metadata_program = Pubkey::from_str("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s").unwrap();
    assert_eq!(metadata.owner, &metadata_program);
    let seed = &[
        b"metadata".as_ref(),
        metadata_program.as_ref(),
        mint.as_ref(),
    ];
    let (metadata_addr, _bump) = Pubkey::find_program_address(seed, &metadata_program);
    assert_eq!(metadata_addr, metadata.key());

    Metadata::from_account_info(metadata)
}

pub fn assert_valid_token(
    metadata: &AccountInfo,
    mint: &Pubkey,
    first_creators: Vec<Pubkey>,
) -> Result<bool, ProgramError> {
    let metadata_data = assert_valid_metadata(metadata, mint)?;
    let creators = metadata_data.data.creators.unwrap();
    let mut first_creators_iter = first_creators.into_iter();

    if first_creators_iter.any(|creator| creator == creators[0].address)
        && creators[0].verified == true
    {
        return Ok(true);
    }
    if first_creators_iter
        .any(|creator| creator == Pubkey::from_str("11111111111111111111111111111111").unwrap())
    {
        return Ok(true);
    }

    Ok(false)
}
