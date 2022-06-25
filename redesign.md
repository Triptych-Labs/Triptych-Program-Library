# Redesign

## Reward Tokens

* Must be reused across all quests.
* Honouring initialised probabilities.

### Process

There is a disparate process to create a reward token and its associated chance
probability.

For instance after deploying the contract, reward tokens must be initialized
with:

* Chance/Probability (Out of 100)
* Amount

### Association

During Quest creation, reward token mints are required as an input and will
be validated for the sum probabilities equaling 100.

### Federation

Upon Quest Ending for a participant, the contract will RNG an unsigned 64-bit
integer and modulate over 100 to yield a percentile used to represent the
threshold of probabilities for the reward tokens.

## Notes

The reward tokens need a shared authority that is not the hosting _Quest_ as such
_Quest_ identity will vary, such shared identity can be appropriated via
`quests` PDA.

