package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func EnterQuest(rpcClient *rpc.Client, initializer, questPda, deposit solana.PublicKey, side string, proposalIndex *uint64) *questing.Instruction {
	var index uint64 = 0
	if proposalIndex != nil {
		index = *proposalIndex
	} else {
		questRecorder, _ := quests.GetQuestRecorder(questPda, initializer)
		questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
		if questRecorderData == nil {
			return nil
		}
		index = questRecorderData.Proposals
	}
	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, index)

	depositMetadata, _ := utils.GetMetadata(deposit)
	nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)

	enterIx := questing.NewEnterQuestInstructionBuilder().
		SetInitializerAccount(initializer).
		SetPixelballzMetadataAccount(depositMetadata).
		SetPixelballzTokenAccountAccount(nftTokenAccount).
		SetQuestAccount(questPda).
		SetQuestProposalAccount(questProposal).
		SetQuestProposalBump(questProposalBump).
		SetQuestProposalIndex(index).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID).
		SetSideEnum(side)

	if e := enterIx.Validate(); e != nil {
		fmt.Println(e.Error())
		panic("...")
	}

	return enterIx.Build()
}
