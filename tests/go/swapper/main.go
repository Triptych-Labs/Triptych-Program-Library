package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go/programs/system"

	"github.com/btcsuite/btcutil/base58"

	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
	"triptych.labs/swapper/swaps/ops"
	swaps_ops "triptych.labs/swapper/swaps/ops"
	"triptych.labs/utils"

	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

const LEFT = 2
const RIGHT = 1

var STAKING_MINT solana.PublicKey

func init() {
	swapper.SetProgramID(solana.MustPublicKeyFromBase58("EocXeJ7KvMZD9k1mVWVHP6KWpRhsLP666wtahSBEMawr"))
}

func main() {
	fromMint := solana.MustPublicKeyFromBase58("C2Y5XRNLEu7CipiVocMrpKzaivXVS14BEFFLPD7wffw6")
	toMint := solana.MustPublicKeyFromBase58("3nNeqjQ724xe4E8kKVGP5t8NFt4HSLksvN5NwMFphCSM")
	_ = fromMint
	_ = toMint
	/*
		  toMint, err := solana.PrivateKeyFromSolanaKeygenFile("./mint.key")
			if err != nil {
				panic(err)
			}
	*/

	// CreateNTokenAccountsOfMint(solana.MustPublicKeyFromBase58("6c5EBgbPnpdZgKhXW4uTtcYojXqVNnVQbS1cdCHo8Zmu"), 2)
	// enableSwaps()

	// createTokenAndMetadata(toMint)
	createSwap(fromMint, toMint)
	// makeSwap()

	// verifyMetadata()
	// getSwaps()

	hash := sha256.Sum256([]byte("account:QuestAccount"))
	encoded := base58.Encode(hash[:8])
	fmt.Println(string(encoded))

	{
		oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
		if err != nil {
			panic(err)
		}

		for i := range make([]int, 3) {
			swapPda, _ := swaps.GetSwap(oracle.PublicKey(), uint64(i))
			fmt.Println(i, swapPda)
		}
	}

}

func createTokenAndMetadata(toMint solana.PrivateKey) {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	data := token_metadata.DataV2{
		Name:                 "Not ApeCoin",
		Symbol:               "NBA",
		Uri:                  "",
		SellerFeeBasisPoints: 0,
		Creators:             nil,
		Collection:           nil,
		Uses:                 nil,
	}
	metadata := token_metadata.CreateMetadataAccountArgsV2{
		Data: token_metadata.DataV2{
			Name:                 "Not ApeCoin",
			Symbol:               "NBA",
			Uri:                  "",
			SellerFeeBasisPoints: 0,
			Creators:             nil,
			Collection:           nil,
			Uses:                 nil,
		},
		IsMutable: true,
	}
	metadataPda, _ := utils.GetMetadata(toMint.PublicKey())

	instructions := make([]solana.Instruction, 0)

	min, err := rpcClient.GetMinimumBalanceForRentExemption(context.TODO(), token.MINT_SIZE, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	var ix solana.Instruction
	if utils.GetMetadataData(rpcClient, metadataPda) == nil {
		instructions = append(instructions,
			system.NewCreateAccountInstructionBuilder().
				SetOwner(token.ProgramID).
				SetNewAccount(toMint.PublicKey()).
				SetSpace(token.MINT_SIZE).
				SetFundingAccount(oracle.PublicKey()).
				SetLamports(min).
				Build(),

			token.NewInitializeMint2InstructionBuilder().
				SetMintAccount(toMint.PublicKey()).
				SetDecimals(3).
				SetMintAuthority(oracle.PublicKey()).
				SetFreezeAuthority(oracle.PublicKey()).
				Build(),

			atok.NewCreateInstructionBuilder().
				SetPayer(oracle.PublicKey()).
				SetWallet(oracle.PublicKey()).
				SetMint(toMint.PublicKey()).
				Build(),
		)
		ix = token_metadata.NewCreateMetadataAccountV2Instruction(metadata, metadataPda, toMint.PublicKey(), oracle.PublicKey(), oracle.PublicKey(), oracle.PublicKey(), solana.SystemProgramID, solana.SysVarRentPubkey).Build()
	} else {
		ix = token_metadata.NewUpdateMetadataAccountV2Instruction(token_metadata.UpdateMetadataAccountArgsV2{
			Data: &data,
		}, metadataPda, oracle.PublicKey()).Build()
	}
	instructions = append(instructions, ix)

	utils.SendTx(
		"create token",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle, toMint),
		oracle.PublicKey(),
	)
}

func verifyMetadata() {
	rpcClient := rpc.New(utils.NETWORK)

	fromMint := solana.MustPublicKeyFromBase58("FsJTaKL31xeEPgPf8yFysCTZYj8v4B8Vvc8FSNKw3uLX")
	toMint := solana.MustPublicKeyFromBase58("CTM8npagWrtdi85aYix3kpD23yKdboPFMXk9fPWMBoD7")

	mints := []solana.PublicKey{fromMint, toMint}

	metadataData, _ := utils.GetTokensMetadataData(rpcClient, mints)
	metadataJson, _ := json.MarshalIndent(metadataData, "", "  ")
	fmt.Println(string(metadataJson))

}

func getSwaps() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	swaps := ops.GetSwaps(rpcClient, oracle.PublicKey())
	swapsJson, _ := json.MarshalIndent(swaps, "", "  ")
	fmt.Println(string(swapsJson))
}

func CreateNTokenAccountsOfMint(mint solana.PublicKey, amount int) {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	tokenAccounts := make([]string, amount)
	var instructions []solana.Instruction
	for i := range tokenAccounts {
		wallet := solana.NewWallet()
		ata, _ := utils.GetTokenWallet(wallet.PublicKey(), mint)
		tokenAccounts[i] = ata.String()

		instructions = append(instructions,
			atok.NewCreateInstructionBuilder().
				SetPayer(oracle.PublicKey()).
				SetWallet(wallet.PublicKey()).
				SetMint(mint).
				Build(),
		)

	}
	utils.SendTx(
		"list",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
	fmt.Println(tokenAccounts)
}

func formatAsJson(data interface{}) {
	dataJson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dataJson))
}

func enableSwaps() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	ixs := make([]solana.Instruction, 0)
	registerIx := swaps_ops.RegisterSwapRecorder(rpcClient, oracle.PublicKey())
	ixs = append(ixs, registerIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func createSwap(fromMint, toMint solana.PublicKey) {
	ixs := make([]solana.Instruction, 0)
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	proposalIx := swaps_ops.ProposeSwapRecord(rpcClient, oracle.PublicKey(), fromMint, toMint, 1, 1)

	ixs = append(ixs, proposalIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func makeSwap() {
	ixs := make([]solana.Instruction, 0)
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	swapIx := swaps_ops.InvokeSwap(rpcClient, oracle.PublicKey(), oracle.PublicKey(), 1, 1)

	ixs = append(ixs, swapIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}
