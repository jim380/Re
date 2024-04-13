package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
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
	cdc.RegisterConcrete(&MsgSecurityTypesResponse{}, "fix/SecurityTypesResponse", nil)
	cdc.RegisterConcrete(&MsgSecurityTypesRequestReject{}, "fix/SecurityTypesRequestReject", nil)
	cdc.RegisterConcrete(&MsgRegisterAccount{}, "fix/RegisterAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateAccount{}, "fix/UpdateAccount", nil)
	cdc.RegisterConcrete(&MsgDeleteAccount{}, "fix/DeleteAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLogonInitiator{},
		&MsgLogonAcceptor{},
		&MsgTerminateLogon{},
		&MsgLogonReject{},
		&MsgLogoutInitiator{},
		&MsgLogoutAcceptor{},
		&MsgNewOrderSingle{},
		&MsgOrderCancelRequest{},
		&MsgOrderCancelReject{},
		&MsgOrderExecutionReport{},
		&MsgQuoteRequest{},
		&MsgQuoteAcknowledgement{},
		&MsgQuoteRequestReject{},
		&MsgSecurityDefinitionRequest{},
		&MsgSecurityDefinition{},
		&MsgSecurityDefinitionRequestReject{},
		&MsgOrderMassStatusRequest{},
		&MsgOrderMassStatusReport{},
		&MsgOrderMassStatusRequestReject{},
		&MsgTradingSessionStatusRequest{},
		&MsgTradingSessionStatus{},
		&MsgTradingSessionStatusRequestReject{},
		&MsgTradingSessionListRequest{},
		&MsgTradingSessionListResponse{},
		&MsgTradingSessionListRequestReject{},
		&MsgSecurityListRequest{},
		&MsgSecurityListResponse{},
		&MsgSecurityListRequestReject{},
		&MsgSecurityStatusRequest{},
		&MsgSecurityStatusResponse{},
		&MsgSecurityStatusRequestReject{},
		&MsgSecurityTypesRequest{},
		&MsgSecurityTypesResponse{},
		&MsgSecurityTypesRequestReject{},
		&MsgRegisterAccount{},
		&MsgUpdateAccount{},
		&MsgDeleteAccount{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
