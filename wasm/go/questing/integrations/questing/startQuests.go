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

func StartQuests(this js.Value, args []js.Value) interface{} {
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

			enrollmentJson, err := startQuests(holder, quest, proposalIndexes)
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

func startQuests(holder, quest solana.PublicKey, questProposalsIndexes []uint64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([]solana.Instruction, 0)
	txJson := []byte("{}")

	for _, questProposalIndex := range questProposalsIndexes {
		if startQuestIx := ops.StartQuest(rpcClient, holder, quest, questProposalIndex); startQuestIx != nil {
			fmt.Println(startQuestIx.Accounts())
			instructions = append(
				instructions,
				startQuestIx,
			)
		}
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
