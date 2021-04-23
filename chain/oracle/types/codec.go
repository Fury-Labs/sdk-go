package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/oracle interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetPriceFeederPrice{}, "oracle/MsgSetPriceFeederPrice", nil)
	cdc.RegisterConcrete(&MsgRelay{}, "oracle/MsgRelay", nil)
	cdc.RegisterConcrete(&MsgRelayCoinbaseMessages{}, "oracle/MsgRelayCoinbaseMessages", nil)


	cdc.RegisterConcrete(&GrantBandOraclePrivilegeProposal{}, "oracle/GrantBandOraclePrivilegeProposal", nil)
	cdc.RegisterConcrete(&RevokeBandOraclePrivilegeProposal{}, "oracle/RevokeBandOraclePrivilegeProposal", nil)
	cdc.RegisterConcrete(&GrantPriceFeederPrivilegeProposal{}, "oracle/GrantPriceFeederPrivilegeProposal", nil)
	cdc.RegisterConcrete(&RevokePriceFeederPrivilegeProposal{}, "oracle/RevokePriceFeederPrivilegeProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetPriceFeederPrice{},
		&MsgRelay{},
		&MsgRelayCoinbaseMessages{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&GrantBandOraclePrivilegeProposal{},
		&RevokeBandOraclePrivilegeProposal{},
		&GrantPriceFeederPrivilegeProposal{},
		&RevokePriceFeederPrivilegeProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/oracle module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/oracle and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
