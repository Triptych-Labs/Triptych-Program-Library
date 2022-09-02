package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func FlushQuestRecord(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, questProposalIndex uint64) []solana.Instruction {
	var instructions = make([]solana.Instruction, 0)

	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, questProposalIndex)
	questProposalData := quests.GetQuestProposalData(rpcClient, questProposal)

	questData := quests.GetQuestData(rpcClient, questPda)
	_, questPdaBump := quests.GetQuest(questData.Oracle, questData.Index)

	for _, deposit := range questProposalData.DepositingLeft {
		nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)

		flushIx := questing.NewFlushQuestRecordInstructionBuilder().
			SetInitializerAccount(initializer).
			SetPixelballzMintAccount(deposit).
			SetPixelballzTokenAccountAccount(nftTokenAccount).
			SetQuestAccount(questPda).
			SetQuestBump(questPdaBump).
			SetQuestProposalAccount(questProposal).
			SetQuestProposalBump(questProposalBump).
			SetQuestProposalIndex(questProposalIndex).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		if e := flushIx.Validate(); e != nil {
			fmt.Println(e.Error())
			panic("...")
		}

		instructions = append(instructions, flushIx.Build())
	}
	for _, deposit := range questProposalData.DepositingRight {
		nftTokenAccount, _ := utils.GetTokenWallet(initializer, deposit)

		flushIx := questing.NewFlushQuestRecordInstructionBuilder().
			SetInitializerAccount(initializer).
			SetPixelballzMintAccount(deposit).
			SetPixelballzTokenAccountAccount(nftTokenAccount).
			SetQuestAccount(questPda).
			SetQuestBump(questPdaBump).
			SetQuestProposalAccount(questProposal).
			SetQuestProposalBump(questProposalBump).
			SetQuestProposalIndex(questProposalIndex).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		if e := flushIx.Validate(); e != nil {
			fmt.Println(e.Error())
			panic("...")
		}
		instructions = append(instructions, flushIx.Build())
	}

	return instructions
}
