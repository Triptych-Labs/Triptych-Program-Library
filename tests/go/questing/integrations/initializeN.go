package integrations

import (
	"fmt"
	"math"

	"triptych.labs/questing"
	"triptych.labs/questing/quests"
	quest_ops "triptych.labs/questing/quests/ops"
	"triptych.labs/utils"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type StakingQuestScope struct {
	name          string
	duration      int
	left          int
	right         int
	leftCreators  [5]solana.PublicKey
	rightCreators [5]solana.PublicKey
	yieldPer      int
	yieldPerTime  int
}

type RewardQuestScope struct {
	name          string
	duration      int
	left          int
	right         int
	leftCreators  [5]solana.PublicKey
	rightCreators [5]solana.PublicKey
	rewards       []questing.Reward
	tender        questing.Tender
	tenderSplits  []questing.Split
}

var GEN1 = [5]solana.PublicKey{
	solana.MustPublicKeyFromBase58("3riM3gFAvvGVWfLkbDT8CMrcnewqfmRiHYFUko2Gd4DB"),
	solana.MustPublicKeyFromBase58("3riM3gFAvvGVWfLkbDT8CMrcnewqfmRiHYFUko2Gd4DB"),
	solana.MustPublicKeyFromBase58("3riM3gFAvvGVWfLkbDT8CMrcnewqfmRiHYFUko2Gd4DB"),
	solana.MustPublicKeyFromBase58("3riM3gFAvvGVWfLkbDT8CMrcnewqfmRiHYFUko2Gd4DB"),
	solana.MustPublicKeyFromBase58("3riM3gFAvvGVWfLkbDT8CMrcnewqfmRiHYFUko2Gd4DB"),
}

var GEN2 = [5]solana.PublicKey{
	solana.MustPublicKeyFromBase58("6oVAspyLfV7iWYivvHokcXg9X5LcCLcWvsa7XL1rbEM8"),
	solana.MustPublicKeyFromBase58("6oVAspyLfV7iWYivvHokcXg9X5LcCLcWvsa7XL1rbEM8"),
	solana.MustPublicKeyFromBase58("6oVAspyLfV7iWYivvHokcXg9X5LcCLcWvsa7XL1rbEM8"),
	solana.MustPublicKeyFromBase58("6oVAspyLfV7iWYivvHokcXg9X5LcCLcWvsa7XL1rbEM8"),
	solana.MustPublicKeyFromBase58("6oVAspyLfV7iWYivvHokcXg9X5LcCLcWvsa7XL1rbEM8"),
}

func CreateNStakingQuests() {

	scopes := []StakingQuestScope{
		{
			name:          "Gen1",
			duration:      0,
			left:          1,
			right:         0,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			yieldPer:      50,
			yieldPerTime:  10,
		},
		{
			name:          "Gen2",
			duration:      0,
			left:          0,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			yieldPer:      50,
			yieldPerTime:  10,
		},
		{
			name:          "Gen1+Gen2 Duo",
			duration:      0,
			left:          1,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			yieldPer:      150,
			yieldPerTime:  10,
		},
		{
			name:          "Gen1+Gen2 Quatro",
			duration:      0,
			left:          2,
			right:         2,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			yieldPer:      350,
			yieldPerTime:  10,
		},
	}

	createNStakingQuests(scopes)
}

func CreateNRewardQuests() {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	ixs := make([]solana.Instruction, 0)
	questRewardIx, questRewardMint := quest_ops.RegisterQuestsStakingReward(oracle.PublicKey(), "qstNBA WL", "qstNBAWL")
	ixs = append(ixs, questRewardIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle, questRewardMint),
		oracle.PublicKey(),
	)

	tenderMint := solana.MustPublicKeyFromBase58("3tQsckZ7R9Tec2rGghGL8CD8RGHTMcoZ8XRQRJnwujGr")

	tenderMintMeta := utils.GetTokenMintData(rpcClient, tenderMint)

	scopes := []RewardQuestScope{
		{
			name:          "60% Whitelist",
			duration:      60 * 60 * 24,
			left:          1,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			rewards: []questing.Reward{
				{
					MintAddress:   questRewardMint.PublicKey(),
					Threshold:     60,
					Amount:        10,
					AuthorityEnum: 0,
					Cap:           math.MaxInt64,
					Counter:       0,
				},
			},
			tender: questing.Tender{
				MintAddress: tenderMint,
				Amount:      utils.ConvertUiAmountToAmount(float64(10), tenderMintMeta.Decimals),
			},
			tenderSplits: []questing.Split{
				{
					TokenAddress: solana.PublicKey{},
					OpCode:       0,
					Share:        100,
				},
			},
		},
		{
			name:          "100% Whitelist",
			duration:      60 * 60 * 24,
			left:          1,
			right:         1,
			leftCreators:  GEN1,
			rightCreators: GEN2,
			rewards: []questing.Reward{
				{
					MintAddress:   questRewardMint.PublicKey(),
					Threshold:     100,
					Amount:        10,
					AuthorityEnum: 0,
					Cap:           math.MaxInt64,
					Counter:       0,
				},
			},
			tender: questing.Tender{
				MintAddress: tenderMint,
				Amount:      utils.ConvertUiAmountToAmount(float64(20), tenderMintMeta.Decimals),
			},
			tenderSplits: []questing.Split{
				{
					TokenAddress: solana.PublicKey{},
					OpCode:       0,
					Share:        100,
				},
			},
		},
	}

	createNRewardQuests(scopes)
}

func createNStakingQuests(scopes []StakingQuestScope) {
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	ixs := make([]solana.Instruction, 0)
	stakingRewardIx, stakingMint := quest_ops.RegisterQuestsStakingReward(oracle.PublicKey(), "qstCoin", "QSTC")
	ixs = append(ixs, stakingRewardIx)

	utils.SendTx(
		"list",
		ixs,
		append(make([]solana.PrivateKey, 0), oracle, stakingMint),
		oracle.PublicKey(),
	)

	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	for i, scope := range scopes {
		questIxs := make([]solana.Instruction, 0)
		questData := questing.Quest{
			Index:           questsData.Quests + uint64(i),
			Name:            scope.name,
			Duration:        int64(scope.duration),
			Oracle:          oracle.PublicKey(),
			WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
			Tender:          nil,
			TenderSplits:    nil,
			Rewards:         []questing.Reward{},
			StakingConfig: &questing.StakingConfig{
				MintAddress:  stakingMint.PublicKey(),
				YieldPer:     uint64(scope.yieldPer),     // 10 secounds
				YieldPerTime: uint64(scope.yieldPerTime), // 5 tokens
			},
			PairsConfig: &questing.PairsConfig{
				Left:          uint8(scope.left),
				LeftCreators:  scope.leftCreators,
				Right:         uint8(scope.right),
				RightCreators: scope.rightCreators,
			},
		}

		creationIx, _ := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData)
		questIxs = append(questIxs, creationIx)

		utils.SendTx(
			"list",
			questIxs,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
}

func createNRewardQuests(scopes []RewardQuestScope) {
	// CTM8npagWrtdi85aYix3kpD23yKdboPFMXk9fPWMBoD7
	rpcClient := rpc.New(utils.NETWORK)
	oracle, err := solana.PrivateKeyFromSolanaKeygenFile("./oracle.key")
	if err != nil {
		panic(err)
	}

	questsPda, _ := quests.GetQuests(oracle.PublicKey())
	questsData := quests.GetQuestsData(rpcClient, questsPda)
	for i, scope := range scopes {
		questIxs := make([]solana.Instruction, 0)
		questData := questing.Quest{
			Index:           questsData.Quests + uint64(i),
			Name:            scope.name,
			Duration:        int64(scope.duration),
			Oracle:          oracle.PublicKey(),
			WlCandyMachines: []solana.PublicKey{oracle.PublicKey()},
			Tender:          &scope.tender,
			TenderSplits:    &scope.tenderSplits,
			Rewards:         scope.rewards,
			StakingConfig:   nil,
			PairsConfig: &questing.PairsConfig{
				Left:          uint8(scope.left),
				LeftCreators:  scope.leftCreators,
				Right:         uint8(scope.right),
				RightCreators: scope.rightCreators,
			},
		}
		fmt.Println(scope.tender, scope.tenderSplits)

		creationIx, _ := quest_ops.CreateQuest(rpcClient, oracle.PublicKey(), questData)
		questIxs = append(questIxs, creationIx)

		utils.SendTx(
			"list",
			questIxs,
			append(make([]solana.PrivateKey, 0), oracle),
			oracle.PublicKey(),
		)
	}
}
