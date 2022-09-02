package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"triptych.labs/escrow"
	escrow_ops "triptych.labs/escrow/wallet/ops"
	"triptych.labs/flipper"
	"triptych.labs/flipper/flips/ops"
	"triptych.labs/utils"
)

const LEFT = 2
const RIGHT = 1

var STAKING_MINT solana.PublicKey

func init() {
	flipper.SetProgramID(solana.MustPublicKeyFromBase58("DbceHs8B185bS1tgnXgeKUNFFEgrXw5DaZG8R1cxJYGf"))
	escrow.SetProgramID(solana.MustPublicKeyFromBase58("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ"))
}

func main() {
	op := os.Args[1]
	switch op {
	case "create":
		{
			CreateFlip()
		}
	case "register":
		{
			Register()
		}
	case "flip":
		{
			NewFlip()
		}
	case "fetch":
		{
			Fetch()
		}
	case "collect":
		{
			Collect()
		}
	case "etc":
		{
			Etc()
		}
	}
}

func Register() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	initEscrowIx := escrow_ops.InitializeEscrow(oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), initEscrowIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func Fetch() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	flipStats := ops.GetStatistics(rpcClient, oracle.PublicKey())
	flipDataJson, _ := json.MarshalIndent(flipStats, "", "  ")
	fmt.Println(string(flipDataJson))
}

func Etc() {
	t := time.Now().UTC()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(rounded.Unix())

	cost := 0.00251952
	fmt.Println(int(cost * float64(solana.LAMPORTS_PER_SOL)))
}

func CreateFlip() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	createIx := ops.CreateFlip(oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), createIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}

func NewFlip() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	var selection = 0
	var times = 0

	selection, err = strconv.Atoi(os.Args[2])
	if err != nil {
		selection = 0
	}
	times, err = strconv.Atoi(os.Args[3])
	if err != nil {
		times = 0
	}

	instructions := make([]solana.Instruction, 0)
	for i := range make([]int, times) {
		if i%5 == 0 && i != 0 {

			utils.SendTx(
				"create",
				instructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)
			instructions = make([]solana.Instruction, 0)
		}

		createIx := ops.NewFlip(oracle.PublicKey(), oracle.PublicKey(), uint64(0.1*float64(solana.LAMPORTS_PER_SOL)), uint8(selection), "wallet")

		instructions = append(instructions, createIx)
	}
	utils.SendTx(
		"create",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}

func Collect() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	amount := uint64(float64(23.8) * float64(solana.LAMPORTS_PER_SOL))
	collectIx := ops.WithdrawHouse(oracle.PublicKey(), amount)

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), collectIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}

