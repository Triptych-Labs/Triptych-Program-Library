package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/flipper"
	"triptych.labs/flipper/flips"
)

func GetStatistics(rpcClient *rpc.Client, oracle solana.PublicKey) flipper.Flip {
	flip, _, _ := flips.GetFlip(oracle)
	flipData := flips.GetFlipData(rpcClient, flip)
	if flipData == nil {
		return flipper.Flip{}
	}

	return *flipData
}
