// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EnableQuests is the `enableQuests` instruction.
type EnableQuests struct {

	// [0] = [WRITE, SIGNER] oracle
	//
	// [1] = [WRITE] quests
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEnableQuestsInstructionBuilder creates a new `EnableQuests` instruction builder.
func NewEnableQuestsInstructionBuilder() *EnableQuests {
	nd := &EnableQuests{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOracleAccount sets the "oracle" account.
func (inst *EnableQuests) SetOracleAccount(oracle ag_solanago.PublicKey) *EnableQuests {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(oracle).WRITE().SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *EnableQuests) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestsAccount sets the "quests" account.
func (inst *EnableQuests) SetQuestsAccount(quests ag_solanago.PublicKey) *EnableQuests {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(quests).WRITE()
	return inst
}

// GetQuestsAccount gets the "quests" account.
func (inst *EnableQuests) GetQuestsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *EnableQuests) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *EnableQuests {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *EnableQuests) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst EnableQuests) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EnableQuests,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EnableQuests) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EnableQuests) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Quests is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *EnableQuests) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EnableQuests")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       oracle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       quests", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj EnableQuests) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *EnableQuests) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewEnableQuestsInstruction declares a new EnableQuests instruction with the provided parameters and accounts.
func NewEnableQuestsInstruction(
	// Accounts:
	oracle ag_solanago.PublicKey,
	quests ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *EnableQuests {
	return NewEnableQuestsInstructionBuilder().
		SetOracleAccount(oracle).
		SetQuestsAccount(quests).
		SetSystemProgramAccount(systemProgram)
}
