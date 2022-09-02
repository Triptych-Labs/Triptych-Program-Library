package questing

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing/quests/ops"
	"triptych.labs/utils"
)

func GetQuestsKPIs(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	holder := solana.MustPublicKeyFromBase58(args[1].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					fmt.Println("kpis", r)
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Please retry")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)
			rpcClient := rpc.New(utils.NETWORK)

			questsJSON, err := getQuestsKPIs(rpcClient, oracle, holder)
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}
			dst := js.Global().Get("Uint8Array").New(len(questsJSON))
			js.CopyBytesToJS(dst, questsJSON)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func getQuestsKPIs(rpcClient *rpc.Client, oracle, holder solana.PublicKey) ([]byte, error) {
	questsKPIs := ops.GetQuested(rpcClient, oracle, holder)

	return json.Marshal(questsKPIs)
}
