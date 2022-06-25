package ops

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
	"triptych.labs/questing/quests"
)

type QuestedMeta struct {
	Questee      questing.Questee
	QuestAccount questing.QuestAccount
}

// QuestedMetaMap { [questPda]: []QuestedMeta }
type QuestedMetaMap map[solana.PublicKey][]QuestedMeta

func GetQuested(oracle, holder solana.PublicKey) QuestedMetaMap {
	questedMetaMap := make(QuestedMetaMap)
	questAccounts := quests.GetQuestAccountsDataForInitializer(holder)
	for _, questAccount := range questAccounts {
		questPda, _ := quests.GetQuest(oracle, questAccount.Index)
		questedMetaMap[questPda] = make([]QuestedMeta, 0)
	}
	for _, questAccount := range questAccounts {
		questPda, _ := quests.GetQuest(oracle, questAccount.Index)
		questee, _ := quests.GetQuesteeAccount(questAccount.DepositTokenMint)
		questeeData := quests.GetQuesteeData(questee)
		questedMetaMap[questPda] = append(
			questedMetaMap[questPda],
			QuestedMeta{
				Questee:      *questeeData,
				QuestAccount: questAccount,
			},
		)
	}

	return questedMetaMap
}
