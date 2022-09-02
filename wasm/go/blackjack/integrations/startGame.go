package integrations

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/blackjack/game/ops"
	"triptych.labs/escrow/wallet"
	escrow_ops "triptych.labs/escrow/wallet/ops"
	"triptych.labs/utils"
)

func StartGame(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	holder := solana.MustPublicKeyFromBase58(args[1].String())
	amountInp := args[2].String()
	operator := args[3].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			amount, _ := strconv.Atoi(amountInp)

			startGameTxJson, err := startGame(oracle, holder, uint64(amount), operator)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			dst := js.Global().Get("Uint8Array").New(len(startGameTxJson))
			js.CopyBytesToJS(dst, startGameTxJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func startGame(oracle, holder solana.PublicKey, amount uint64, operator string) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)
	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	escrow, _ := wallet.GetEscrow(holder)
	if escrowData := wallet.GetEscrowData(rpcClient, escrow); escrowData == nil {
		instructions = append(
			instructions,
			escrow_ops.InitializeEscrow(holder),
		)
	}

	if registrationIx := ops.RegisterPlayer(rpcClient, holder); registrationIx != nil {
		instructions = append(
			instructions,
			registrationIx,
		)
	}

	fmt.Println("args", "oracle", oracle, "holder", holder)
	newGameIx := ops.NewGame(
		rpcClient,
		oracle,
		holder,
		amount,
		operator,
	)
	if newGameIx != nil {
		instructions = append(
			instructions,
			newGameIx,
		)
	}

	if len(instructions) > 0 {
		txBuilder := solana.NewTransactionBuilder()
		for _, ix := range instructions {
			txBuilder = txBuilder.AddInstruction(ix)
		}
		txB, _ := txBuilder.Build()
		txJson, _ = json.MarshalIndent(txB, "", "  ")

	}

	return txJson, nil

}
