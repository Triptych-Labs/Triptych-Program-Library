// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EndQuest is the `endQuest` instruction.
type EndQuest struct {
	DepositTokenAccountBump *uint8
	QuestsBump              *uint8

	// [0] = [WRITE] questQuesteeReceipt
	//
	// [1] = [WRITE] quest
	//
	// [2] = [] quests
	//
	// [3] = [WRITE, SIGNER] initializer
	//
	// [4] = [WRITE] depositTokenAccount
	//
	// [5] = [WRITE] pixelballzMint
	//
	// [6] = [WRITE] pixelballzTokenAccount
	//
	// [7] = [WRITE] questAcc
	//
	// [8] = [WRITE] questor
	//
	// [9] = [WRITE] questee
	//
	// [10] = [] oracle
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] associatedTokenProgram
	//
	// [13] = [] slotHashes
	//
	// [14] = [] systemProgram
	//
	// [15] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEndQuestInstructionBuilder creates a new `EndQuest` instruction builder.
func NewEndQuestInstructionBuilder() *EndQuest {
	nd := &EndQuest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetDepositTokenAccountBump sets the "depositTokenAccountBump" parameter.
func (inst *EndQuest) SetDepositTokenAccountBump(depositTokenAccountBump uint8) *EndQuest {
	inst.DepositTokenAccountBump = &depositTokenAccountBump
	return inst
}

// SetQuestsBump sets the "questsBump" parameter.
func (inst *EndQuest) SetQuestsBump(questsBump uint8) *EndQuest {
	inst.QuestsBump = &questsBump
	return inst
}

// SetQuestQuesteeReceiptAccount sets the "questQuesteeReceipt" account.
func (inst *EndQuest) SetQuestQuesteeReceiptAccount(questQuesteeReceipt ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(questQuesteeReceipt).WRITE()
	return inst
}

// GetQuestQuesteeReceiptAccount gets the "questQuesteeReceipt" account.
func (inst *EndQuest) GetQuestQuesteeReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuestAccount sets the "quest" account.
func (inst *EndQuest) SetQuestAccount(quest ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(quest).WRITE()
	return inst
}

// GetQuestAccount gets the "quest" account.
func (inst *EndQuest) GetQuestAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetQuestsAccount sets the "quests" account.
func (inst *EndQuest) SetQuestsAccount(quests ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(quests)
	return inst
}

// GetQuestsAccount gets the "quests" account.
func (inst *EndQuest) GetQuestsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetInitializerAccount sets the "initializer" account.
func (inst *EndQuest) SetInitializerAccount(initializer ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *EndQuest) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetDepositTokenAccountAccount sets the "depositTokenAccount" account.
func (inst *EndQuest) SetDepositTokenAccountAccount(depositTokenAccount ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(depositTokenAccount).WRITE()
	return inst
}

// GetDepositTokenAccountAccount gets the "depositTokenAccount" account.
func (inst *EndQuest) GetDepositTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPixelballzMintAccount sets the "pixelballzMint" account.
func (inst *EndQuest) SetPixelballzMintAccount(pixelballzMint ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(pixelballzMint).WRITE()
	return inst
}

// GetPixelballzMintAccount gets the "pixelballzMint" account.
func (inst *EndQuest) GetPixelballzMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPixelballzTokenAccountAccount sets the "pixelballzTokenAccount" account.
func (inst *EndQuest) SetPixelballzTokenAccountAccount(pixelballzTokenAccount ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(pixelballzTokenAccount).WRITE()
	return inst
}

// GetPixelballzTokenAccountAccount gets the "pixelballzTokenAccount" account.
func (inst *EndQuest) GetPixelballzTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetQuestAccAccount sets the "questAcc" account.
func (inst *EndQuest) SetQuestAccAccount(questAcc ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(questAcc).WRITE()
	return inst
}

// GetQuestAccAccount gets the "questAcc" account.
func (inst *EndQuest) GetQuestAccAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetQuestorAccount sets the "questor" account.
func (inst *EndQuest) SetQuestorAccount(questor ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(questor).WRITE()
	return inst
}

// GetQuestorAccount gets the "questor" account.
func (inst *EndQuest) GetQuestorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetQuesteeAccount sets the "questee" account.
func (inst *EndQuest) SetQuesteeAccount(questee ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(questee).WRITE()
	return inst
}

// GetQuesteeAccount gets the "questee" account.
func (inst *EndQuest) GetQuesteeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOracleAccount sets the "oracle" account.
func (inst *EndQuest) SetOracleAccount(oracle ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *EndQuest) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *EndQuest) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *EndQuest) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *EndQuest) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *EndQuest) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *EndQuest) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *EndQuest) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *EndQuest) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *EndQuest) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetRentAccount sets the "rent" account.
func (inst *EndQuest) SetRentAccount(rent ag_solanago.PublicKey) *EndQuest {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *EndQuest) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst EndQuest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EndQuest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EndQuest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EndQuest) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DepositTokenAccountBump == nil {
			return errors.New("DepositTokenAccountBump parameter is not set")
		}
		if inst.QuestsBump == nil {
			return errors.New("QuestsBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.QuestQuesteeReceipt is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Quest is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Quests is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.DepositTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PixelballzMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PixelballzTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.QuestAcc is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Questor is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Questee is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *EndQuest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EndQuest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DepositTokenAccountBump", *inst.DepositTokenAccountBump))
						paramsBranch.Child(ag_format.Param("             QuestsBump", *inst.QuestsBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   questQuesteeReceipt", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 quest", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                quests", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           initializer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          depositToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        pixelballzMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       pixelballzToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              questAcc", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("               questor", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("               questee", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                oracle", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("            slotHashes", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj EndQuest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DepositTokenAccountBump` param:
	err = encoder.Encode(obj.DepositTokenAccountBump)
	if err != nil {
		return err
	}
	// Serialize `QuestsBump` param:
	err = encoder.Encode(obj.QuestsBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EndQuest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DepositTokenAccountBump`:
	err = decoder.Decode(&obj.DepositTokenAccountBump)
	if err != nil {
		return err
	}
	// Deserialize `QuestsBump`:
	err = decoder.Decode(&obj.QuestsBump)
	if err != nil {
		return err
	}
	return nil
}

// NewEndQuestInstruction declares a new EndQuest instruction with the provided parameters and accounts.
func NewEndQuestInstruction(
	// Parameters:
	depositTokenAccountBump uint8,
	questsBump uint8,
	// Accounts:
	questQuesteeReceipt ag_solanago.PublicKey,
	quest ag_solanago.PublicKey,
	quests ag_solanago.PublicKey,
	initializer ag_solanago.PublicKey,
	depositTokenAccount ag_solanago.PublicKey,
	pixelballzMint ag_solanago.PublicKey,
	pixelballzTokenAccount ag_solanago.PublicKey,
	questAcc ag_solanago.PublicKey,
	questor ag_solanago.PublicKey,
	questee ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *EndQuest {
	return NewEndQuestInstructionBuilder().
		SetDepositTokenAccountBump(depositTokenAccountBump).
		SetQuestsBump(questsBump).
		SetQuestQuesteeReceiptAccount(questQuesteeReceipt).
		SetQuestAccount(quest).
		SetQuestsAccount(quests).
		SetInitializerAccount(initializer).
		SetDepositTokenAccountAccount(depositTokenAccount).
		SetPixelballzMintAccount(pixelballzMint).
		SetPixelballzTokenAccountAccount(pixelballzTokenAccount).
		SetQuestAccAccount(questAcc).
		SetQuestorAccount(questor).
		SetQuesteeAccount(questee).
		SetOracleAccount(oracle).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSlotHashesAccount(slotHashes).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
