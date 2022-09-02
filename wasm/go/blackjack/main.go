package main

import (
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	blackjack_program "triptych.labs/blackjack"
	escrow_program "triptych.labs/escrow"
	"triptych.labs/wasm/v2/integrations"
)

func main() {
	global := js.Global()
	done := make(chan struct{})
	blackjack_program.SetProgramID(solana.MustPublicKeyFromBase58("4D3g6DHPDiE3gD9G2yfRnEehkQKJwfh5VkExDgeHvhxr"))
	escrow_program.SetProgramID(solana.MustPublicKeyFromBase58("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ"))

	startGame := js.FuncOf(integrations.StartGame)
	defer startGame.Release()
	global.Set("start_game", startGame)

	<-done
}
