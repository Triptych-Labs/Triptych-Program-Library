package integrations

import (
	"encoding/json"
	"strconv"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/escrow/wallet"
	escrow_ops "triptych.labs/escrow/wallet/ops"
	"triptych.labs/flipper/flips/ops"
	"triptych.labs/utils"
)

func InvokeFlip(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	holder := solana.MustPublicKeyFromBase58(args[1].String())
	amountInp := args[2].String()
	selectionInp := args[3].String()
	operator := args[4].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			amount, _ := strconv.Atoi(amountInp)
			selection, _ := strconv.Atoi(selectionInp)

			flipTxJson, err := invokeFlip(oracle, holder, uint64(amount), uint8(selection), operator)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			dst := js.Global().Get("Uint8Array").New(len(flipTxJson))
			js.CopyBytesToJS(dst, flipTxJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func invokeFlip(oracle, holder solana.PublicKey, amount uint64, selection uint8, operator string) ([]byte, error) {
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

	if flipIx := ops.NewFlip(oracle, holder, amount, selection, operator); flipIx != nil {
		instructions = append(
			instructions,
			flipIx,
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
