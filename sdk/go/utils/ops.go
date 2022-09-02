package utils

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

type TokensMetadataResponse struct {
	TokenMintMeta token.Mint       `json:"tokenMintMeta"`
	TokenMetadata MetadataResponse `json:"tokenMetadata"`
}

func FindTokenResponse(tokenMint solana.PublicKey, mintsData map[solana.PublicKey]token.Mint, metadatasData []MetadataResponse) TokensMetadataResponse {
	mintMeta := mintsData[tokenMint]
	metadata := func() MetadataResponse {
		for _, metadata := range metadatasData {
			if metadata.Mint == tokenMint {
				return metadata
			}
		}
		panic("unforeseen")
	}()

	return TokensMetadataResponse{
		TokenMintMeta: mintMeta,
		TokenMetadata: metadata,
	}
}

func GetTokensMetadataData(rpcClient *rpc.Client, tokenMints []solana.PublicKey) (map[solana.PublicKey]token.Mint, []MetadataResponse) {
	tokenMetadatas := make([]solana.PublicKey, len(tokenMints))
	for i, mint := range tokenMints {
		mintMetadata, _ := GetMetadata(mint)
		tokenMetadatas[i] = mintMetadata
	}

	mintsData := GetTokenMintsData(rpcClient, tokenMints)
	metadatasData := []MetadataResponse{}
	metadatasDataResponse := GetMetadatasData(rpcClient, tokenMetadatas)
	if metadatasDataResponse != nil {
		metadatasData = *metadatasDataResponse
	}

	return mintsData, metadatasData
}

