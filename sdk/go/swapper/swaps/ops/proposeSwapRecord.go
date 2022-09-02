package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
)

func ProposeSwapRecord(rpcClient *rpc.Client, oracle, fromMint, toMint solana.PublicKey, per, exch uint64) solana.Instruction {
	swapRecorder, _ := swaps.GetSwapRecorder(oracle)
	swapRecorderData := swaps.GetSwapRecorderData(rpcClient, swapRecorder)

	swapPda, _ := swaps.GetSwap(oracle, swapRecorderData.Proposals)
	swapPool, _ := swaps.GetSwapPool(toMint, swapRecorder)

	proposalIx := swapper.NewProposeSwapRecordInstructionBuilder().
		SetExchange(exch).
		SetFromMintAccount(fromMint).
		SetOracleAccount(oracle).
		SetPer(per).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSplits([]swapper.Split{{TokenAddress: solana.PublicKey{}, OpCode: 1, Share: 100}}).
		SetSwapAccount(swapPda).
		SetSwapPoolAccount(swapPool).
		SetSwapRecorderAccount(swapRecorder).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetToMintAccount(toMint).
		SetTokenProgramAccount(solana.TokenProgramID)

	if e := proposalIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return proposalIx.Build()
}
