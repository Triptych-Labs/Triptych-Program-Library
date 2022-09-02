package swapper

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper/swaps/ops"
	"triptych.labs/utils"
)

func GetSwaps(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			enrollmentJson, err := getSwaps(oracle)
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

func getSwaps(oracle solana.PublicKey) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)
	txJson := []byte("[]")

	swaps := ops.GetSwaps(rpcClient, oracle)

	txJson, _ = json.MarshalIndent(swaps, "", "  ")

	fmt.Println(string(txJson))
	return txJson, nil

}
