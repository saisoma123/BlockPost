package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// Registers the msg and msg_server types with the app itself
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBlockPostMessage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBlockPostMessageResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParamsResponse{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
