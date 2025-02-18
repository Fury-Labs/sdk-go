package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	ExtensionOptionsEthereumTxI interface{}
)

// RegisterInterfaces registers the client interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgEthereumTx{},
	)

	registry.RegisterInterface("kaiju.evm.v1beta1.ExtensionOptionsEthereumTx", (*ExtensionOptionsEthereumTxI)(nil))
	registry.RegisterImplementations(
		(*ExtensionOptionsEthereumTxI)(nil),
		&ExtensionOptionsEthereumTx{},
	)
}

var (
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)
