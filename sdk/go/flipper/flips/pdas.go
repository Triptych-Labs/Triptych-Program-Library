package flips

import (
	"encoding/binary"
	"time"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/flipper"
)

func GetHouse(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("house"),
			oracle.Bytes(),
		},
		flipper.ProgramID,
	)
	return addr, bump
}

func GetFlip(
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
		flipper.ProgramID,
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
		flipper.ProgramID,
	)
	return addr, bump
}
