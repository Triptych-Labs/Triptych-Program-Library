package utils

import (
	"context"
	"strings"

	ag_binary "github.com/gagliardetto/binary"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/gagliardetto/solana-go"
)

type MetadataResponse struct {
	Name   string           `json:"name"`
	Symbol string           `json:"symbol"`
	Mint   solana.PublicKey `json:"mint"`
}

func GetMetadata(mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			token_metadata.ProgramID.Bytes(),
			mint.Bytes(),
		},
		token_metadata.ProgramID,
	)
	return addr, err
}

func GetMasterEdition(mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			token_metadata.ProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		token_metadata.ProgramID,
	)
	return addr, err
}

func GetMetadataData(rpcClient *rpc.Client, metadataPda solana.PublicKey) *token_metadata.Metadata {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), metadataPda, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data token_metadata.Metadata
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetMetadatasData(rpcClient *rpc.Client, metadataPdas []solana.PublicKey) *[]MetadataResponse {
	metadatas := make([]MetadataResponse, 0)

	response, _ := rpcClient.GetMultipleAccounts(context.TODO(), metadataPdas...)
	if response == nil {
		return nil
	}
	if len(response.Value) == 0 {
		return nil
	}

	for _, bin := range response.Value {
		var data token_metadata.Metadata
		decoder := ag_binary.NewBorshDecoder(bin.Data.GetBinary())
		err := data.UnmarshalWithDecoder(decoder)
		if err != nil {
			panic(err)
		}

		metadatas = append(metadatas, MetadataResponse{
			Name:   strings.Replace(data.Data.Name, "\u0000", "", -1),
			Symbol: strings.Replace(data.Data.Symbol, "\u0000", "", -1),
			Mint:   data.Mint,
		})
	}

	return &metadatas

}
