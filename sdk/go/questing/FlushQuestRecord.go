// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FlushQuestRecord is the `flushQuestRecord` instruction.
type FlushQuestRecord struct {
	QuestProposalIndex *uint64
	QuestProposalBump  *uint8
	QuestBump          *uint8

	// [0] = [WRITE] questProposal
	//
	// [1] = [WRITE] quest
	//
	// [2] = [WRITE, SIGNER] initializer
	//
	// [3] = [WRITE] pixelballzMint
	//
	// [4] = [WRITE] pixelballzTokenAccount
	//
	// [5] = [] tokenProgram
	//
	// [6] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFlushQuestRecordInstructionBuilder creates a new `FlushQuestRecord` instruction builder.
func NewFlushQuestRecordInstructionBuilder() *FlushQuestRecord {
	nd := &FlushQuestRecord{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetQuestProposalIndex sets the "questProposalIndex" parameter.
func (inst *FlushQuestRecord) SetQuestProposalIndex(questProposalIndex uint64) *FlushQuestRecord {
	inst.QuestProposalIndex = &questProposalIndex
	return inst
}

// SetQuestProposalBump sets the "questProposalBump" parameter.
func (inst *FlushQuestRecord) SetQuestProposalBump(questProposalBump uint8) *FlushQuestRecord {
	inst.QuestProposalBump = &questProposalBump
	return inst
}

// SetQuestBump sets the "questBump" parameter.
func (inst *FlushQuestRecord) SetQuestBump(questBump uint8) *FlushQuestRecord {
	inst.QuestBump = &questBump
	return inst
}

// SetQuestProposalAccount sets the "questProposal" account.
func (inst *FlushQuestRecord) SetQuestProposalAccount(questProposal ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(questProposal).WRITE()
	return inst
}

// GetQuestProposalAccount gets the "questProposal" account.
func (inst *FlushQuestRecord) GetQuestProposalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestAccount sets the "quest" account.
func (inst *FlushQuestRecord) SetQuestAccount(quest ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(quest).WRITE()
	return inst
}

// GetQuestAccount gets the "quest" account.
func (inst *FlushQuestRecord) GetQuestAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetInitializerAccount sets the "initializer" account.
func (inst *FlushQuestRecord) SetInitializerAccount(initializer ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *FlushQuestRecord) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPixelballzMintAccount sets the "pixelballzMint" account.
func (inst *FlushQuestRecord) SetPixelballzMintAccount(pixelballzMint ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(pixelballzMint).WRITE()
	return inst
}

// GetPixelballzMintAccount gets the "pixelballzMint" account.
func (inst *FlushQuestRecord) GetPixelballzMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPixelballzTokenAccountAccount sets the "pixelballzTokenAccount" account.
func (inst *FlushQuestRecord) SetPixelballzTokenAccountAccount(pixelballzTokenAccount ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(pixelballzTokenAccount).WRITE()
	return inst
}

// GetPixelballzTokenAccountAccount gets the "pixelballzTokenAccount" account.
func (inst *FlushQuestRecord) GetPixelballzTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *FlushQuestRecord) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *FlushQuestRecord) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *FlushQuestRecord) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *FlushQuestRecord {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *FlushQuestRecord) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst FlushQuestRecord) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FlushQuestRecord,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FlushQuestRecord) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FlushQuestRecord) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.QuestProposalIndex == nil {
			return errors.New("QuestProposalIndex parameter is not set")
		}
		if inst.QuestProposalBump == nil {
			return errors.New("QuestProposalBump parameter is not set")
		}
		if inst.QuestBump == nil {
			return errors.New("QuestBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.QuestProposal is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Quest is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PixelballzMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PixelballzTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *FlushQuestRecord) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FlushQuestRecord")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("QuestProposalIndex", *inst.QuestProposalIndex))
						paramsBranch.Child(ag_format.Param(" QuestProposalBump", *inst.QuestProposalBump))
						paramsBranch.Child(ag_format.Param("         QuestBump", *inst.QuestBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  questProposal", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          quest", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    initializer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" pixelballzMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("pixelballzToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj FlushQuestRecord) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `QuestProposalIndex` param:
	err = encoder.Encode(obj.QuestProposalIndex)
	if err != nil {
		return err
	}
	// Serialize `QuestProposalBump` param:
	err = encoder.Encode(obj.QuestProposalBump)
	if err != nil {
		return err
	}
	// Serialize `QuestBump` param:
	err = encoder.Encode(obj.QuestBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FlushQuestRecord) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `QuestProposalIndex`:
	err = decoder.Decode(&obj.QuestProposalIndex)
	if err != nil {
		return err
	}
	// Deserialize `QuestProposalBump`:
	err = decoder.Decode(&obj.QuestProposalBump)
	if err != nil {
		return err
	}
	// Deserialize `QuestBump`:
	err = decoder.Decode(&obj.QuestBump)
	if err != nil {
		return err
	}
	return nil
}

// NewFlushQuestRecordInstruction declares a new FlushQuestRecord instruction with the provided parameters and accounts.
func NewFlushQuestRecordInstruction(
	// Parameters:
	questProposalIndex uint64,
	questProposalBump uint8,
	questBump uint8,
	// Accounts:
	questProposal ag_solanago.PublicKey,
	quest ag_solanago.PublicKey,
	initializer ag_solanago.PublicKey,
	pixelballzMint ag_solanago.PublicKey,
	pixelballzTokenAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *FlushQuestRecord {
	return NewFlushQuestRecordInstructionBuilder().
		SetQuestProposalIndex(questProposalIndex).
		SetQuestProposalBump(questProposalBump).
		SetQuestBump(questBump).
		SetQuestProposalAccount(questProposal).
		SetQuestAccount(quest).
		SetInitializerAccount(initializer).
		SetPixelballzMintAccount(pixelballzMint).
		SetPixelballzTokenAccountAccount(pixelballzTokenAccount).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram)
}