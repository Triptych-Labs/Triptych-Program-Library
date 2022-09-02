package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	escrow_program "triptych.labs/escrow"
	flipper_program "triptych.labs/flipper"
	"triptych.labs/wasm/v2/integrations"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	flipper_program.SetProgramID(solana.MustPublicKeyFromBase58("DbceHs8B185bS1tgnXgeKUNFFEgrXw5DaZG8R1cxJYGf"))
	escrow_program.SetProgramID(solana.MustPublicKeyFromBase58("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ"))

	/*
	   getSwaps := js.FuncOf(integrations.GetSwaps)
	   defer getSwaps.Release()
	   global.Set("get_swaps", getSwaps)
	*/

	invokeFlip := js.FuncOf(integrations.InvokeFlip)
	defer invokeFlip.Release()
	global.Set("invoke_flip", invokeFlip)

	// getStatistics
	getStatistics := js.FuncOf(integrations.GetStatistics)
	defer getStatistics.Release()
	global.Set("get_statistics", getStatistics)

	<-done
}
