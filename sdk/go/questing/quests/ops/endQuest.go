package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func EndQuest(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, questProposalIndex uint64) *questing.Instruction {
	questData := quests.GetQuestData(rpcClient, questPda)
	_, questPdaBump := quests.GetQuest(questData.Oracle, questData.Index)
	questsPda, questsPdaBump := quests.GetQuests(questData.Oracle)
	questsData := quests.GetQuestsData(rpcClient, questsPda)

	questRecorder, questRecorderBump := quests.GetQuestRecorder(questPda, initializer)
	questProposal, questProposalBump := quests.GetQuestProposal(questPda, initializer, questProposalIndex)
	questAccount, _ := quests.GetQuestAccount(initializer, questProposal, questPda)
	endQuestIx := questing.NewEndQuestInstructionBuilder().
		SetInitializerAccount(initializer).
		SetOracleAccount(questData.Oracle).
		SetQuestAccAccount(questAccount).
		SetQuestAccount(questPda).
		SetQuestBump(questPdaBump).
		SetQuestProposalAccount(questProposal).
		SetQuestProposalBump(questProposalBump).
		SetQuestProposalIndex(questProposalIndex).
		SetQuestRecorderAccount(questRecorder).
		SetQuestRecorderBump(questRecorderBump).
		SetQuestsAccount(questsPda).
		SetQuestsBump(questsPdaBump).
		SetRentAccount(solana.SysVarRentPubkey).
		SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID).
		SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	for _, reward := range questData.Rewards {
		endQuestIx.Append(&solana.AccountMeta{PublicKey: reward.MintAddress, IsWritable: true, IsSigner: false})
	}
	for _, reward := range questData.Rewards {
		rewardAta, _ := utils.GetTokenWallet(initializer, reward.MintAddress)
		endQuestIx.Append(&solana.AccountMeta{PublicKey: rewardAta, IsWritable: true, IsSigner: false})
	}

	for _, reward := range questsData.Rewards {
		endQuestIx.Append(&solana.AccountMeta{PublicKey: reward.MintAddress, IsWritable: true, IsSigner: false})
	}
	for _, reward := range questsData.Rewards {
		rewardAta, _ := utils.GetTokenWallet(initializer, reward.MintAddress)
		endQuestIx.Append(&solana.AccountMeta{PublicKey: rewardAta, IsWritable: true, IsSigner: false})
	}

	if e := endQuestIx.Validate(); e != nil {
		fmt.Println(e.Error())
		return nil
	}

	return endQuestIx.Build()
}
