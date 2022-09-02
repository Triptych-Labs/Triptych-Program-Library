package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
	"triptych.labs/utils"
)

type mintsMeta struct {
	From utils.TokensMetadataResponse `json:"from"`
	To   utils.TokensMetadataResponse `json:"to"`
}

type SwapResponse struct {
	MintsMeta mintsMeta            `json:"mintsMeta"`
	SwapMeta  swapper.SwapProposal `json:"swapMeta"`
}

type GetSwapsResponse []SwapResponse

func GetSwaps(rpcClient *rpc.Client, oracle solana.PublicKey) GetSwapsResponse {
	swapRecorder, _ := swaps.GetSwapRecorder(oracle)

	swapRecorderData := swaps.GetSwapRecorderData(rpcClient, swapRecorder)

	swapPdas := make([]solana.PublicKey, swapRecorderData.Proposals)
	for i := range swapPdas {
		swapPda, _ := swaps.GetSwap(oracle, uint64(i))
		swapPdas[i] = swapPda
	}

	swapsData := swaps.GetSwapsData(rpcClient, swapPdas)

	mints := make([]solana.PublicKey, 0)
	for _, swapData := range *swapsData {
		mints = append(mints, swapData.FromMint, swapData.ToMint)
	}

	mintsData, metadatasData := utils.GetTokensMetadataData(rpcClient, mints)

	swapsResponse := make(GetSwapsResponse, 0)
	for _, swapData := range *swapsData {
		response := SwapResponse{
			MintsMeta: mintsMeta{
				From: utils.FindTokenResponse(swapData.FromMint, mintsData, metadatasData),
				To:   utils.FindTokenResponse(swapData.ToMint, mintsData, metadatasData),
			},
			SwapMeta: swapData,
		}

		swapsResponse = append(swapsResponse, response)
	}

	return swapsResponse
}

