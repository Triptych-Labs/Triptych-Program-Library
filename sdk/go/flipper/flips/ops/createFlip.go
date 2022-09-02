package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
	"triptych.labs/flipper"
	"triptych.labs/flipper/flips"
)

func NewFlip(oracle, initializer solana.PublicKey, amount uint64, selection uint8, operator string) *flipper.Instruction {
	flip, flipBump, dailyEpoch := flips.GetFlip(oracle)
	house, houseBump := flips.GetHouse(oracle)
	escrowPda, escrowPdaBump := wallet.GetEscrow(initializer)

	fees := solana.MustPublicKeyFromBase58("G8UyAzcphHUE4ZCtH8YQCFATehWEhLN53Pf1aWi2SPCM")

	ix := flipper.NewNewFlipInstructionBuilder().
		SetAmount(amount).
		SetEscrowAccount(escrowPda).
		SetEscrowBump(escrowPdaBump).
		SetEscrowProgramAccount(escrow.ProgramID).
		SetDailyEpoch(dailyEpoch).
		SetFeesAccount(fees).
		SetFlipAccount(flip).
		SetFlipBump(flipBump).
		SetFlipperProgramAccount(flipper.ProgramID).
		SetHouseAccount(house).
		SetHouseBump(houseBump).
		SetInitializerAccount(initializer).
		SetOracleAccount(oracle).
		SetOperator(operator).
		SetSelection(selection).
		SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")).
		SetSystemProgramAccount(solana.SystemProgramID)

	err := ix.Validate()
	if err != nil {
		panic(err)
	}

	return ix.Build()
}
