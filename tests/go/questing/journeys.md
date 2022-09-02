# e2e tests

## [] Create N quests reusing Staking Reward Mint
```
stakingMint := new Keypair()
_ = initializeStakingReward({withStakingMint: stakingMint})

for _ in range(N):
    stakingReward := {mint: stakingMint.PublicKey()}
    _ = createQuest({withStakingReward: stakingReward})
```


## [] Create N quests reusing Quest Reward Mint
```
questRewardMint := new Keypair()
_ = initializeQuestReward({withQuestRewardMint: questRewardMint})

for _ in range(N):
    questReward := {mint: questRewardMint.PublicKey()}
    _ = createQuest({withQuestReward: questReward})
```
