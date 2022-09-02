package ops

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/blackjack"
	"triptych.labs/blackjack/game"
	"triptych.labs/escrow"
	"triptych.labs/escrow/wallet"
)

func NewGame(
	rpcClient *rpc.Client,
	oracle,
	initializer solana.PublicKey,
	amount uint64,
	operator string,
) *blackjack.Instruction {
	escrowPda, escrowPdaBump := wallet.GetEscrow(initializer)
	housePda, housePdaBump := game.GetHouse(oracle)
	gamesPda, _ := game.GetGames(initializer)

	gameIndex := uint64(0)
	if gamesData := game.GetGamesData(rpcClient, gamesPda); gamesData != nil {
		gameIndex = gamesData.Games
	}

	gamePda, _ := game.GetGame(initializer, gameIndex)
	statsPda, _, dailyEpoch := game.GetStats(oracle)
	fmt.Println(gamesPda, gamePda, housePda)

	newGameIx := blackjack.NewStartGameInstructionBuilder().
		SetAmount(amount).
		SetBlackjackProgramAccount(blackjack.ProgramID).
		SetDailyEpoch(dailyEpoch).
		SetEscrowAccount(escrowPda).
		SetEscrowBump(escrowPdaBump).
		SetEscrowProgramAccount(escrow.ProgramID).
		SetGameAccount(gamePda).
		SetGamesAccount(gamesPda).
		SetHouseAccount(housePda).
		SetHouseBump(housePdaBump).
		SetInitializerAccount(initializer).
		SetOracleAccount(oracle).
		SetSlotHashesAccount(solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")).
		SetStatsAccount(statsPda).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetWallet("wallet")

	return newGameIx.Build()
}
