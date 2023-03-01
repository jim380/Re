package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAccount{}, "fix/CreateAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateAccount{}, "fix/UpdateAccount", nil)
	cdc.RegisterConcrete(&MsgDeleteAccount{}, "fix/DeleteAccount", nil)
	cdc.RegisterConcrete(&MsgLogonInitiator{}, "fix/LogonInitiator", nil)
	cdc.RegisterConcrete(&MsgLogonAcceptor{}, "fix/LogonAcceptor", nil)
	cdc.RegisterConcrete(&MsgTerminateLogon{}, "fix/TerminateLogon", nil)
	cdc.RegisterConcrete(&MsgLogonReject{}, "fix/LogonReject", nil)
	cdc.RegisterConcrete(&MsgLogoutInitiator{}, "fix/LogoutInitiator", nil)
	cdc.RegisterConcrete(&MsgLogoutAcceptor{}, "fix/LogoutAcceptor", nil)
	cdc.RegisterConcrete(&MsgNewOrderSingle{}, "fix/NewOrderSingle", nil)
	cdc.RegisterConcrete(&MsgOrderCancelRequest{}, "fix/OrderCancelRequest", nil)
	cdc.RegisterConcrete(&MsgOrderCancelReject{}, "fix/OrderCancelReject", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAccount{},
		&MsgUpdateAccount{},
		&MsgDeleteAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogonInitiator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogonAcceptor{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTerminateLogon{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogonReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogoutInitiator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogoutAcceptor{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNewOrderSingle{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderCancelRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderCancelReject{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
