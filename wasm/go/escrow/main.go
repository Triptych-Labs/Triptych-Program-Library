package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	escrow_program "triptych.labs/escrow"
	"triptych.labs/wasm/v2/integrations"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	escrow_program.SetProgramID(solana.MustPublicKeyFromBase58("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ"))

	getEscrow := js.FuncOf(integrations.GetEscrow)
	defer getEscrow.Release()
	global.Set("get_escrow", getEscrow)

	drainEscrow := js.FuncOf(integrations.DrainEscrow)
	defer drainEscrow.Release()
	global.Set("drain_escrow", drainEscrow)

	depositEscrow := js.FuncOf(integrations.DepositEscrow)
	defer depositEscrow.Release()
	global.Set("deposit_escrow", depositEscrow)

	<-done
}
