package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
)

func DepositEscrow(initializer solana.PublicKey, amount uint64) *escrow.Instruction {
	escrowPda, escrowPdaBump := wallet.GetEscrow(initializer)

	initIx := escrow.NewDepositEscrowInstructionBuilder().
		SetAmount(amount).
		SetEscrowAccount(escrowPda).
		SetEscrowBump(escrowPdaBump).
		SetInitializerAccount(initializer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return initIx.Build()
}
