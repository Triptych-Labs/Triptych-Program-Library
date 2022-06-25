package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func EnrollQuestor(oracle solana.PublicKey) solana.Instruction {
	questor, _ := quests.GetQuestorAccount(oracle)

	enrollQuestorIx := questing.NewEnrollQuestorInstructionBuilder().
		SetInitializerAccount(oracle).
		SetQuestorAccount(questor).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := enrollQuestorIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return enrollQuestorIx.Build()
}
