package game

import (
	"encoding/binary"
	"time"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/blackjack"
)

func GetHouse(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("house"),
			oracle.Bytes(),
		},
		blackjack.ProgramID,
	)
	return addr, bump
}

func GetGames(
	initializer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("game"),
			initializer.Bytes(),
		},
		blackjack.ProgramID,
	)
	return addr, bump
}

func GetGame(
	initializer solana.PublicKey,
	index uint64,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("game"),
			initializer.Bytes(),
			buf,
		},
		blackjack.ProgramID,
	)
	return addr, bump
}

func GetStats(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8, uint64) {
	buf := make([]byte, 8)
	t := time.Now().UTC()
	dailyEpoch := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	binary.LittleEndian.PutUint64(buf, uint64(dailyEpoch.Unix()))
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("flip"),
			oracle.Bytes(),
			buf,
		},
		blackjack.ProgramID,
	)
	return addr, bump, uint64(dailyEpoch.Unix())
}

func GetEscrow(
	initializer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("escrow"),
			initializer.Bytes(),
		},
		blackjack.ProgramID,
	)
	return addr, bump
}
