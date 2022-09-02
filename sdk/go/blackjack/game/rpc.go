package game

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/blackjack"
)

func GetGamesData(rpcClient *rpc.Client, games solana.PublicKey) *blackjack.Games {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), games, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data blackjack.Games
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetGameData(rpcClient *rpc.Client, game solana.PublicKey) *blackjack.Game {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), game, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data blackjack.Game
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}
