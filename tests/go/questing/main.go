package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go/programs/token"

	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	"triptych.labs/questing/quests/ops"
	quest_ops "triptych.labs/questing/quests/ops"
	"triptych.labs/utils"

	"github.com/gagliardetto/solana-go"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
)

const DEVNET = "https://sparkling-dark-shadow.solana-devnet.quiknode.pro/0e9964e4d70fe7f856e7d03bc7e41dc6a2b84452/"
const TESTNET = "https://api.testnet.solana.com"
const NETWORK = DEVNET

var MINT = solana.MustPublicKeyFromBase58("6c5EBgbPnpdZgKhXW4uTtcYojXqVNnVQbS2cdCHo8Zmu")

func init() {
	questing.SetProgramID(solana.MustPublicKeyFromBase58("Cr4keTx8UQiQ5F9TzTGdQ5dkcMHjxhYSAaHkHhUSABCk"))
}

func main() {
	// enable()
	// verifyBatchUpload()
	// catalogBatches()
	// treasure()
	// list()
	// verifyList()
	// treasureCMs()
	// treasureVerify()
	// treasureVerifyCM()
	// mintRare()
	// holder_nft_metadata()
	// burn()

	// marketCreate()
	// verifyMarketCreate()
	// marketList()
	// verifyMarketList()
	// marketFulfill()

	// GetMarketMintMeta()
	// GetMarketListingsData()

	// enableVias()
	// enableViaForRarityToken()

	// CreateNTokenAccountsOfMint(MINT, 2)
	// enableQuestsAndCreateQuest()
	// CreateAndAmmendEntitlementQuest()
	// startQuest()
	startAndEndQuest()
	// ETZoY7cJfD8N7EVx5tShRYS1vxgv3F4Dkavjb52kGRyj
	// treasureVerify()

	hash := sha256.Sum256([]byte("account:QuestAccount"))
	encoded := base58.Encode(hash[:8])
	fmt.Println(string(encoded))

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

func enableQuestsAndCreateQuest() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	signers := make([]solana.PrivateKey, 0)
	rewardsMints := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
		solana.NewWallet().PrivateKey,
		solana.NewWallet().PrivateKey,
	}
	signers = append(signers, rewardsMints...)

	quest_ops.EnableQuests(oracle)

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
		quest_ops.RegisterQuestReward(oracle, reward, rewardsMints[i])
	}

	ixs := make([]solana.Instruction, 0)
	questData := questing.Quest{
		Index:           0,
		Name:            "aaa",
		Duration:        100,
		Oracle:          oracle.PublicKey(),
		WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
		Entitlement:     nil,
		Tender:          nil,
		/*
			Tender: &questing.Tender{
				MintAddress: MINT,
				Amount:      5,
			},
		*/
	}
	ix, questIndex := quest_ops.CreateQuest(oracle.PublicKey(), questData)
	questData.Index = questIndex
	ixs = append(ixs, ix)

	rewardIxs := ops.AppendQuestRewards(oracle.PublicKey(), questData)
	ixs = append(ixs, rewardIxs...)

	utils.SendTx(
		"list",
		ixs,
		append(signers, oracle),
		oracle.PublicKey(),
	)

	{
		questsPda, _ := quests.GetQuests(oracle.PublicKey())
		questsData := quests.GetQuestsData(questsPda)
		quest, _ := quests.GetQuest(oracle.PublicKey(), questsData.Quests-1)
		fmt.Println(quest, questsData.Quests, questsData.Quests-1)
		questData := quests.GetQuestData(quest)
		{
			questDataJson, _ := json.MarshalIndent(questData, "", "  ")
			fmt.Println(string(questDataJson))
		}

	}
}
func CreateAndAmmendEntitlementQuest() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	ix, _ := quest_ops.CreateQuest(oracle.PublicKey(), questing.Quest{
		Index:           0,
		Name:            "aaa",
		Duration:        100,
		Oracle:          oracle.PublicKey(),
		WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
		Entitlement:     nil,
		Tender: &questing.Tender{
			MintAddress: MINT,
			Amount:      5,
		},
	})
	utils.SendTx(
		"list",
		append(make([]solana.Instruction, 0), ix),
		append(make([]solana.PrivateKey, 0), oracle),
		oracle.PublicKey(),
	)

	{
		questsPda, _ := quests.GetQuests(oracle.PublicKey())
		questsData := quests.GetQuestsData(questsPda)
		quest, _ := quests.GetQuest(oracle.PublicKey(), questsData.Quests-1)
		fmt.Println(quest, questsData.Quests)
		questData := quests.GetQuestData(quest)
		{
			questDataJson, _ := json.MarshalIndent(questData, "", "  ")
			fmt.Println(string(questDataJson))
		}

		{
			quest_ops.AmmendQuestWithEntitlement(
				oracle,
				*questData,
				questing.Reward{
					MintAddress: solana.MustPublicKeyFromBase58("ETZoY7cJfD8N7EVx5tShRYS1vxgv3F4Dkavjb52kGRyj"),
					Amount:      50,
				},
			)
		}
	}
	{
		questsPda, _ := quests.GetQuests(oracle.PublicKey())
		questsData := quests.GetQuestsData(questsPda)
		quest, _ := quests.GetQuest(oracle.PublicKey(), questsData.Quests-1)
		questData := quests.GetQuestData(quest)
		{
			questDataJson, _ := json.MarshalIndent(questData, "", "  ")
			fmt.Println(string(questDataJson))
		}
	}
}

func enrollQuestor() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	pixelBallzMint := solana.MustPublicKeyFromBase58("7zaCj11reNw4FMxY5UqR8mjNdatgB4vgdN17eKAwMGie")

	_, _ = oracle, pixelBallzMint
}

func startQuest() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	for range make([]int, 2) {
		var instructions []solana.Instruction

		pixelBallzMint := solana.NewWallet()
		pixelBallzTokenAddress, _ := utils.GetTokenWallet(oracle.PublicKey(), pixelBallzMint.PublicKey())
		// ballzMint := solana.MustPublicKeyFromBase58("6NGcNWFVksoeXf1xgAvKubQgS6rW5EZ2oVwqAa1eHySz")
		// ballzTokenAddress := solana.MustPublicKeyFromBase58("57hobyD843HjijKTLAbiKcfPCdBY3bdDPgvKR4ggoGaz")
		// pixelBallzMint := solana.MustPublicKeyFromBase58("7zaCj11reNw4FMxY5UqR8mjNdatgB4vgdN17eKAwMGie")
		// pixelBallTokenAddress := solana.MustPublicKeyFromBase58("DpXfu5sQpfGM2wSRPq1nUs4iKqVkjSwCehDrErhysZLP")
		{

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
		}
		utils.SendTx(
			"list",
			instructions,
			append(make([]solana.PrivateKey, 0), oracle, pixelBallzMint.PrivateKey),
			oracle.PublicKey(),
		)

		/*
			fmt.Println("sleeping")
			time.Sleep(15 * time.Second)
		*/

		questInstructions := make([]solana.Instruction, 0)

		questor, _ := quests.GetQuestorAccount(oracle.PublicKey())
		questorData := quests.GetQuestorData(questor)
		if questorData == nil {
			questInstructions = append(
				questInstructions,
				ops.EnrollQuestor(oracle.PublicKey()),
			)
		}

		questee, _ := quests.GetQuesteeAccount(pixelBallzMint.PublicKey())
		questeeData := quests.GetQuesteeData(questee)
		if questeeData == nil {
			questInstructions = append(
				questInstructions,
				ops.EnrollQuestee(oracle.PublicKey(), pixelBallzMint.PublicKey(), pixelBallzTokenAddress),
			)
		}

		questPda, _ := quests.GetQuest(oracle.PublicKey(), 3)
		questAccount, _ := quests.GetQuestAccount(questor, questee, questPda)
		questDeposit, _ := quests.GetQuestDepositTokenAccount(questee, questPda)

		questData := quests.GetQuestData(questPda)

		startQuestIx := questing.NewStartQuestInstructionBuilder().
			SetDepositTokenAccountAccount(questDeposit).
			SetInitializerAccount(oracle.PublicKey()).
			SetPixelballzMintAccount(pixelBallzMint.PublicKey()).
			SetPixelballzTokenAccountAccount(pixelBallzTokenAddress).
			SetQuestAccAccount(questAccount).
			SetQuestAccount(questPda).
			SetQuesteeAccount(questee).
			SetQuestorAccount(questor).
			SetRentAccount(solana.SysVarRentPubkey).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		if questData.Tender != nil && questData.TenderSplits != nil {
			tenderTokenAccount, _ := utils.GetTokenWallet(oracle.PublicKey(), questData.Tender.MintAddress)
			startQuestIx.Append(&solana.AccountMeta{tenderTokenAccount, true, false})
			for _, tenderSplit := range *questData.TenderSplits {
				startQuestIx.Append(&solana.AccountMeta{tenderSplit.TokenAddress, true, false})
			}
		}

		if err = startQuestIx.Validate(); err != nil {
			panic(err)
		} else {
			questInstructions = append(
				questInstructions,
				startQuestIx.Build(),
			)
		}

		utils.SendTx(
			"init cm",
			questInstructions,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
}

func startAndEndQuest() {
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}
	{
		questsPda, _ := quests.GetQuests(oracle.PublicKey())
		questsData := quests.GetQuestsData(questsPda)
		quest, _ := quests.GetQuest(oracle.PublicKey(), questsData.Quests-1)
		questData := quests.GetQuestData(quest)
		{
			questDataJson, _ := json.MarshalIndent(questData, "", "  ")
			fmt.Println(string(questDataJson))
		}
	}

	pixelBallzMint := solana.NewWallet()
	pixelBallzTokenAddress, _ := utils.GetTokenWallet(oracle.PublicKey(), pixelBallzMint.PublicKey())
	// ballzMint := solana.MustPublicKeyFromBase58("6NGcNWFVksoeXf1xgAvKubQgS6rW5EZ2oVwqAa1eHySz")
	// ballzTokenAddress := solana.MustPublicKeyFromBase58("57hobyD843HjijKTLAbiKcfPCdBY3bdDPgvKR4ggoGaz")
	// pixelBallzMint := solana.MustPublicKeyFromBase58("7zaCj11reNw4FMxY5UqR8mjNdatgB4vgdN17eKAwMGie")
	// pixelBallTokenAddress := solana.MustPublicKeyFromBase58("DpXfu5sQpfGM2wSRPq1nUs4iKqVkjSwCehDrErhysZLP")
	{
		var instructions []solana.Instruction
		{

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
		}
		utils.SendTx(
			"list",
			instructions,
			append(make([]solana.PrivateKey, 0), oracle, pixelBallzMint.PrivateKey),
			oracle.PublicKey(),
		)

		/*
		   fmt.Println("sleeping")
		   time.Sleep(15 * time.Second)
		*/

		questInstructions := make([]solana.Instruction, 0)

		questor, _ := quests.GetQuestorAccount(oracle.PublicKey())
		questorData := quests.GetQuestorData(questor)
		if questorData == nil {
			questInstructions = append(
				questInstructions,
				ops.EnrollQuestor(oracle.PublicKey()),
			)
		}

		questee, _ := quests.GetQuesteeAccount(pixelBallzMint.PublicKey())
		questeeData := quests.GetQuesteeData(questee)
		if questeeData == nil {
			questInstructions = append(
				questInstructions,
				ops.EnrollQuestee(oracle.PublicKey(), pixelBallzMint.PublicKey(), pixelBallzTokenAddress),
			)
		}

		questPda, _ := quests.GetQuest(oracle.PublicKey(), 0)
		questAccount, _ := quests.GetQuestAccount(questor, questee, questPda)
		questDeposit, _ := quests.GetQuestDepositTokenAccount(questee, questPda)

		questData := quests.GetQuestData(questPda)

		startQuestIx := questing.NewStartQuestInstructionBuilder().
			SetDepositTokenAccountAccount(questDeposit).
			SetInitializerAccount(oracle.PublicKey()).
			SetPixelballzMintAccount(pixelBallzMint.PublicKey()).
			SetPixelballzTokenAccountAccount(pixelBallzTokenAddress).
			SetQuestAccAccount(questAccount).
			SetQuestAccount(questPda).
			SetQuesteeAccount(questee).
			SetQuestorAccount(questor).
			SetRentAccount(solana.SysVarRentPubkey).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		if questData.Tender != nil && questData.TenderSplits != nil {
			tenderTokenAccount, _ := utils.GetTokenWallet(oracle.PublicKey(), questData.Tender.MintAddress)
			startQuestIx.Append(&solana.AccountMeta{PublicKey: tenderTokenAccount, IsWritable: true, IsSigner: false})
			for _, tenderSplit := range *questData.TenderSplits {
				startQuestIx.Append(&solana.AccountMeta{PublicKey: tenderSplit.TokenAddress, IsWritable: true, IsSigner: false})
			}
		}

		if err = startQuestIx.Validate(); err != nil {
			panic(err)
		} else {
			questInstructions = append(
				questInstructions,
				startQuestIx.Build(),
			)
		}

		utils.SendTx(
			"init cm",
			questInstructions,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
	fmt.Println("Sleeping...")
	time.Sleep(5 * time.Second)
	fmt.Println("Slept")
	{
		questInstructions := make([]solana.Instruction, 0)

		questor, _ := quests.GetQuestorAccount(oracle.PublicKey())

		questee, _ := quests.GetQuesteeAccount(pixelBallzMint.PublicKey())

		questPda, questPdaBump := quests.GetQuest(oracle.PublicKey(), 0)
		questAccount, _ := quests.GetQuestAccount(questor, questee, questPda)
		questDeposit, questDepositBump := quests.GetQuestDepositTokenAccount(questee, questPda)
		questQuesteeReceipt, _ := quests.GetQuestQuesteeReceiptAccount(questor, questee, questPda)

		endQuestIx := questing.NewEndQuestInstructionBuilder().
			SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID).
			SetDepositTokenAccountAccount(questDeposit).
			SetDepositTokenAccountBump(questDepositBump).
			SetInitializerAccount(oracle.PublicKey()).
			SetOracleAccount(oracle.PublicKey()).
			SetPixelballzMintAccount(pixelBallzMint.PublicKey()).
			SetPixelballzTokenAccountAccount(pixelBallzTokenAddress).
			SetQuestAccAccount(questAccount).
			SetQuestAccount(questPda).
			SetQuesteeAccount(questee).
			SetQuestorAccount(questor).
			SetQuestQuesteeReceiptAccount(questQuesteeReceipt).
			SetRentAccount(solana.SysVarRentPubkey).
			SetQuestBump(questPdaBump).
			SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetTokenProgramAccount(solana.TokenProgramID)

		questData := quests.GetQuestData(questPda)
		rewardMints := make([]solana.AccountMeta, 0)
		rewardAtas := make([]solana.AccountMeta, 0)
		fmt.Println("---------", len(endQuestIx.AccountMetaSlice))
		for _, reward := range questData.Rewards {
			endQuestIx.Append(&solana.AccountMeta{PublicKey: reward.MintAddress, IsWritable: true, IsSigner: false})
		}
		for _, reward := range questData.Rewards {
			rewardAta, _ := utils.GetTokenWallet(oracle.PublicKey(), reward.MintAddress)
			endQuestIx.Append(&solana.AccountMeta{PublicKey: rewardAta, IsWritable: true, IsSigner: false})
		}

		if err = endQuestIx.Validate(); err != nil {
			panic(err)
		} else {
			fmt.Println("---------", len(endQuestIx.AccountMetaSlice), len(rewardMints), len(rewardAtas))
			questInstructions = append(
				questInstructions,
				endQuestIx.Build(),
			)
		}

		utils.SendTx(
			"init cm",
			questInstructions,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}

}
