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

func OnboardFromSingletons(this js.Value, args []js.Value) interface{} {
	holder := solana.MustPublicKeyFromBase58(args[0].String())
	quest := solana.MustPublicKeyFromBase58(args[1].String())
	depositingInp := args[2].String()

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
			var depositing []solana.PublicKey
			json.Unmarshal([]byte(depositingInp), &depositing)

			if len(depositing) == 0 {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("invalid deposits")
				reject.Invoke(errorObject)
				return
			}

			enrollmentJson, err := onboardFromSingletons(holder, quest, depositing)
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

func onboardFromSingletons(holder, quest solana.PublicKey, depositing []solana.PublicKey) ([]byte, error) {
	// TODO batch 3 onboards at a time
	rpcClient := rpc.New(utils.NETWORK)

	transactions := make([]*solana.Transaction, 0)
	responseJson := []byte("[]")

	instructions := make([]solana.Instruction, 0)
	for i, deposit := range depositing {
		proposalIndexOffset := uint64(i)
		questProposalIx, questProposalIndex := ops.NewQuestProposal(rpcClient, holder, quest, []solana.PublicKey{deposit}, []solana.PublicKey{}, &proposalIndexOffset)

		instructions = append(
			instructions,
			questProposalIx,
			ops.EnterQuest(rpcClient, holder, quest, deposit, "left", questProposalIndex),
			ops.StartQuest(rpcClient, holder, quest, *questProposalIndex),
		)

	}
	txBuilder := solana.NewTransactionBuilder()
	for _, ix := range instructions {
		txBuilder = txBuilder.AddInstruction(ix)
	}
	txB, _ := txBuilder.Build()
	transactions = append(transactions, txB)

	responseJson, _ = json.MarshalIndent(transactions, "", "  ")

	fmt.Println("....", string(responseJson))
	return responseJson, nil

}
