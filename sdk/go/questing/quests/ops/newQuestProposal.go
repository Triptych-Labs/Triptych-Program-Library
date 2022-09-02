package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

func NewQuestProposal(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, depositingLeft, depositingRight []solana.PublicKey, proposalIndexOffset *uint64) (*questing.Instruction, *uint64) {
	questRecorder, _ := quests.GetQuestRecorder(questPda, initializer)

	var index uint64 = 0
	questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
	index = questRecorderData.Proposals
	if proposalIndexOffset != nil {
		index += *proposalIndexOffset
	}
	questProposal, _ := quests.GetQuestProposal(questPda, initializer, index)

	questProposalIx := questing.NewProposeQuestRecordInstructionBuilder().
		SetDepositingLeft(depositingLeft).
		SetDepositingRight(depositingRight).
		SetInitializerAccount(initializer).
		SetQuestAccount(questPda).
		SetQuestProposalAccount(questProposal).
		SetQuestRecorderAccount(questRecorder).
		SetSystemProgramAccount(solana.SystemProgramID)

	if e := questProposalIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return questProposalIx.Build(), &index
}

