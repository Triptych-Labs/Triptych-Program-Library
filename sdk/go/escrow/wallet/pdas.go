package wallet

import (
	"github.com/gagliardetto/solana-go"
	"triptych.labs/escrow"
)

func GetEscrow(
	initializer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("escrow_wallet"),
			initializer.Bytes(),
		},
		escrow.ProgramID,
	)
	return addr, bump
}
