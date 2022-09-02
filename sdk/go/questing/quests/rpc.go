package quests

import (
	"context"
	"crypto/sha256"
	"log"

	ag_binary "github.com/gagliardetto/binary"

	"github.com/gagliardetto/solana-go/rpc"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
)

func GetQuestsData(rpcClient *rpc.Client, quests solana.PublicKey) *questing.Quests {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), quests, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.Quests
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestData(rpcClient *rpc.Client, quest solana.PublicKey) *questing.Quest {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), quest, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.Quest
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestorData(rpcClient *rpc.Client, questor solana.PublicKey) *questing.Questor {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questor, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.Questor
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}
func GetQuesteeData(rpcClient *rpc.Client, questee solana.PublicKey) *questing.Questee {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questee, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.Questee
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestQuesteeReceiptAccountData(rpcClient *rpc.Client, questQuesteeReceipt solana.PublicKey) *questing.QuestQuesteeEndReceipt {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questQuesteeReceipt, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.QuestQuesteeEndReceipt
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestAccountData(rpcClient *rpc.Client, questAccount solana.PublicKey) *questing.QuestAccount {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questAccount, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.QuestAccount
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestRecorderData(rpcClient *rpc.Client, questRecorder solana.PublicKey) *questing.QuestRecorder {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questRecorder, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.QuestRecorder
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestProposalData(rpcClient *rpc.Client, questProposal solana.PublicKey) *questing.QuestProposal {
	bin, _ := rpcClient.GetAccountInfoWithOpts(context.TODO(), questProposal, &rpc.GetAccountInfoOpts{Commitment: "confirmed"})
	if bin == nil {
		return nil
	}
	var data questing.QuestProposal
	decoder := ag_binary.NewBorshDecoder(bin.Value.Data.GetBinary())
	err := data.UnmarshalWithDecoder(decoder)
	if err != nil {
		panic(err)
	}

	return &data

}

func GetQuestsProposals(rpcClient *rpc.Client, initializer, questPda solana.PublicKey) ([]questing.QuestProposal, []*questing.QuestAccount) {
	batch := 100

	questRecorder, _ := GetQuestRecorder(questPda, initializer)
	questRecorderData := GetQuestRecorderData(rpcClient, questRecorder)
	if questRecorderData == nil {
		return []questing.QuestProposal{}, []*questing.QuestAccount{}
	}
	if questRecorderData.Proposals == 0 {
		return []questing.QuestProposal{}, []*questing.QuestAccount{}
	}

	questProposalsData := make([]questing.QuestProposal, questRecorderData.Proposals)
	questProposalsPdas := make([]solana.PublicKey, questRecorderData.Proposals)
	questAccData := make([]*questing.QuestAccount, questRecorderData.Proposals)
	questAccPdas := make([]solana.PublicKey, questRecorderData.Proposals)

	for i := range questProposalsPdas {
		questProposalsPdas[i], _ = GetQuestProposal(questPda, initializer, uint64(i))
	}

	for i := range questProposalsPdas {
		questAccPdas[i], _ = GetQuestAccount(initializer, questProposalsPdas[i], questPda)
	}

	for i := 0; i < len(questProposalsPdas); i += batch {
		j := i + batch
		if j > len(questProposalsPdas) {
			j = len(questProposalsPdas)
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), questProposalsPdas[i:j]...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			var data questing.QuestProposal
			if accounts.Value[ii%batch] == nil {
				questProposalsData[ii] = data
				continue
			}
			if accounts.Value[ii%batch].Data == nil {
				questProposalsData[ii] = data
				continue
			}
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			questProposalsData[ii] = data
		}

	}

	for i := 0; i < len(questAccPdas); i += batch {
		j := i + batch
		if j > len(questAccPdas) {
			j = len(questAccPdas)
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), questAccPdas[i:j]...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			if accounts.Value[ii%batch] == nil {
				questAccData[ii] = nil
				continue
			}
			if len(accounts.Value[ii%batch].Data.GetBinary()) == 0 {
				questAccData[ii] = nil
				continue
			}
			var data questing.QuestAccount
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			questAccData[ii] = &data
		}

	}

	return questProposalsData, questAccData

}

func GetQuestKPIs(rpcClient *rpc.Client, oracle, initializer solana.PublicKey) ([]questing.QuestRecorder, []questing.QuestAccount) {
	batch := 100

	questAccHash := sha256.Sum256([]byte("account:QuestAccount"))
	recorderHash := sha256.Sum256([]byte("account:QuestRecorder"))
	zero := uint64(0)
	recorderAccounts, _ := rpcClient.GetProgramAccountsWithOpts(context.TODO(), questing.ProgramID, &rpc.GetProgramAccountsOpts{
		Encoding: "base64",
		DataSlice: &rpc.DataSlice{
			Offset: &zero,
			Length: &zero,
		},
		Filters: append(
			make([]rpc.RPCFilter, 0),
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 0,
					Bytes:  recorderHash[:8],
				},
			},
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 80,
					Bytes:  oracle.Bytes(),
				},
			},
		),
	})
	questAccAccounts, _ := rpcClient.GetProgramAccountsWithOpts(context.TODO(), questing.ProgramID, &rpc.GetProgramAccountsOpts{
		Encoding: "base64",
		DataSlice: &rpc.DataSlice{
			Offset: &zero,
			Length: &zero,
		},
		Filters: append(
			make([]rpc.RPCFilter, 0),
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 0,
					Bytes:  questAccHash[:8],
				},
			},
			rpc.RPCFilter{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 64,
					Bytes:  initializer.Bytes(),
				},
			},
		),
	})
	if recorderAccounts == nil || questAccAccounts == nil {
		return []questing.QuestRecorder{}, []questing.QuestAccount{}
	}

	recordersData := make([]questing.QuestRecorder, len(recorderAccounts))
	questAccData := make([]questing.QuestAccount, len(questAccAccounts))

	for i := 0; i < len(recorderAccounts); i += batch {
		j := i + batch
		if j > len(recorderAccounts) {
			j = len(recorderAccounts)
		}
		recorders := make([]solana.PublicKey, batch)
		for ii, account := range recorderAccounts[i:j] {
			recorders[ii] = account.Pubkey
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), recorders...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			var data questing.QuestRecorder
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			recordersData[ii] = data
		}

	}

	for i := 0; i < len(questAccAccounts); i += batch {
		j := i + batch
		if j > len(questAccAccounts) {
			j = len(questAccAccounts)
		}
		recorders := make([]solana.PublicKey, batch)
		for ii, account := range questAccAccounts[i:j] {
			recorders[ii] = account.Pubkey
		}

		accounts, err := rpcClient.GetMultipleAccounts(context.TODO(), recorders...)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(accounts.Value) == 0 {
			log.Println("empty accounts")
			continue
		}

		for ii := i; ii < j; ii++ {
			var data questing.QuestAccount
			decoder := ag_binary.NewBorshDecoder(accounts.Value[ii%batch].Data.GetBinary())
			err := data.UnmarshalWithDecoder(decoder)
			if err != nil {
				panic(err)
			}
			questAccData[ii] = data
		}

	}

	return recordersData, questAccData

}

