package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
)

func DrainEscrow(initializer solana.PublicKey, amount uint64) *escrow.Instruction {
	escrowPda, escrowPdaBump := wallet.GetEscrow(initializer)

	drainIx := escrow.NewDrainEscrowInstructionBuilder().
		SetAmount(amount).
		SetCallerProgramAccount(solana.SystemProgramID).
		SetCollectorAccount(initializer).
		SetEscrowAccount(escrowPda).
		SetEscrowBump(escrowPdaBump).
		SetInitializerAccount(initializer)

	return drainIx.Build()
}

