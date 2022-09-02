package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func CreateQuest(rpcClient *rpc.Client, oracle solana.PublicKey, questData questing.Quest) (solana.Instruction, uint64) {
	questsPda, _ := quests.GetQuests(oracle)
	quest, _ := quests.GetQuest(oracle, questData.Index)
	createQuestIx := questing.NewCreateQuestInstructionBuilder().
		SetDuration(questData.Duration).
		SetName(questData.Name).
		SetOracleAccount(oracle).
		SetQuestAccount(quest).
		SetQuestIndex(questData.Index).
		SetQuestsAccount(questsPda).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetWlCandyMachines(questData.WlCandyMachines).
		SetXp(questData.Xp).
		SetEnabled(true).
		SetRequiredLevel(questData.RequiredLevel).
		SetRewards(questData.Rewards)

	if questData.Tender != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		createQuestIx.SetTender(*questData.Tender)
		createQuestIx.SetTenderSplits(*questData.TenderSplits)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}

	if questData.StakingConfig != nil {
		createQuestIx.SetStakingConfig(*questData.StakingConfig)
	}

	if questData.PairsConfig != nil {
		createQuestIx.SetPairsConfig(*questData.PairsConfig)
	}

	if e := createQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return createQuestIx.Build(), questData.Index
}

