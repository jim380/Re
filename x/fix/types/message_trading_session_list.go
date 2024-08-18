package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradingSessionListRequest = "trading_session_list_request"

var _ sdk.Msg = &MsgTradingSessionListRequest{}

func NewMsgTradingSessionListRequest(creator string, sessionID string, tradSesReqID string, tradingSessionID string, tradingSessionSubID string, securityExchange string, tradSesMethod string, tradSesMode string, subscriptionRequestType string) *MsgTradingSessionListRequest {
	return &MsgTradingSessionListRequest{
		Creator:                 creator,
		SessionID:               sessionID,
		TradSesReqID:            tradSesReqID,
		TradingSessionID:        tradingSessionID,
		TradingSessionSubID:     tradingSessionSubID,
		SecurityExchange:        securityExchange,
		TradSesMethod:           tradSesMethod,
		TradSesMode:             tradSesMode,
		SubscriptionRequestType: subscriptionRequestType,
	}
}

func (msg *MsgTradingSessionListRequest) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionListRequest) Type() string {
	return TypeMsgTradingSessionListRequest
}

func (msg *MsgTradingSessionListRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionListRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionListRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradingSessionListResponse = "trading_session_list_response"

var _ sdk.Msg = &MsgTradingSessionListResponse{}

func NewMsgTradingSessionListResponse(creator string, sessionID string, tradSesReqID string, noTradingSessions string, tradingSessionID string, tradingSessionSubID string, securityExchange string, tradSesMethod string, tradSesMode string, unsolicitedIndicator string, tradSesStatus string, tradSesStatusRejReason string, tradSesStartTime string, tradSesOpenTime string, tradSesPreCloseTime string, tradSesCloseTime string, tradSesEndTime string) *MsgTradingSessionListResponse {
	return &MsgTradingSessionListResponse{
		Creator:                creator,
		SessionID:              sessionID,
		TradSesReqID:           tradSesReqID,
		NoTradingSessions:      noTradingSessions,
		TradingSessionID:       tradingSessionID,
		TradingSessionSubID:    tradingSessionSubID,
		SecurityExchange:       securityExchange,
		TradSesMethod:          tradSesMethod,
		TradSesMode:            tradSesMode,
		UnsolicitedIndicator:   unsolicitedIndicator,
		TradSesStatus:          tradSesStatus,
		TradSesStatusRejReason: tradSesStatusRejReason,
		TradSesStartTime:       tradSesStartTime,
		TradSesOpenTime:        tradSesOpenTime,
		TradSesPreCloseTime:    tradSesPreCloseTime,
		TradSesCloseTime:       tradSesCloseTime,
		TradSesEndTime:         tradSesEndTime,
	}
}

func (msg *MsgTradingSessionListResponse) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionListResponse) Type() string {
	return TypeMsgTradingSessionListResponse
}

func (msg *MsgTradingSessionListResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionListResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionListResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradingSessionListRequestReject = "trading_session_list_request_reject"

var _ sdk.Msg = &MsgTradingSessionListRequestReject{}

func NewMsgTradingSessionListRequestReject(creator string, sessionID string, tradSesReqID string, tradSesStatus string, tradSesStatusRejReason string, text string) *MsgTradingSessionListRequestReject {
	return &MsgTradingSessionListRequestReject{
		Creator:                creator,
		SessionID:              sessionID,
		TradSesReqID:           tradSesReqID,
		TradSesStatus:          tradSesStatus,
		TradSesStatusRejReason: tradSesStatusRejReason,
		Text:                   text,
	}
}

func (msg *MsgTradingSessionListRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionListRequestReject) Type() string {
	return TypeMsgTradingSessionListRequestReject
}

func (msg *MsgTradingSessionListRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionListRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionListRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
