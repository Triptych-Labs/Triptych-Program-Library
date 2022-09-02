module triptych.labs/swapper_tests/v2

go 1.16

require (
	github.com/btcsuite/btcutil v1.0.2
	github.com/gagliardetto/binary v0.6.1 // indirect
	github.com/gagliardetto/metaplex-go v0.2.1
	github.com/gagliardetto/solana-go v1.4.0
	github.com/mr-tron/base58 v1.2.0 // indirect
	triptych.labs/swapper v0.0.0
	triptych.labs/utils v0.0.0
)

replace triptych.labs/swapper => ../../../sdk/go/swapper

replace triptych.labs/utils => ../../../sdk/go/utils
