package swapper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper/swaps/ops"
	"triptych.labs/utils"
)

func InvokeSwap(this js.Value, args []js.Value) interface{} {
	holder := solana.MustPublicKeyFromBase58(args[0].String())
	oracle := solana.MustPublicKeyFromBase58(args[1].String())
	proposalIndexInp := args[2].String()
	amountInp := args[3].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			proposalIndex, _ := strconv.Atoi(proposalIndexInp)
			amount, _ := strconv.ParseFloat(amountInp, 64)

			enrollmentJson, err := invokeSwap(holder, oracle, uint64(proposalIndex), float64(amount))
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			dst := js.Global().Get("Uint8Array").New(len(enrollmentJson))
			js.CopyBytesToJS(dst, enrollmentJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func invokeSwap(holder, oracle solana.PublicKey, swapIndex uint64, amount float64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	if claimIx := ops.InvokeSwap(rpcClient, oracle, holder, swapIndex, amount); claimIx != nil {
		instructions = append(
			instructions,
			claimIx,
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

	fmt.Println(string(txJson))
	return txJson, nil

}
