package swaps

import (
	"context"

	ag_binary "github.com/gagliardetto/binary"

	"github.com/gagliardetto/solana-go/rpc"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/swapper"
)

func GetSwapRecorderData(rpcClient *rpc.Client, quests solana.PublicKey) *swapper.SwapRecorder {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), quests, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data swapper.SwapRecorder
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetSwapData(rpcClient *rpc.Client, quests solana.PublicKey) *swapper.SwapProposal {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), quests, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data swapper.SwapProposal
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetSwapsData(rpcClient *rpc.Client, swaps []solana.PublicKey) *[]swapper.SwapProposal {
	response, _ := rpcClient.GetMultipleAccounts(context.TODO(), swaps...)
	if response == nil {
		return nil
	}
	if len(response.Value) == 0 {
		return nil
	}

	swapsData := make([]swapper.SwapProposal, len(swaps))

	for i, swapData := range response.Value {
		var data swapper.SwapProposal
		decoder := ag_binary.NewBorshDecoder(swapData.Data.GetBinary())
		err := data.UnmarshalWithDecoder(decoder)
		if err != nil {
			panic(err)
		}
		swapsData[i] = data
	}

	return &swapsData

}
