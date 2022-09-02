// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package flipper

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateFlip is the `createFlip` instruction.
type CreateFlip struct {

	// [0] = [WRITE, SIGNER] oracle
	//
	// [1] = [WRITE] house
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateFlipInstructionBuilder creates a new `CreateFlip` instruction builder.
func NewCreateFlipInstructionBuilder() *CreateFlip {
	nd := &CreateFlip{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOracleAccount sets the "oracle" account.
func (inst *CreateFlip) SetOracleAccount(oracle ag_solanago.PublicKey) *CreateFlip {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).WRITE().SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *CreateFlip) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetHouseAccount sets the "house" account.
func (inst *CreateFlip) SetHouseAccount(house ag_solanago.PublicKey) *CreateFlip {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(house).WRITE()
	return inst
}

// GetHouseAccount gets the "house" account.
func (inst *CreateFlip) GetHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateFlip) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateFlip {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateFlip) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CreateFlip) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateFlip,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateFlip) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateFlip) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.House is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateFlip) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateFlip")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        house", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CreateFlip) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CreateFlip) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCreateFlipInstruction declares a new CreateFlip instruction with the provided parameters and accounts.
func NewCreateFlipInstruction(
	// Accounts:
	oracle ag_solanago.PublicKey,
	house ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateFlip {
	return NewCreateFlipInstructionBuilder().
		SetOracleAccount(oracle).
		SetHouseAccount(house).
		SetSystemProgramAccount(systemProgram)
}