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
	cdc.RegisterConcrete(&MsgOrderExecutionReport{}, "fix/OrderExecutionReport", nil)
	cdc.RegisterConcrete(&MsgQuoteRequest{}, "fix/QuoteRequest", nil)
	cdc.RegisterConcrete(&MsgQuoteAcknowledgement{}, "fix/QuoteAcknowledgement", nil)
	cdc.RegisterConcrete(&MsgQuoteRequestReject{}, "fix/QuoteRequestReject", nil)
	cdc.RegisterConcrete(&MsgSecurityDefinitionRequest{}, "fix/SecurityDefinitionRequest", nil)
	cdc.RegisterConcrete(&MsgSecurityDefinition{}, "fix/SecurityDefinition", nil)
	cdc.RegisterConcrete(&MsgSecurityDefinitionRequestReject{}, "fix/SecurityDefinitionRequestReject", nil)
	cdc.RegisterConcrete(&MsgOrderMassStatusRequest{}, "fix/OrderMassStatusRequest", nil)
	cdc.RegisterConcrete(&MsgOrderMassStatusReport{}, "fix/OrderMassStatusReport", nil)
	cdc.RegisterConcrete(&MsgOrderMassStatusRequestReject{}, "fix/OrderMassStatusRequestReject", nil)
	cdc.RegisterConcrete(&MsgTradingSessionStatusRequest{}, "fix/TradingSessionStatusRequest", nil)
	cdc.RegisterConcrete(&MsgTradingSessionStatus{}, "fix/TradingSessionStatus", nil)
	cdc.RegisterConcrete(&MsgTradingSessionStatusRequestReject{}, "fix/TradingSessionStatusRequestReject", nil)
	cdc.RegisterConcrete(&MsgTradingSessionListRequest{}, "fix/TradingSessionListRequest", nil)
	cdc.RegisterConcrete(&MsgTradingSessionListResponse{}, "fix/TradingSessionListResponse", nil)
	cdc.RegisterConcrete(&MsgTradingSessionListRequestReject{}, "fix/TradingSessionListRequestReject", nil)
	cdc.RegisterConcrete(&MsgSecurityListRequest{}, "fix/SecurityListRequest", nil)
	cdc.RegisterConcrete(&MsgSecurityListResponse{}, "fix/SecurityListResponse", nil)
	cdc.RegisterConcrete(&MsgSecurityListRequestReject{}, "fix/SecurityListRequestReject", nil)
	cdc.RegisterConcrete(&MsgSecurityStatusRequest{}, "fix/SecurityStatusRequest", nil)
	cdc.RegisterConcrete(&MsgSecurityStatusResponse{}, "fix/SecurityStatusResponse", nil)
	cdc.RegisterConcrete(&MsgSecurityStatusRequestReject{}, "fix/SecurityStatusRequestReject", nil)
	cdc.RegisterConcrete(&MsgSecurityTypesRequest{}, "fix/SecurityTypesRequest", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderExecutionReport{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgQuoteRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgQuoteAcknowledgement{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgQuoteRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityDefinitionRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityDefinition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityDefinitionRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderMassStatusRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderMassStatusReport{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOrderMassStatusRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionStatusRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionStatus{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionStatusRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionListRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionListResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTradingSessionListRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityListRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityListResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityListRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityStatusRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityStatusResponse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityStatusRequestReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSecurityTypesRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
