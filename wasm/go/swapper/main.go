package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	swapper_program "triptych.labs/swapper"
	"triptych.labs/wasm/v2/integrations/swapper"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	swapper_program.SetProgramID(solana.MustPublicKeyFromBase58("EocXeJ7KvMZD9k1mVWVHP6KWpRhsLP666wtahSBEMawr"))

	getSwaps := js.FuncOf(swapper.GetSwaps)
	defer getSwaps.Release()
	global.Set("get_swaps", getSwaps)

	invokeSwap := js.FuncOf(swapper.InvokeSwap)
	defer invokeSwap.Release()
	global.Set("invoke_swap", invokeSwap)

	<-done
}
