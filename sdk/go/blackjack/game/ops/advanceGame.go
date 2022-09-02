package ops

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/blackjack"
	"triptych.labs/blackjack/game"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
)

func AdvanceGame(
	rpcClient *rpc.Client,
	oracle,
	initializer solana.PublicKey,
	operator string,
	gameIndex uint64,
) *blackjack.Instruction {
	escrowPda, escrowPdaBump := wallet.GetEscrow(initializer)
	housePda, housePdaBump := game.GetHouse(oracle)

	gamePda, gamePdaBump := game.GetGame(initializer, gameIndex)
	statsPda, statsPdaBump, dailyEpoch := game.GetStats(oracle)

	advGameIx := blackjack.NewPlayerTurnInstructionBuilder().
		SetBlackjackProgramAccount(blackjack.ProgramID).
		SetDailyEpoch(dailyEpoch).
		SetEscrowAccount(escrowPda).
		SetEscrowBump(escrowPdaBump).
		SetEscrowProgramAccount(escrow.ProgramID).
		SetGameAccount(gamePda).
		SetGameBump(gamePdaBump).
		SetGameIndex(gameIndex).
		SetHouseAccount(housePda).
		SetHouseBump(housePdaBump).
		SetInitializerAccount(initializer).
		SetOperator(operator).
		SetOracleAccount(oracle).
		SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")).
		SetStatsAccount(statsPda).
		SetStatsBump(statsPdaBump).
		SetSystemProgramAccount(solana.SystemProgramID)

	if err := advGameIx.Validate(); err != nil {
		panic(err)
	}

	return advGameIx.Build()
}
