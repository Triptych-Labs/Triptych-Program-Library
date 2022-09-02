package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/blackjack"
	"triptych.labs/blackjack/game"
)

func RegisterPlayer(rpcClient *rpc.Client, initializer solana.PublicKey) *blackjack.Instruction {
	gamesPda, _ := game.GetGames(initializer)
	gamesData := game.GetGamesData(rpcClient, gamesPda)
	if gamesData != nil {
		return nil
	}

	registerIx := blackjack.NewRegisterPlayerInstructionBuilder().
		SetGamesAccount(gamesPda).
		SetInitializerAccount(initializer).
		SetSystemProgramAccount(solana.SystemProgramID)

	return registerIx.Build()
}

