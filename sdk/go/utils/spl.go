package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

type TokenListMeta struct {
	Address solana.PublicKey `json:"address"`
	Symbol  string           `json:"symbol"`
	Name    string           `json:"name"`
}

type TokenList struct {
	Tokens []TokenListMeta `json:"tokens"`
}

func GetTokenWallet(wallet solana.PublicKey, mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			wallet.Bytes(),
			solana.TokenProgramID.Bytes(),
			mint.Bytes(),
		},
		solana.SPLAssociatedTokenAccountProgramID,
	)
	return addr, err
}

func FetchTokenMeta() []TokenListMeta {
	var tokenList TokenList
	tokenListUrl := fmt.Sprint(CDN + "/solana.tokenlist.json")
	res, err := http.DefaultClient.Get(tokenListUrl)
	if err != nil {
		return tokenList.Tokens
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tokenList.Tokens
	}

	err = json.Unmarshal(data, &tokenList)
	if err != nil {
		return tokenList.Tokens
	}

	return tokenList.Tokens

}

func ConvertUiAmountToAmount(uiAmount float64, decimals uint8) uint64 {
	return uint64(uiAmount * math.Pow10(int(decimals)))
}

func ConvertAmountToUiAmount(amount uint64, decimals uint8) float64 {
	return float64(amount) / math.Pow10(int(decimals))
}

func GetTokenMintData(rpcClient *rpc.Client, tokenMint solana.PublicKey) *token.Mint {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), tokenMint, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data token.Mint
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetTokenMintsData(rpcClient *rpc.Client, tokenMints []solana.PublicKey) map[solana.PublicKey]token.Mint {
	mints := make(map[solana.PublicKey]token.Mint)

	response, _ := rpcClient.GetMultipleAccounts(context.TODO(), tokenMints...)
	if response == nil {
		return nil
	}
	if len(response.Value) == 0 {
		return nil
	}

	for i, bin := range response.Value {
		var data token.Mint
		decoder := ag_binary.NewBorshDecoder(bin.Data.GetBinary())
		err := data.UnmarshalWithDecoder(decoder)
		if err != nil {
			panic(err)
		}

		mints[tokenMints[i]] = data
	}

	return mints

}

