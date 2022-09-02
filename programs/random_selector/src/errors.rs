use anchor_lang::prelude::*;

#[error_code]
pub enum FlipError {
    #[msg("Unexpected questing state")]
    UnexpectedQuestingState,
    #[msg("Invalid initizalizer")]
    InvalidInitializer,
    #[msg("Is timelocked")]
    IsTimelocked,
    #[msg("Numerical overflow error!")]
    NumericalOverflowError,
    #[msg("Index greater than length!")]
    IndexGreaterThanLength,
    #[msg("Unable to find an unused config line near your random number index")]
    CannotFindUsableConfigLine,
    #[msg("Uuid must be exactly of 6 length")]
    UuidMustBeExactly6Length,
    #[msg("Invalid string")]
    InvalidString,
    #[msg("Suspicious Transaction")]
    SuspiciousTransaction,
    #[msg("Invalid mint")]
    InvalidMint,
    #[msg("Not enough xp")]
    NotEnoughXp,
    #[msg("Invalid conviction")]
    InvalidConviction,
    #[msg("Invalid Completion")]
    InvalidCompletion,
    #[msg("Ratio too big")]
    RatioTooBig,
    #[msg("Proposal Started")]
    ProposalStarted,
    #[msg("Proposal Not Finished")]
    ProposalNotFinished,
    #[msg("Invalid Reward Mint")]
    InvalidRewardMint,
}
