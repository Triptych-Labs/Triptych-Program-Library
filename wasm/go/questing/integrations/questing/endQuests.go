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

func EndQuests(this js.Value, args []js.Value) interface{} {
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

			enrollmentJson, err := endQuests(holder, quest, proposalIndexes)
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

func endQuests(holder, quest solana.PublicKey, questProposalsIndexes []uint64) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([][]solana.Instruction, len(questProposalsIndexes))
	transactions := make([]solana.Transaction, 0)
	txJson := []byte("[]")

	for i, questProposalIndex := range questProposalsIndexes {
		innerInstructions := make([]solana.Instruction, 0)
		if claimIx := ops.ClaimQuestStakingReward(rpcClient, holder, quest, questProposalIndex); claimIx != nil {
			innerInstructions = append(
				innerInstructions,
				claimIx,
			)
		}
		if endQuestIx := ops.EndQuest(rpcClient, holder, quest, questProposalIndex); endQuestIx != nil {
			innerInstructions = append(
				innerInstructions,
				endQuestIx,
			)
			for _, ix := range ops.FlushQuestRecord(rpcClient, holder, quest, questProposalIndex) {
				innerInstructions = append(
					innerInstructions,
					ix,
				)
			}
		}

		instructions[i] = innerInstructions
	}

	if len(instructions) > 0 {
		for _, ixs := range instructions {
			txBuilder := solana.NewTransactionBuilder()
			for _, ix := range ixs {
				txBuilder = txBuilder.AddInstruction(ix)
			}
			txB, _ := txBuilder.Build()
			transactions = append(transactions, *txB)
		}
	}

	txJson, _ = json.MarshalIndent(transactions, "", "  ")
	fmt.Println(string(txJson))
	return txJson, nil

}
