package quests

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"triptych.labs/questing"
)

func GetQuests(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("oracle"),
			oracle.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuest(
	oracle solana.PublicKey,
	index uint64,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("oracle"),
			oracle.Bytes(),
			buf,
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestEntitlementTokenAccount(
	oracle solana.PublicKey,
	index uint64,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("oracle"),
			[]byte("entitlement"),
			oracle.Bytes(),
			buf,
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestDepositTokenAccount(
	pixelBallzMint solana.PublicKey,
	quest solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			pixelBallzMint.Bytes(),
			quest.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestAccount(
	questor solana.PublicKey,
	questProposal solana.PublicKey,
	quest solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			questor.Bytes(),
			questProposal.Bytes(),
			quest.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestorAccount(
	oracle solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			oracle.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuesteeAccount(
	pixelBallzMint solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("questing"),
			pixelBallzMint.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestQuesteeReceiptAccount(
	questor solana.PublicKey,
	questee solana.PublicKey,
	quest solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("quest_reward"),
			questor.Bytes(),
			questee.Bytes(),
			quest.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetRewardToken(
	oracle solana.PublicKey,
	mint solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("oracle"),
			oracle.Bytes(),
			mint.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestRecorder(
	quest solana.PublicKey,
	initializer solana.PublicKey,
) (solana.PublicKey, uint8) {
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			[]byte("quest_recorder"),
			quest.Bytes(),
			initializer.Bytes(),
		},
		questing.ProgramID,
	)
	return addr, bump
}

func GetQuestProposal(
	quest solana.PublicKey,
	initializer solana.PublicKey,
	proposalId uint64,
) (solana.PublicKey, uint8) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, proposalId)
	addr, bump, _ := solana.FindProgramAddress(
		[][]byte{
			quest.Bytes(),
			initializer.Bytes(),
			buf,
		},
		questing.ProgramID,
	)
	return addr, bump
}
