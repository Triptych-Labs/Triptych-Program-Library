package ops

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/utils"
)

type QuestedMeta struct {
	TotalStaked    uint64  `json:"totalStaked"`
	StakingRewards float64 `json:"stakingRewards"`
}

// QuestedMetaMap { [questPda]: []QuestedMeta }
type QuestedMetaMap map[solana.PublicKey]QuestedMeta

func GetQuested(rpcClient *rpc.Client, oracle, initializer solana.PublicKey) QuestedMetaMap {
	questedMetaMap := make(QuestedMetaMap)
	recorderAccounts, questAccounts := quests.GetQuestKPIs(rpcClient, oracle, initializer)

	questsData := make(map[solana.PublicKey]questing.Quest)

	for _, account := range recorderAccounts {
		if _, ok := questedMetaMap[account.Quest]; !ok {
			questsData[account.Quest] = *quests.GetQuestData(rpcClient, account.Quest)

			questedMetaMap[account.Quest] = QuestedMeta{
				TotalStaked:    0,
				StakingRewards: 0,
			}
		}

		record := questedMetaMap[account.Quest]

		record.TotalStaked += uint64(len(account.Staked))

		questedMetaMap[account.Quest] = record
	}

	for _, account := range questAccounts {
		questPda, _ := quests.GetQuest(oracle, account.Index)
		if !account.Quest.Equals(questPda) {
			continue
		}

		record := questedMetaMap[questPda]

		accountJson, _ := json.MarshalIndent(account, "", "  ")
		fmt.Println(string(accountJson))
		if *account.Completed != true && account.StartTime != 0 {
			if questsData[questPda].StakingConfig != nil {
				duration := int64(0)
				now := time.Now().UTC().Unix()
				// TODO REIMPLEMENT
				/*
					if now >= account.EndTime {
						duration = account.EndTime - account.LastClaim
					} else {
					}
				*/
				duration = now - account.LastClaim
				beta := float64(duration) / float64(questsData[questPda].StakingConfig.YieldPerTime)
				alpha := uint64(float64(questsData[questPda].StakingConfig.YieldPer) * beta)
				// questsData[questPda].StakingConfig.YieldPer
				record.StakingRewards += utils.ConvertAmountToUiAmount(alpha, 1)
			}
		}

		questedMetaMap[questPda] = record
	}

	return questedMetaMap
}
