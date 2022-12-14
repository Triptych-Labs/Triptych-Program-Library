// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package draffle

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BuyTickets is the `buyTickets` instruction.
type BuyTickets struct {
	Amount *uint32

	// [0] = [] raffle
	//
	// [1] = [WRITE] entrants
	//
	// [2] = [WRITE] proceeds
	//
	// [3] = [WRITE] buyerTokenAccount
	//
	// [4] = [WRITE, SIGNER] buyerTransferAuthority
	//
	// [5] = [WRITE] feeAcc
	//
	// [6] = [] tokenProgram
	//
	// [7] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyTicketsInstructionBuilder creates a new `BuyTickets` instruction builder.
func NewBuyTicketsInstructionBuilder() *BuyTickets {
	nd := &BuyTickets{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *BuyTickets) SetAmount(amount uint32) *BuyTickets {
	inst.Amount = &amount
	return inst
}

// SetRaffleAccount sets the "raffle" account.
func (inst *BuyTickets) SetRaffleAccount(raffle ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(raffle)
	return inst
}

// GetRaffleAccount gets the "raffle" account.
func (inst *BuyTickets) GetRaffleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetEntrantsAccount sets the "entrants" account.
func (inst *BuyTickets) SetEntrantsAccount(entrants ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(entrants).WRITE()
	return inst
}

// GetEntrantsAccount gets the "entrants" account.
func (inst *BuyTickets) GetEntrantsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProceedsAccount sets the "proceeds" account.
func (inst *BuyTickets) SetProceedsAccount(proceeds ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(proceeds).WRITE()
	return inst
}

// GetProceedsAccount gets the "proceeds" account.
func (inst *BuyTickets) GetProceedsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBuyerTokenAccountAccount sets the "buyerTokenAccount" account.
func (inst *BuyTickets) SetBuyerTokenAccountAccount(buyerTokenAccount ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(buyerTokenAccount).WRITE()
	return inst
}

// GetBuyerTokenAccountAccount gets the "buyerTokenAccount" account.
func (inst *BuyTickets) GetBuyerTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBuyerTransferAuthorityAccount sets the "buyerTransferAuthority" account.
func (inst *BuyTickets) SetBuyerTransferAuthorityAccount(buyerTransferAuthority ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(buyerTransferAuthority).WRITE().SIGNER()
	return inst
}

// GetBuyerTransferAuthorityAccount gets the "buyerTransferAuthority" account.
func (inst *BuyTickets) GetBuyerTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFeeAccAccount sets the "feeAcc" account.
func (inst *BuyTickets) SetFeeAccAccount(feeAcc ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(feeAcc).WRITE()
	return inst
}

// GetFeeAccAccount gets the "feeAcc" account.
func (inst *BuyTickets) GetFeeAccAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *BuyTickets) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *BuyTickets) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *BuyTickets) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *BuyTickets {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *BuyTickets) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst BuyTickets) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BuyTickets,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BuyTickets) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BuyTickets) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Raffle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Entrants is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Proceeds is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BuyerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BuyerTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FeeAcc is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *BuyTickets) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BuyTickets")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                raffle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              entrants", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              proceeds", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            buyerToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("buyerTransferAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                feeAcc", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj BuyTickets) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BuyTickets) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewBuyTicketsInstruction declares a new BuyTickets instruction with the provided parameters and accounts.
func NewBuyTicketsInstruction(
	// Parameters:
	amount uint32,
	// Accounts:
	raffle ag_solanago.PublicKey,
	entrants ag_solanago.PublicKey,
	proceeds ag_solanago.PublicKey,
	buyerTokenAccount ag_solanago.PublicKey,
	buyerTransferAuthority ag_solanago.PublicKey,
	feeAcc ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *BuyTickets {
	return NewBuyTicketsInstructionBuilder().
		SetAmount(amount).
		SetRaffleAccount(raffle).
		SetEntrantsAccount(entrants).
		SetProceedsAccount(proceeds).
		SetBuyerTokenAccountAccount(buyerTokenAccount).
		SetBuyerTransferAuthorityAccount(buyerTransferAuthority).
		SetFeeAccAccount(feeAcc).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram)
}
