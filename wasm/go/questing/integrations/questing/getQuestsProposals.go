package questing

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

type ProposalResponse struct {
	Index           uint64
	Fulfilled       bool
	Started         bool
	Finished        bool
	Withdrawn       bool
	DepositingLeft  []solana.PublicKey
	DepositingRight []solana.PublicKey
	RecordLeft      []bool
	RecordRight     []bool

	// (from) QuestAccount
	StartTime int64
	EndTime   int64
}

func GetQuestsProposals(this js.Value, args []js.Value) interface{} {
	oracle := solana.MustPublicKeyFromBase58(args[0].String())
	holder := solana.MustPublicKeyFromBase58(args[1].String())

	handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]

		go func() {
			defer func(_reject *js.Value) {
				if r := recover(); r != nil {
					fmt.Println("props", r)
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New("Please retry")
					_reject.Invoke(errorObject)
					return
				}
			}(&reject)
			questsProposalsJson, err := json.Marshal(getQuestsProposals(oracle, holder))
			if err != nil {
				errorConstructor := js.Global().Get("Error")
				errorObject := errorConstructor.New("unauthorized")
				reject.Invoke(errorObject)
				return
			}

			dst := js.Global().Get("Uint8Array").New(len(questsProposalsJson))
			js.CopyBytesToJS(dst, questsProposalsJson)

			resolve.Invoke(dst)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}

func getQuestProposals(rpcClient *rpc.Client, holder, questPda solana.PublicKey) []ProposalResponse {

	// fetch the recorder
	// fetch every proposal of the recorder
	questRecorder, _ := quests.GetQuestRecorder(questPda, holder)
	questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
	if questRecorderData == nil {
		return []ProposalResponse{}
	}
	if questRecorderData.Proposals == 0 {
		return []ProposalResponse{}
	}

	questProposalsData, questAccountsData := quests.GetQuestsProposals(rpcClient, holder, questPda)
	questProposalsResponse := make([]ProposalResponse, len(questProposalsData))
	for i := range questProposalsData {

		var startTime int64 = 0
		var endTime int64 = 0

		questProposalData := questProposalsData[i]
		questAccountData := questAccountsData[i]
		if questAccountData != nil {
			startTime = questAccountData.StartTime
			endTime = questAccountData.EndTime
		}

		questProposalsResponse[i] = ProposalResponse{
			Index:           questProposalData.Index,
			Fulfilled:       questProposalData.Fulfilled,
			Started:         questProposalData.Started,
			Finished:        questProposalData.Finished,
			Withdrawn:       questProposalData.Withdrawn,
			DepositingLeft:  questProposalData.DepositingLeft,
			DepositingRight: questProposalData.DepositingRight,
			RecordLeft:      questProposalData.RecordLeft,
			RecordRight:     questProposalData.RecordRight,
			StartTime:       startTime,
			EndTime:         endTime,
		}
	}

	return questProposalsResponse
}

func getQuestsProposals(oracle, holder solana.PublicKey) map[solana.PublicKey][]ProposalResponse {

	rpcClient := rpc.New(utils.NETWORK)
	questsProposalsData := make(map[solana.PublicKey][]ProposalResponse)

	questsData := getQuests(rpcClient, oracle)
	for quest := range questsData {
		questsProposalsData[quest] = getQuestProposals(rpcClient, holder, quest)
	}

	return questsProposalsData

}
