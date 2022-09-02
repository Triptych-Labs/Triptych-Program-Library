package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"triptych.labs/blackjack"
	"triptych.labs/blackjack/game"
	"triptych.labs/blackjack/game/ops"
	"triptych.labs/escrow"
	"triptych.labs/utils"
)

func init() {
	blackjack.SetProgramID(solana.MustPublicKeyFromBase58("4D3g6DHPDiE3gD9G2yfRnEehkQKJwfh5VkExDgeHvhxr"))
	escrow.SetProgramID(solana.MustPublicKeyFromBase58("2wbpcaSSP3H6uaqMhgQAjtwWCsLqoBQMaJ1b3MGe5WFJ"))
}

func main() {
	op := os.Args[1]
	switch op {
	case "create":
		{
			CreateGame()
		}
	case "register":
		{
			Register()
		}
	case "newGame":
		{
			NewGame()
		}
	case "hit":
		{
			Hit()
		}
	case "stand":
		{
			Stand()
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

func Hit() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	gameIndex := uint64(0)
	gamesPda, _ := game.GetGames(oracle.PublicKey())
	if gamesData := game.GetGamesData(rpcClient, gamesPda); gamesData != nil {
		gameIndex = gamesData.Games - 1
	}
	instructions := make([]solana.Instruction, 0)

	hitIx := ops.AdvanceGame(rpcClient, oracle.PublicKey(), oracle.PublicKey(), "hit", gameIndex)
	instructions = append(instructions, hitIx)

	utils.SendTx(
		"create",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

	{
		gamePda, _ := game.GetGame(oracle.PublicKey(), gameIndex)
		gameData := game.GetGameData(rpcClient, gamePda)
		j, _ := json.MarshalIndent(gameData, "", "  ")
		fmt.Println(string(j))
	}
}

func Stand() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	gameIndex := uint64(0)
	gamesPda, _ := game.GetGames(oracle.PublicKey())
	if gamesData := game.GetGamesData(rpcClient, gamesPda); gamesData != nil {
		gameIndex = gamesData.Games - 1
	}
	instructions := make([]solana.Instruction, 0)

	standIx := ops.AdvanceGame(rpcClient, oracle.PublicKey(), oracle.PublicKey(), "stand", gameIndex)
	instructions = append(instructions, standIx)

	utils.SendTx(
		"create",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

	{
		gamePda, _ := game.GetGame(oracle.PublicKey(), gameIndex)
		gameData := game.GetGameData(rpcClient, gamePda)
		j, _ := json.MarshalIndent(gameData, "", "  ")
		fmt.Println(string(j))
	}
}

func Register() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	registerIx := ops.RegisterPlayer(rpcClient, oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), registerIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func Fetch() {
	/*
		rpcClient := rpc.New(utils.NETWORK)
		oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
		if err != nil {
			panic(err)
		}
		flipStats := ops.GetStatistics(rpcClient, oracle.PublicKey())
		flipDataJson, _ := json.MarshalIndent(flipStats, "", "  ")
		fmt.Println(string(flipDataJson))
	*/
}

func Etc() {
	t := time.Now().UTC()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(rounded.Unix())

	cost := 0.00251952
	fmt.Println(int(cost * float64(solana.LAMPORTS_PER_SOL)))
}

func CreateGame() {
	/*
		// rpcClient := rpc.New(utils.NETWORK)
	*/
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	createIx := ops.CreateBlackjack(oracle.PublicKey())

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), createIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

}

func NewGame() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	newGameIx := ops.NewGame(
		rpcClient,
		oracle.PublicKey(),
		oracle.PublicKey(),
		uint64(0.1*float64(solana.LAMPORTS_PER_SOL)),
		"wallet",
	)

	utils.SendTx(
		"create",
		append(make([]solana.Instruction, 0), newGameIx),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
}

func Collect() {
	/*
		// rpcClient := rpc.New(utils.NETWORK)
		oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
		if err != nil {
			panic(err)
		}

		amount := 2 * solana.LAMPORTS_PER_SOL
		collectIx := ops.WithdrawHouse(oracle.PublicKey(), amount)

		utils.SendTx(
			"create",
			append(make([]solana.Instruction, 0), collectIx),
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	*/

}
