package wallet

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/escrow"
)

func GetEscrowData(rpcClient *rpc.Client, escrowPda solana.PublicKey) *escrow.Escrow {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), escrowPda, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data escrow.Escrow
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}
