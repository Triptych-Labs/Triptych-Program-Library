package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func RegisterQuestRecorder(rpcClient *rpc.Client, initializer, questPda solana.PublicKey) *questing.Instruction {
	questRecorder, _ := quests.GetQuestRecorder(questPda, initializer)
	if questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder); questRecorderData != nil {
		return nil
	}

	createQuestRecorderIx := questing.NewRegisterQuestRecorderInstructionBuilder().
		SetInitializerAccount(initializer).
		SetQuestAccount(questPda).
		SetQuestRecorderAccount(questRecorder).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := createQuestRecorderIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return createQuestRecorderIx.Build()
}
