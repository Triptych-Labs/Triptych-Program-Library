module triptych.labs/blackjack

go 1.16

require (
	github.com/btcsuite/btcutil v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/gagliardetto/binary v0.6.1
	github.com/gagliardetto/gofuzz v1.2.2
	github.com/gagliardetto/solana-go v1.4.0
	github.com/gagliardetto/treeout v0.1.4
	github.com/stretchr/testify v1.7.1
	triptych.labs/escrow v0.0.0
	triptych.labs/utils v0.0.0 // indirect
)

replace triptych.labs/utils => ../utils

replace triptych.labs/escrow => ../escrow
