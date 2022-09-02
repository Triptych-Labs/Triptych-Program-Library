package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
)

func InitializeEscrow(initializer solana.PublicKey) *escrow.Instruction {
	escrowPda, _ := wallet.GetEscrow(initializer)

	initIx := escrow.NewInitializeEscrowInstructionBuilder().
		SetEscrowAccount(escrowPda).
		SetInitializerAccount(initializer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return initIx.Build()
}
