package ops

import (
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"triptych.labs/swapper"
	"triptych.labs/swapper/swaps"
	"triptych.labs/utils"
)

func InvokeSwap(rpcClient *rpc.Client, oracle, initializer solana.PublicKey, swapIndex uint64, amount float64) solana.Instruction {
	swapPda, swapPdaBump := swaps.GetSwap(oracle, swapIndex)
	swapData := swaps.GetSwapData(rpcClient, swapPda)

	swapRecorder, swapRecorderBump := swaps.GetSwapRecorder(swapData.Oracle)
	log.Println("...", oracle, initializer, swapData.Oracle)
	swapPool, _ := swaps.GetSwapPool(swapData.ToMint, swapRecorder)

	fromTokenAccount, _ := utils.GetTokenWallet(initializer, swapData.FromMint)
	toTokenAccount, _ := utils.GetTokenWallet(initializer, swapData.ToMint)

	fromTokenMintMeta := utils.GetTokenMintData(rpcClient, swapData.FromMint)

	amountFmt := utils.ConvertUiAmountToAmount(amount, fromTokenMintMeta.Decimals)

	swapIx := swapper.NewInvokeSwapInstructionBuilder().
		SetAmount(amountFmt).
		SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID).
		SetFromMintAccount(swapData.FromMint).
		SetFromTokenAccountAccount(fromTokenAccount).
		SetInitializerAccount(initializer).
		SetOracleAccount(swapData.Oracle).
		SetRentAccount(solana.SysVarRentPubkey).
		SetSwapAccount(swapPda).
		SetSwapBump(swapPdaBump).
		SetSwapIndex(swapIndex).
		SetSwapPoolAccount(swapPool).
		SetSwapRecorderAccount(swapRecorder).
		SetSwapRecorderBump(swapRecorderBump).
		SetSystemProgramAccount(solana.SystemProgramID).
		SetToMintAccount(swapData.ToMint).
		SetToTokenAccountAccount(toTokenAccount).
		SetTokenProgramAccount(solana.TokenProgramID)

	return swapIx.Build()
}
