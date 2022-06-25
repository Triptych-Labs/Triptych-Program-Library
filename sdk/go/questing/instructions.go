// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package questing

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Questing"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_EnrollQuestor = ag_binary.TypeID([8]byte{24, 214, 177, 154, 205, 7, 26, 46})

	Instruction_EnrollQuestee = ag_binary.TypeID([8]byte{142, 147, 49, 24, 11, 200, 232, 83})

	Instruction_UpdateQuestee = ag_binary.TypeID([8]byte{56, 254, 41, 193, 39, 42, 198, 0})

	Instruction_EnableQuests = ag_binary.TypeID([8]byte{196, 145, 208, 170, 4, 80, 81, 162})

	Instruction_CreateQuest = ag_binary.TypeID([8]byte{112, 49, 32, 224, 255, 173, 5, 7})

	Instruction_RegisterQuestReward = ag_binary.TypeID([8]byte{36, 83, 38, 130, 109, 90, 36, 132})

	Instruction_StartQuest = ag_binary.TypeID([8]byte{212, 181, 79, 44, 200, 242, 13, 105})

	Instruction_EndQuest = ag_binary.TypeID([8]byte{250, 149, 247, 239, 238, 172, 90, 150})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_EnrollQuestor:
		return "EnrollQuestor"
	case Instruction_EnrollQuestee:
		return "EnrollQuestee"
	case Instruction_UpdateQuestee:
		return "UpdateQuestee"
	case Instruction_EnableQuests:
		return "EnableQuests"
	case Instruction_CreateQuest:
		return "CreateQuest"
	case Instruction_RegisterQuestReward:
		return "RegisterQuestReward"
	case Instruction_StartQuest:
		return "StartQuest"
	case Instruction_EndQuest:
		return "EndQuest"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"enroll_questor", (*EnrollQuestor)(nil),
		},
		{
			"enroll_questee", (*EnrollQuestee)(nil),
		},
		{
			"update_questee", (*UpdateQuestee)(nil),
		},
		{
			"enable_quests", (*EnableQuests)(nil),
		},
		{
			"create_quest", (*CreateQuest)(nil),
		},
		{
			"register_quest_reward", (*RegisterQuestReward)(nil),
		},
		{
			"start_quest", (*StartQuest)(nil),
		},
		{
			"end_quest", (*EndQuest)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
