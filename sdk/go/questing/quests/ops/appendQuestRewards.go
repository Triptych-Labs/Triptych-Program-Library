package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func AppendQuestRewards(rpcClient *rpc.Client, oracle solana.PublicKey, questData questing.Quest) []solana.Instruction {
	questsPda, _ := quests.GetQuests(oracle)
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	fmt.Println(questsData.Quests)
	quest, _ := quests.GetQuest(oracle, questData.Index)

	appendRewardIxs := make([]solana.Instruction, 0)

	rewards := questsData.Rewards
	for _, reward := range rewards {
		appendRewardIx := questing.NewRegisterQuestRewardInstructionBuilder().
			SetOracleAccount(oracle).
			SetQuestAccount(quest).
			SetRentAccount(solana.SysVarRentPubkey).
			SetReward(reward).
			SetRewardMintAccount(reward.MintAddress).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		if e := appendRewardIx.Validate(); e != nil {
			fmt.Println(e.Error())
			panic("...")
		}

		appendRewardIxs = append(appendRewardIxs, appendRewardIx.Build())

	}

	return appendRewardIxs
}

