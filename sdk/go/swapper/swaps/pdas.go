package swaps

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/swapper"
)

func GetSwapRecorder(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("swap_recorder"),
			oracle.Bytes(),
		},
		swapper.ProgramID,
	)
	return addr, bump
}

func GetSwap(
	oracle solana.PublicKey,
	index uint64,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			oracle.Bytes(),
			buf,
		},
		swapper.ProgramID,
	)
	return addr, bump
}

func GetSwapPool(
	mint solana.PublicKey,
	swapRecorder solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("swaping"),
			mint.Bytes(),
			swapRecorder.Bytes(),
		},
		swapper.ProgramID,
	)
	return addr, bump
}
