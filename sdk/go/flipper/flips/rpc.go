package flips

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/flipper"
)

func GetEscrowData(rpcClient *rpc.Client, escrow solana.PublicKey) *flipper.Escrow {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), escrow, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data flipper.Escrow
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetFlipData(rpcClient *rpc.Client, flip solana.PublicKey) *flipper.Flip {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), flip, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data flipper.Flip
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}
