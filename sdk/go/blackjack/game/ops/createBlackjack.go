package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/blackjack"
	"triptych.labs/blackjack/game"
)

func CreateBlackjack(oracle solana.PublicKey) *blackjack.Instruction {
	housePda, _ := game.GetHouse(oracle)

	createIx := blackjack.NewCreateBlackjackInstructionBuilder().
		SetHouseAccount(housePda).
		SetOracleAccount(oracle).
		SetSystemProgramAccount(solana.SystemProgramID)

	return createIx.Build()
}
