package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
)

func RegisterSwapRecorder(rpcClient *rpc.Client, oracle solana.PublicKey) solana.Instruction {
	swapRecorder, _ := swaps.GetSwapRecorder(oracle)

	registerIx := swapper.NewRegisterSwapRecorderInstructionBuilder().
		SetOracleAccount(oracle).
		SetSwapRecorderAccount(swapRecorder).
		SetSystemProgramAccount(solana.SystemProgramID)

	return registerIx.Build()
}
