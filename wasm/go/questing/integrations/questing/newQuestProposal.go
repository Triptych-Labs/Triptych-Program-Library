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

type NewQuestProposalResponse struct {
	Transaction   *solana.Transaction `json:"transaction"`
	ProposalIndex uint64              `json:"proposalIndex"`
}

func NewQuestProposal(this js.Value, args []js.Value) interface{} {
	holder := solana.MustPublicKeyFromBase58(args[0].String())
	quest := solana.MustPublicKeyFromBase58(args[1].String())
	depositingLeftInp, depositingRightInp := args[2].String(), args[3].String()

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
			var depositingLeft, depositingRight []solana.PublicKey
			json.Unmarshal([]byte(depositingLeftInp), &depositingLeft)
			json.Unmarshal([]byte(depositingRightInp), &depositingRight)

			if len(depositingLeft) == 0 && len(depositingRight) == 0 {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("invalid deposits")
				reject.Invoke(errorObject)
				return
			}

			enrollmentJson, err := newQuestProposal(holder, quest, depositingLeft, depositingRight)
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

func newQuestProposal(holder, quest solana.PublicKey, depositingLeft, depositingRight []solana.PublicKey) ([]byte, error) {
	rpcClient := rpc.New(utils.NETWORK)

	instructions := make([]solana.Instruction, 0)
	responseJson := []byte("{}")

	questProposalIx, questProposalIndex := ops.NewQuestProposal(rpcClient, holder, quest, depositingLeft, depositingRight, nil)
	fmt.Println("xxxx", questProposalIx)
	if questProposalIx != nil {
		instructions = append(instructions, questProposalIx)

		for _, deposit := range depositingLeft {
			instructions = append(instructions, ops.EnterQuest(rpcClient, holder, quest, deposit, "left", nil))
		}

		for _, deposit := range depositingRight {
			instructions = append(instructions, ops.EnterQuest(rpcClient, holder, quest, deposit, "right", nil))
		}

		if startQuestIx := ops.StartQuest(rpcClient, holder, quest, *questProposalIndex); startQuestIx != nil {
			fmt.Println(startQuestIx.Accounts())
			instructions = append(
				instructions,
				startQuestIx,
			)
		}

		txBuilder := solana.NewTransactionBuilder()
		for _, ix := range instructions {
			txBuilder = txBuilder.AddInstruction(ix)
		}

		txB, _ := txBuilder.Build()
		response := NewQuestProposalResponse{
			Transaction:   txB,
			ProposalIndex: *questProposalIndex,
		}

		responseJson, _ = json.MarshalIndent(response, "", "  ")
	}

	fmt.Println("....", string(responseJson))
	return responseJson, nil

}
