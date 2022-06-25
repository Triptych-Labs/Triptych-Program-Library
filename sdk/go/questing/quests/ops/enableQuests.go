package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func EnableQuests(oracle solana.PrivateKey) {
	quests, _ := quests.GetQuests(oracle.PublicKey())
	enableQuestIx := questing.NewEnableQuestsInstructionBuilder().
		SetOracleAccount(oracle.PublicKey()).
		SetQuestsAccount(quests).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := enableQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	utils.SendTx(
		"sell",
		append(make([]solana.Instruction, 0), enableQuestIx.Build()),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func RegisterQuestReward(oracle solana.PrivateKey, reward questing.Reward, rewardMint solana.PrivateKey) {
	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	rewardToken, _ := quests.GetRewardToken(oracle.PublicKey(), reward.MintAddress)
	enableQuestIx := questing.NewRegisterQuestRewardInstructionBuilder().
		SetSystemProgramAccount(solana.SystemProgramID).
		SetQuestsAccount(questsPda).
		SetOracleAccount(oracle.PublicKey()).
		SetRentAccount(solana.SysVarRentPubkey).
		SetReward(reward).
		SetRewardMintAccount(reward.MintAddress).
		SetRewardTokenAccount(rewardToken).
		SetTokenProgramAccount(solana.TokenProgramID)

	if e := enableQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	utils.SendTx(
		"sell",
		append(make([]solana.Instruction, 0), enableQuestIx.Build()),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

