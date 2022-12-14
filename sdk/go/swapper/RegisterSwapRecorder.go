// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package swapper

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RegisterSwapRecorder is the `registerSwapRecorder` instruction.
type RegisterSwapRecorder struct {

	// [0] = [WRITE, SIGNER] oracle
	//
	// [1] = [WRITE] swapRecorder
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRegisterSwapRecorderInstructionBuilder creates a new `RegisterSwapRecorder` instruction builder.
func NewRegisterSwapRecorderInstructionBuilder() *RegisterSwapRecorder {
	nd := &RegisterSwapRecorder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOracleAccount sets the "oracle" account.
func (inst *RegisterSwapRecorder) SetOracleAccount(oracle ag_solanago.PublicKey) *RegisterSwapRecorder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).WRITE().SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *RegisterSwapRecorder) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSwapRecorderAccount sets the "swapRecorder" account.
func (inst *RegisterSwapRecorder) SetSwapRecorderAccount(swapRecorder ag_solanago.PublicKey) *RegisterSwapRecorder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(swapRecorder).WRITE()
	return inst
}

// GetSwapRecorderAccount gets the "swapRecorder" account.
func (inst *RegisterSwapRecorder) GetSwapRecorderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RegisterSwapRecorder) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RegisterSwapRecorder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RegisterSwapRecorder) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst RegisterSwapRecorder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RegisterSwapRecorder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RegisterSwapRecorder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RegisterSwapRecorder) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SwapRecorder is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RegisterSwapRecorder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RegisterSwapRecorder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" swapRecorder", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj RegisterSwapRecorder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RegisterSwapRecorder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRegisterSwapRecorderInstruction declares a new RegisterSwapRecorder instruction with the provided parameters and accounts.
func NewRegisterSwapRecorderInstruction(
	// Accounts:
	oracle ag_solanago.PublicKey,
	swapRecorder ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RegisterSwapRecorder {
	return NewRegisterSwapRecorderInstructionBuilder().
		SetOracleAccount(oracle).
		SetSwapRecorderAccount(swapRecorder).
		SetSystemProgramAccount(systemProgram)
}
