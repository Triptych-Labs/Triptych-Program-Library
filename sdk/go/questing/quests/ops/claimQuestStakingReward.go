package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

func ClaimQuestStakingReward(rpcClient *rpc.Client, initializer, questPda solana.PublicKey, questProposalIndex uint64) *questing.Instruction {
	questData := quests.GetQuestData(rpcClient, questPda)
	questsPda, questsPdaBump := quests.GetQuests(questData.Oracle)
	questProposal, _ := quests.GetQuestProposal(questPda, initializer, questProposalIndex)
	questAccount, _ := quests.GetQuestAccount(initializer, questProposal, questPda)
	questAccountData := quests.GetQuestAccountData(rpcClient, questAccount)
	if questAccountData == nil {
		return nil
	}
	if questData.StakingConfig == nil {
		return nil
	}

	rewardTokenAccount, _ := utils.GetTokenWallet(initializer, questData.StakingConfig.MintAddress)

	claimIx := questing.NewClaimQuestStakingRewardInstructionBuilder().
		SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID).
		SetInitializerAccount(initializer).
		SetQuestAccAccount(questAccount).
		SetQuestAccount(questPda).
		SetQuestsAccount(questsPda).
		SetQuestsBump(questsPdaBump).
		SetRentAccount(solana.SysVarRentPubkey).
		SetRewardMintAccount(questData.StakingConfig.MintAddress).
		SetRewardTokenAccountAccount(rewardTokenAccount).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetTokenProgramAccount(solana.TokenProgramID)

	if e := claimIx.Validate(); e != nil {
		fmt.Println(e.Error())
		return nil
	}

	return claimIx.Build()
}
