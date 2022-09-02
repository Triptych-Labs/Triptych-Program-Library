package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go/programs/token"

	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	quest_ops "triptych.labs/questing/quests/ops"
	"triptych.labs/questing_tests/v2/integrations"
	"triptych.labs/utils"

	"github.com/gagliardetto/solana-go"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
)

const DEVNET = "https://sparkling-dark-shadow.solana-devnet.quiknode.pro/0e9964e4d70fe7f856e7d03bc7e41dc6a2b84452/"
const TESTNET = "https://api.testnet.solana.com"
const NETWORK = DEVNET

const LEFT = 1
const RIGHT = 2

var STAKING_MINT solana.PublicKey

func init() {
	questing.SetProgramID(solana.MustPublicKeyFromBase58("9iMuz8Lf27R9Y2jQhWM1wrSVtPB4Tt5wqkh1opjMTK11"))
}

func main() {
	// CreateNTokenAccountsOfMint(solana.MustPublicKeyFromBase58("6c5EBgbPnpdZgKhXW4uTtcYojXqVNnVQbS2cdCHo8Zmu"), 2)
	enableQuests()
	// createQuest()

	// createStakingQuest()
	// createQuestReward(1)
	integrations.CreateNStakingQuests()
	// integrations.CreateNRewardQuests()

	// startAndEndQuest()

	hash := sha256.Sum256([]byte("account:QuestAccount"))
	encoded := base58.Encode(hash[:8])
	fmt.Println(string(encoded))

	// GetQuestPdas()
	// GetQuestsKPIs()

}
func GetQuestPdas() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	for i := range make([]int, 7) {
		questPda, _ := quests.GetQuest(oracle.PublicKey(), uint64(i))
		questData := quests.GetQuestData(rpcClient, questPda)
		log.Println(i, questPda, questData)
	}

}

func GetQuestsKPIs() {
	rpcClient := rpc.New("https://devnet.genesysgo.net")
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	kpis := quest_ops.GetQuested(rpcClient, oracle.PublicKey(), solana.MustPublicKeyFromBase58("J2Y8TpYwpNshLcmnbjp9frPW9wHKRfWf3Yc26SjD1qmv"))
	j, _ := json.MarshalIndent(kpis, "", "  ")
	log.Println(string(j))

}

func CreateNTokenAccountsOfMint(mint solana.PublicKey, amount int) {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	tokenAccounts := make([]string, amount)
	var instructions []solana.Instruction
	for i := range tokenAccounts {
		wallet := solana.NewWallet()
		ata, _ := utils.GetTokenWallet(wallet.PublicKey(), mint)
		tokenAccounts[i] = ata.String()

		instructions = append(instructions,
			atok.NewCreateInstructionBuilder().
				SetPayer(oracle.PublicKey()).
				SetWallet(wallet.PublicKey()).
				SetMint(mint).
				Build(),
		)

	}
	utils.SendTx(
		"list",
		instructions,
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
	fmt.Println(tokenAccounts)
}

func formatAsJson(data interface{}) {
	dataJson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(dataJson))
}

func enableQuests() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	quest_ops.EnableQuests(oracle)
	rewardsMints := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
		solana.NewWallet().PrivateKey,
		solana.NewWallet().PrivateKey,
	}

	rewards := []questing.Reward{
		{
			MintAddress: rewardsMints[0].PublicKey(),
			Threshold:   40,
			Amount:      2,
		},
		{
			MintAddress: rewardsMints[1].PublicKey(),
			Threshold:   40,
			Amount:      4,
		},
		{
			MintAddress: rewardsMints[2].PublicKey(),
			Threshold:   20,
			Amount:      1,
		},
	}

	for i, reward := range rewards {
		_, _ = i, reward
		// quest_ops.RegisterQuestReward(oracle, reward, rewardsMints[i])
	}
}

func createStakingQuest() {
	rpcClient := rpc.New("https://devnet.genesysgo.net")
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	stakingRewardIx, stakingMint := quest_ops.RegisterQuestsStakingReward(oracle.PublicKey(), "Not qstApeCoin", "qstNBA")
	// stakingMint := solana.MustPublicKeyFromBase58("FsJTaKL31xeEPgPf8yFysCTZYj8v4B8Vvc8FSNKw3uLX")

	ixs := make([]solana.Instruction, 0)
	questData := questing.Quest{
		Index:           0,
		Name:            "Gen1 Quest",
		Duration:        60 * 5,
		Oracle:          oracle.PublicKey(),
		WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
		Tender:          nil,
		TenderSplits:    nil,
		StakingConfig: &questing.StakingConfig{
			MintAddress:  stakingMint.PublicKey(),
			YieldPer:     10, // 10 secounds
			YieldPerTime: 5,  // 5 tokens
		},
		PairsConfig: &questing.PairsConfig{
			Left: LEFT,
			LeftCreators: [5]solana.PublicKey{
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
			},
			Right: RIGHT,
			RightCreators: [5]solana.PublicKey{
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
			},
		},
	}

	creationIx, questIndex := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData)
	ixs = append(ixs, creationIx)
	ixs = append(ixs, stakingRewardIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle, stakingMint),
		// append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
	rewardsMints := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
	}

	rewards := []questing.Reward{
		{
			MintAddress: rewardsMints[0].PublicKey(),
			Threshold:   50,
			Amount:      10,
		},
	}

	for i, reward := range rewards {
		_, _ = i, reward
		_ = questIndex
		// quest_ops.RegisterQuestReward(oracle, questIndex, reward, rewardsMints[i])
	}

}

func createQuest() {
	rpcClient := rpc.New("https://devnet.genesysgo.net")
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	tenderMint := solana.MustPublicKeyFromBase58("62crD8aYknDsrsEyLrqM7YWwPZtaPbW9SzoDtifpd1dY")
	tenderMintMeta, _ := utils.GetTokensMetadataData(rpcClient, []solana.PublicKey{tenderMint})

	ixs := make([]solana.Instruction, 0)
	questData := questing.Quest{
		Index:           0,
		Name:            "Gen2 WL Quest",
		Duration:        60 * 15,
		Oracle:          oracle.PublicKey(),
		WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
		Tender: &questing.Tender{
			MintAddress: tenderMint,
			Amount:      utils.ConvertUiAmountToAmount(float64(5), tenderMintMeta[tenderMint].Decimals),
		},
		TenderSplits: &[]questing.Split{
			{
				TokenAddress: solana.PublicKey{},
				OpCode:       0,
				Share:        100,
			},
		},
		StakingConfig: nil,
		PairsConfig: &questing.PairsConfig{
			Left: LEFT,
			LeftCreators: [5]solana.PublicKey{
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
				solana.MustPublicKeyFromBase58("HRRfCf1Uvak1cqvbXjhH7FcenK3LJxJL3cc2MF8yajBq"),
			},
			Right: RIGHT,
			RightCreators: [5]solana.PublicKey{
				solana.MustPublicKeyFromBase58("CM3ZXwWgJp6ec7BTEGreQZAhtNmaUfhZRMuR3QVidGyx"),
				solana.MustPublicKeyFromBase58("CM3ZXwWgJp6ec7BTEGreQZAhtNmaUfhZRMuR3QVidGyx"),
				solana.MustPublicKeyFromBase58("CM3ZXwWgJp6ec7BTEGreQZAhtNmaUfhZRMuR3QVidGyx"),
				solana.MustPublicKeyFromBase58("CM3ZXwWgJp6ec7BTEGreQZAhtNmaUfhZRMuR3QVidGyx"),
				solana.MustPublicKeyFromBase58("CM3ZXwWgJp6ec7BTEGreQZAhtNmaUfhZRMuR3QVidGyx"),
			},
		},
	}
	creationIx, questIndex := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData)
	ixs = append(ixs, creationIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle),
		// append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)
	rewardsMints := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
	}

	rewards := []questing.Reward{
		{
			MintAddress: rewardsMints[0].PublicKey(),
			Threshold:   50,
			Amount:      10,
		},
	}

	for i, reward := range rewards {
		quest_ops.RegisterQuestReward(oracle, questIndex, reward, rewardsMints[i], "Not qstNBAGen2 Whitelist", "qstNBAG2WL")
	}

}

func createQuestReward(questIndex uint64) {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	rewardsMints := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
	}

	rewards := []questing.Reward{
		{
			MintAddress: rewardsMints[0].PublicKey(),
			Threshold:   50,
			Amount:      10,
		},
	}

	for i, reward := range rewards {
		quest_ops.RegisterQuestReward(oracle, questIndex, reward, rewardsMints[i], "Not qstNBAGen2 Whitelist", "qstNBAG2WL")
	}

}

func startAndEndQuest() {
	rpcClient := rpc.New("https://devnet.genesysgo.net")
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	questPda, _ := quests.GetQuest(oracle.PublicKey(), questsData.Quests-1)

	nfts := make([]solana.PrivateKey, 0)
	for range make([]int, LEFT+RIGHT) {
		nfts = append(nfts, solana.NewWallet().PrivateKey)
	}

	questRecorder, _ := quests.GetQuestRecorder(questPda, oracle.PublicKey())

	{
		{ //@Mint-NFTs
			for i := range nfts {
				pixelBallzMint := nfts[i]
				pixelBallzTokenAddress, _ := utils.GetTokenWallet(oracle.PublicKey(), pixelBallzMint.PublicKey())
				var instructions []solana.Instruction

				client := rpc.New(NETWORK)
				min, err := client.GetMinimumBalanceForRentExemption(context.TODO(), token.MINT_SIZE, rpc.CommitmentFinalized)
				if err != nil {
					panic(err)
				}

				instructions = append(instructions,
					system.NewCreateAccountInstructionBuilder().
						SetOwner(token.ProgramID).
						SetNewAccount(pixelBallzMint.PublicKey()).
						SetSpace(token.MINT_SIZE).
						SetFundingAccount(oracle.PublicKey()).
						SetLamports(min).
						Build(),

					token.NewInitializeMint2InstructionBuilder().
						SetMintAccount(pixelBallzMint.PublicKey()).
						SetDecimals(0).
						SetMintAuthority(oracle.PublicKey()).
						SetFreezeAuthority(oracle.PublicKey()).
						Build(),

					atok.NewCreateInstructionBuilder().
						SetPayer(oracle.PublicKey()).
						SetWallet(oracle.PublicKey()).
						SetMint(pixelBallzMint.PublicKey()).
						Build(),

					token.NewMintToInstructionBuilder().
						SetMintAccount(pixelBallzMint.PublicKey()).
						SetDestinationAccount(pixelBallzTokenAddress).
						SetAuthorityAccount(oracle.PublicKey()).
						SetAmount(1).
						Build(),
				)
				utils.SendTx(
					"list",
					instructions,
					append(make([]solana.PrivateKey, 0), oracle, pixelBallzMint),
					oracle.PublicKey(),
				)
			}
		}

		/*
		   fmt.Println("sleeping")
		   time.Sleep(15 * time.Second)
		*/

		{ //@Land-Quest
			// create quest recorder
			var instructions = make([]solana.Instruction, 0)

			createQuestRecorderIx := quest_ops.RegisterQuestRecorder(rpcClient, oracle.PublicKey(), questPda)
			if createQuestRecorderIx != nil {

				instructions = append(instructions, createQuestRecorderIx)

				utils.SendTx(
					"list",
					instructions,
					append(make([]solana.PrivateKey, 0), oracle),
					oracle.PublicKey(),
				)

			}
		}

		{ //@Affirm-Quest
			// propose quest
			var instructions = make([]solana.Instruction, 0)

			depositingLeft := func() []solana.PublicKey {
				left := make([]solana.PublicKey, 0)
				for _, mint := range nfts[0:LEFT] {
					left = append(left, mint.PublicKey())
				}
				return left
			}()
			depositingRight := func() []solana.PublicKey {
				right := make([]solana.PublicKey, 0)
				for _, mint := range nfts[LEFT:] {
					right = append(right, mint.PublicKey())
				}
				return right
			}()

			questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
			questProposalIndex := questRecorderData.Proposals
			questProposal, questProposalBump := quests.GetQuestProposal(questPda, oracle.PublicKey(), questProposalIndex)

			zero := uint64(0)
			questProposalIx, _ := quest_ops.NewQuestProposal(rpcClient, oracle.PublicKey(), questPda, depositingLeft, depositingRight, &zero)

			instructions = append(instructions, questProposalIx)

			for i, nft := range nfts {
				nftTokenAccount, _ := utils.GetTokenWallet(oracle.PublicKey(), nft.PublicKey())

				enterIx := questing.NewEnterQuestInstructionBuilder().
					SetInitializerAccount(oracle.PublicKey()).
					SetPixelballzTokenAccountAccount(nftTokenAccount).
					SetQuestAccount(questPda).
					SetQuestProposalAccount(questProposal).
					SetQuestProposalBump(questProposalBump).
					SetQuestProposalIndex(questProposalIndex).
					SetRentAccount(solana.SysVarRentPubkey).
					SetSystemProgramAccount(solana.SystemProgramID).
					SetTokenProgramAccount(solana.TokenProgramID)

				if LEFT > i {
					enterIx.SetSideEnum("left")
				}
				if LEFT <= i {
					enterIx.SetSideEnum("right")
				}

				instructions = append(instructions, enterIx.Build())
			}

			utils.SendTx(
				"list",
				instructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)
		}

		{ //@Start-Quest
			questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)

			questInstructions := make([]solana.Instruction, 0)

			startQuestIx := quest_ops.StartQuest(rpcClient, oracle.PublicKey(), questPda, questRecorderData.Proposals-1)

			questInstructions = append(
				questInstructions,
				startQuestIx,
			)

			utils.SendTx(
				"init cm",
				questInstructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)

		}

		{ //@Claim-Quest-Reward
			fmt.Println("sleeping...")
			time.Sleep(10 * time.Second)
			questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)

			questInstructions := make([]solana.Instruction, 0)

			claimIx := quest_ops.ClaimQuestStakingReward(rpcClient, oracle.PublicKey(), questPda, questRecorderData.Proposals-1)

			questInstructions = append(
				questInstructions,
				claimIx,
			)

			utils.SendTx(
				"init cm",
				questInstructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)

		}

		{ //@End-Quest
			questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)

			questInstructions := make([]solana.Instruction, 0)

			endQuestIx := quest_ops.EndQuest(rpcClient, oracle.PublicKey(), questPda, questRecorderData.Proposals-1)

			questInstructions = append(
				questInstructions,
				endQuestIx,
			)

			utils.SendTx(
				"init cm",
				questInstructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)

		}

		{ //@Flush-Quest
			var instructions = make([]solana.Instruction, 0)

			questRecorderData := quests.GetQuestRecorderData(rpcClient, questRecorder)
			questProposalIndex := questRecorderData.Proposals - 1

			instructions = append(instructions, quest_ops.FlushQuestRecord(rpcClient, oracle.PublicKey(), questPda, questProposalIndex)...)

			utils.SendTx(
				"list",
				instructions,
				append(make([]solana.PrivateKey, 0), oracle),
				oracle.PublicKey(),
			)
		}

	}
}
