package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/flipper"
	"triptych.labs/flipper/flips"
)

func CreateFlip(oracle solana.PublicKey) *flipper.Instruction {
	house, _ := flips.GetHouse(oracle)
	ix := flipper.NewCreateFlipInstructionBuilder().
		SetHouseAccount(house).
		SetOracleAccount(oracle).
		SetSystemProgramAccount(solana.SystemProgramID)

	return ix.Build()
}
