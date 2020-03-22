package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	channelexported "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
)

var _ channelexported.PacketDataI = OraclePacketData{}

type OraclePacketData struct {
	Data []byte `json:"data" yaml:"data"`
}

// NewFungibleTokenPacketData contructs a new FungibleTokenPacketData instance
func NewOraclePacketData(data []byte) OraclePacketData {
	return OraclePacketData{
		Data: data,
	}
}

// String returns a string representation of FungibleTokenPacketData
func (o OraclePacketData) String() string {
	return fmt.Sprintf(`OraclePacketData:
	Data:                 %s`,
		o.Data,
	)
}

// ValidateBasic implements channelexported.PacketDataI
func (o OraclePacketData) ValidateBasic() error {
	return nil
}

// GetBytes implements channelexported.PacketDataI
func (o OraclePacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(o))
}

// GetTimeoutHeight implements channelexported.PacketDataI
func (o OraclePacketData) GetTimeoutHeight() uint64 {
	return uint64(18446744073709551615)
}

// Type implements channelexported.PacketDataI
func (ftpd OraclePacketData) Type() string {
	return "zoracle"
}
