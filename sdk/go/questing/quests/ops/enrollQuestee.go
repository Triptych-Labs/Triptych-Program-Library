package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func EnrollQuestee(oracle, pixelBallzMint, pixelBallzTokenAddress solana.PublicKey) solana.Instruction {
	questor, _ := quests.GetQuestorAccount(oracle)
	questee, _ := quests.GetQuesteeAccount(pixelBallzMint)

	enrollQuestorIx := questing.NewEnrollQuesteeInstructionBuilder().
		SetInitializerAccount(oracle).
		SetPixelballzMintAccount(pixelBallzMint).
		SetPixelballzTokenAccountAccount(pixelBallzTokenAddress).
		SetQuesteeAccount(questee).
		SetQuestorAccount(questor).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := enrollQuestorIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return enrollQuestorIx.Build()
}
