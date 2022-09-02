package integrations

import (
	"encoding/json"
	"strconv"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/escrow/wallet"
	"triptych.labs/escrow/wallet/ops"
	"triptych.labs/utils"
)

func DepositEscrow(this js.Value, args []js.Value) interface{} {
	holder := solana.MustPublicKeyFromBase58(args[0].String())
	amountInp := args[1].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			amount, _ := strconv.Atoi(amountInp)

			flipTxJson, err := depositEscrow(holder, uint64(amount))
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

func depositEscrow(holder solana.PublicKey, amount uint64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)
	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	escrow, _ := wallet.GetEscrow(holder)
	if escrowData := wallet.GetEscrowData(rpcClient, escrow); escrowData == nil {
		instructions = append(
			instructions,
			ops.InitializeEscrow(holder),
		)
	}

	if depositIx := ops.DepositEscrow(holder, amount); depositIx != nil {
		instructions = append(
			instructions,
			depositIx,
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
