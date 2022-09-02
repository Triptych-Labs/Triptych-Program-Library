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

func FlushQuestRecords(this js.Value, args []js.Value) interface{} {
	holder := solana.MustPublicKeyFromBase58(args[0].String())
	quest := solana.MustPublicKeyFromBase58(args[1].String())
	proposalIndexesInp := args[2].String()

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Please retry")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)
			var proposalIndexes []uint64
			json.Unmarshal([]byte(proposalIndexesInp), &proposalIndexes)

			if len(proposalIndexes) == 0 {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("invalid proposals length")
				reject.Invoke(errorObject)
				return
			}

			enrollmentJson, err := flushQuestRecords(holder, quest, proposalIndexes)
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

func flushQuestRecords(holder, quest solana.PublicKey, proposalIndexes []uint64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	transactions := make([]solana.Transaction, 0)
	responseJson := []byte("[]")

	for _, proposalIndex := range proposalIndexes {
		txBuilder := solana.NewTransactionBuilder()
		for _, ix := range ops.FlushQuestRecord(rpcClient, holder, quest, proposalIndex) {
			txBuilder = txBuilder.AddInstruction(ix)
		}

		txB, _ := txBuilder.Build()
		transactions = append(transactions, *txB)

	}
	if len(transactions) > 0 {
		responseJson, _ = json.MarshalIndent(transactions, "", "  ")
	}
	fmt.Println("....", string(responseJson))
	return responseJson, nil

}
