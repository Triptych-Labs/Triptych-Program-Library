use crate::{constants::*, state};
use anchor_lang::prelude::*;
use escrow::cpi::accounts::InstallEscrow;
use std::cell::Ref;
use std::result::Result;

pub fn attenuate_player(hand: &mut [u8; 5], recent_slot_hash: &Ref<&mut [u8]>) -> bool {
    msg!("atten_player {:?}", hand);
    let most_recent = &recent_slot_hash[13..25];
    let rng: u8 = ((u64::from_le_bytes([
        most_recent[0],
        most_recent[1],
        most_recent[2],
        most_recent[3],
        most_recent[4],
        most_recent[5],
        most_recent[6],
        most_recent[7],
    ]) % 10) as u8)
        + 1;

    msg!("RNG: {}", rng);
    let first_zero = hand.iter().position(|&card| card == 0).unwrap();
    hand[first_zero] = rng;

    // sanitize the hand for a busted hit
    let busted = sanitize_hand(hand);
    if busted {
        return busted;
    }
    false
}

pub fn attenuate_dealer(hand: &mut [u8; 5], recent_slot_hash: &Ref<&mut [u8]>) -> bool {
    msg!("atten_dealer {:?}", hand);
    // dealer makes 4 hits or until 18 whichever comes first
    for i in 0..=4 {
        let hand_sum: u8 = hand.iter().sum();
        // soft 17 - casino will always hit upto+including 17
        if hand_sum > 17 {
            break;
        }
        let most_recent = &recent_slot_hash[((i * 32) + 13)..((i * 32) + 25)];
        let rng: u8 = ((u64::from_le_bytes([
            most_recent[0],
            most_recent[1],
            most_recent[2],
            most_recent[3],
            most_recent[4],
            most_recent[5],
            most_recent[6],
            most_recent[7],
        ]) % 10) as u8)
            + 1;

        msg!("RNG: {}", rng);
        let first_zero = hand.iter().position(|&card| card == 0).unwrap();
        hand[first_zero] = rng;

        // sanitize the hand for a busted hit
        let busted = sanitize_hand(hand);
        if busted {
            return busted;
        }
    }

    false
}

pub fn sanitize_hand(hand: &[u8; 5]) -> bool {
    let hand_sum: u8 = hand.iter().sum();
    if hand_sum > 21 {
        return true;
    }

    false
}

pub fn declare_player_win<'a>(
    initializer_info: AccountInfo<'a>,
    house_info: AccountInfo<'a>,
    house_bump_bytes: [u8; 1],
    escrow_info: AccountInfo<'a>,
    escrow_bump: u8,
    escrow_program_info: AccountInfo<'a>,
    oracle: Pubkey,
    game: &mut Account<state::Game>,
) -> Result<(), Error> {
    msg!("player won");

    **house_info.try_borrow_mut_lamports()? = house_info
        .lamports()
        .checked_sub(game.bet_amount * 2)
        .unwrap();

    **escrow_info.try_borrow_mut_lamports()? = escrow_info
        .lamports()
        .checked_add(game.bet_amount * 2)
        .unwrap();

    let house_seeds = &[HOUSE.as_ref(), oracle.as_ref(), house_bump_bytes.as_ref()];

    escrow::cpi::install_escrow(
        CpiContext::new_with_signer(
            escrow_program_info,
            InstallEscrow {
                initializer: initializer_info.to_account_info(),
                installer: house_info.to_account_info(),
                escrow: escrow_info.to_account_info(),
            },
            &[&house_seeds[..]],
        ),
        escrow_bump,
        game.bet_amount * 2,
    )
    .unwrap();

    game.terminated = true;
    Ok(())
}

pub fn declare_dealer_win(game: &mut state::Game) -> Result<(), Error> {
    msg!("dealer won");
    game.terminated = true;

    Ok(())
}
