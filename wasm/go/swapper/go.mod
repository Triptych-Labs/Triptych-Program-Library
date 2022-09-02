module triptych.labs/wasm/v2

go 1.16

require (
	triptych.labs/swapper v0.0.0
	triptych.labs/utils v0.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gagliardetto/binary v0.6.1
	github.com/gagliardetto/gofuzz v1.2.2 // indirect
	github.com/gagliardetto/metaplex-go v0.2.1
	github.com/gagliardetto/solana-go v1.4.0
	github.com/gagliardetto/treeout v0.1.4 // indirect
)

replace triptych.labs/swapper => ../../../sdk/go/swapper

replace triptych.labs/utils => ../../../sdk/go/utils
