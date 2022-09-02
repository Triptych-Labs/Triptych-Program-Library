// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package blackjack

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type House struct {
	Oracle    ag_solanago.PublicKey
	PayedOut  uint64
	Collected uint64
}

var HouseDiscriminator = [8]byte{21, 145, 94, 109, 254, 199, 210, 151}

func (obj House) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(HouseDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `PayedOut` param:
	err = encoder.Encode(obj.PayedOut)
	if err != nil {
		return err
	}
	// Serialize `Collected` param:
	err = encoder.Encode(obj.Collected)
	if err != nil {
		return err
	}
	return nil
}

func (obj *House) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(HouseDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[21 145 94 109 254 199 210 151]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `PayedOut`:
	err = decoder.Decode(&obj.PayedOut)
	if err != nil {
		return err
	}
	// Deserialize `Collected`:
	err = decoder.Decode(&obj.Collected)
	if err != nil {
		return err
	}
	return nil
}

type Games struct {
	Initializer ag_solanago.PublicKey
	Games       uint64
}

var GamesDiscriminator = [8]byte{73, 124, 61, 201, 178, 83, 6, 66}

func (obj Games) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(GamesDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initializer` param:
	err = encoder.Encode(obj.Initializer)
	if err != nil {
		return err
	}
	// Serialize `Games` param:
	err = encoder.Encode(obj.Games)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Games) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(GamesDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[73 124 61 201 178 83 6 66]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initializer`:
	err = decoder.Decode(&obj.Initializer)
	if err != nil {
		return err
	}
	// Deserialize `Games`:
	err = decoder.Decode(&obj.Games)
	if err != nil {
		return err
	}
	return nil
}

type Game struct {
	Index        uint64
	Initialized  *bool `bin:"optional"`
	Player       ag_solanago.PublicKey
	BetAmount    uint64
	DailyEpoch   uint64
	Hands        [2][5]uint8
	PlayerBusted bool
	DealerBusted bool
	Terminated   bool
}

var GameDiscriminator = [8]byte{27, 90, 166, 125, 74, 100, 121, 18}

func (obj Game) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(GameDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Initialized` param (optional):
	{
		if obj.Initialized == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Initialized)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Player` param:
	err = encoder.Encode(obj.Player)
	if err != nil {
		return err
	}
	// Serialize `BetAmount` param:
	err = encoder.Encode(obj.BetAmount)
	if err != nil {
		return err
	}
	// Serialize `DailyEpoch` param:
	err = encoder.Encode(obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Serialize `Hands` param:
	err = encoder.Encode(obj.Hands)
	if err != nil {
		return err
	}
	// Serialize `PlayerBusted` param:
	err = encoder.Encode(obj.PlayerBusted)
	if err != nil {
		return err
	}
	// Serialize `DealerBusted` param:
	err = encoder.Encode(obj.DealerBusted)
	if err != nil {
		return err
	}
	// Serialize `Terminated` param:
	err = encoder.Encode(obj.Terminated)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Game) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(GameDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[27 90 166 125 74 100 121 18]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Initialized` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Initialized)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Player`:
	err = decoder.Decode(&obj.Player)
	if err != nil {
		return err
	}
	// Deserialize `BetAmount`:
	err = decoder.Decode(&obj.BetAmount)
	if err != nil {
		return err
	}
	// Deserialize `DailyEpoch`:
	err = decoder.Decode(&obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Deserialize `Hands`:
	err = decoder.Decode(&obj.Hands)
	if err != nil {
		return err
	}
	// Deserialize `PlayerBusted`:
	err = decoder.Decode(&obj.PlayerBusted)
	if err != nil {
		return err
	}
	// Deserialize `DealerBusted`:
	err = decoder.Decode(&obj.DealerBusted)
	if err != nil {
		return err
	}
	// Deserialize `Terminated`:
	err = decoder.Decode(&obj.Terminated)
	if err != nil {
		return err
	}
	return nil
}

type Stats struct {
	Initialized *bool `bin:"optional"`
	Oracle      ag_solanago.PublicKey
	DailyEpoch  uint64
	Games       [4]uint64
}

var StatsDiscriminator = [8]byte{190, 125, 51, 63, 169, 197, 36, 238}

func (obj Stats) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(StatsDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized` param (optional):
	{
		if obj.Initialized == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Initialized)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `DailyEpoch` param:
	err = encoder.Encode(obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Serialize `Games` param:
	err = encoder.Encode(obj.Games)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Stats) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(StatsDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[190 125 51 63 169 197 36 238]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Initialized)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `DailyEpoch`:
	err = decoder.Decode(&obj.DailyEpoch)
	if err != nil {
		return err
	}
	// Deserialize `Games`:
	err = decoder.Decode(&obj.Games)
	if err != nil {
		return err
	}
	return nil
}