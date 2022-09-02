package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func StartQuest(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, questProposalIndex uint64) *questing.Instruction {
	questData := quests.GetQuestData(rpcClient, questPda)
	questRecorder, questRecorderBump := quests.GetQuestRecorder(questPda, initializer)
	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, questProposalIndex)
	questAccount, _ := quests.GetQuestAccount(initializer, questProposal, questPda)

	fmt.Println(initializer, questPda, questAccount, questProposal, questRecorder)
	startQuestIx := questing.NewStartQuestInstructionBuilder().
		SetInitializerAccount(initializer).
		SetQuestAccount(questPda).
		SetQuestAccAccount(questAccount).
		SetQuestProposalAccount(questProposal).
		SetQuestProposalBump(questProposalBump).
		SetQuestProposalIndex(questProposalIndex).
		SetQuestRecorderAccount(questRecorder).
		SetQuestRecorderBump(questRecorderBump).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	if questData.Tender != nil && questData.TenderSplits != nil {
		fmt.Println(*questData.Tender, questData.TenderSplits)
		tenderTokenAccount, _ := utils.GetTokenWallet(initializer, questData.Tender.MintAddress)
		startQuestIx.Append(&solana.AccountMeta{PublicKey: tenderTokenAccount, IsWritable: true, IsSigner: false})
		for _, tenderSplit := range *questData.TenderSplits {
			if tenderSplit.TokenAddress.Equals(solana.SystemProgramID) {
				startQuestIx.Append(&solana.AccountMeta{PublicKey: questData.Tender.MintAddress, IsWritable: true, IsSigner: false})
			} else {
				startQuestIx.Append(&solana.AccountMeta{PublicKey: tenderSplit.TokenAddress, IsWritable: true, IsSigner: false})
			}
		}
	}

	for _, account := range startQuestIx.AccountMetaSlice {
		fmt.Println(account)
	}

	if e := startQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		return nil
	}

	return startQuestIx.Build()
}
