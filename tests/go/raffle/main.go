package main

import (
	"os"
	"strconv"

	"github.com/gagliardetto/solana-go"

	"triptych.labs/flipper"
	"triptych.labs/flipper/flips/ops"
	"triptych.labs/utils"
)

const LEFT = 2
const RIGHT = 1

var STAKING_MINT solana.PublicKey

func init() {
	flipper.SetProgramID(solana.MustPublicKeyFromBase58("DbceHs8B185bS1tgnXgeKUNFFEgrXw5DaZG8R1cxJYGf"))
}

func main() {
	op := os.Args[1]
	switch op {
	case "create":
		{
			CreateFlip()
		}
	case "flip":
		{
			NewFlip()
		}
	case "collect":
		{
			Collect()
		}
	}
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

	var times = 0
	times, err = strconv.Atoi(os.Args[2])
	if err != nil {
		times = 0
	}

	{
		initEscrowIx := ops.InitializeEscrow(oracle.PublicKey())

		utils.SendTx(
			"create",
			append(make([]solana.Instruction, 0), initEscrowIx),
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
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

		createIx := ops.NewFlip(oracle.PublicKey(), oracle.PublicKey(), uint64(0.2*float64(solana.LAMPORTS_PER_SOL)), 0)

		instructions = append(instructions, createIx)
	}

}

func Collect() {
	// rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	collectIx := ops.DrainEscrow(oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), collectIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}
