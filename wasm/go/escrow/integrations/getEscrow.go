package integrations

import (
	"encoding/json"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/escrow/wallet"
	"triptych.labs/utils"
)

func GetEscrow(this js.Value, args []js.Value) interface{} {
	initializer := solana.MustPublicKeyFromBase58(args[0].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Failed to fetch Escrow Metadata")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)

			enrollmentJson := getEscrow(initializer)

			dst := js.Global().Get("Uint8Array").New(len(enrollmentJson))
			js.CopyBytesToJS(dst, enrollmentJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func getEscrow(initializer solana.PublicKey) []byte {
	rpcClient := rpc.New(utils.NETWORK)
	escrow, _ := wallet.GetEscrow(initializer)

	empty := []byte("{}")

	escrowData := wallet.GetEscrowData(rpcClient, escrow)
	if escrowData == nil {
		return empty
	}

	escrowDataJson, _ := json.MarshalIndent(escrowData, "", "  ")

	return escrowDataJson
}
