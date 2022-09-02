package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/flipper"
	"triptych.labs/flipper/flips"
)

func WithdrawHouse(oracle solana.PublicKey, amount uint64) *flipper.Instruction {
	house, houseBump := flips.GetHouse(oracle)
	ix := flipper.NewWithdrawHouseInstructionBuilder().
		SetAmount(amount).
		SetHouseAccount(house).
		SetHouseBump(houseBump).
		SetOracleAccount(oracle)

	return ix.Build()
}

