package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradingSessionStatusRequest = "trading_session_status_request"

var _ sdk.Msg = &MsgTradingSessionStatusRequest{}

func NewMsgTradingSessionStatusRequest(creator string, sessionID string, tradSesReqID string, tradingSessionID string, tradingSessionSubID string, marketID string, subscriptionRequest string, securityID string, securityIDSource string, symbol string, securityExchange string, marketSegmentID string, tradSesReqType int32, tradSesSubReqType int32, tradSesMode int32, tradingSessionDate string, tradingSessionTime string, tradingSessionSubTime string, expirationDate string) *MsgTradingSessionStatusRequest {
	return &MsgTradingSessionStatusRequest{
		Creator:               creator,
		SessionID:             sessionID,
		TradSesReqID:          tradSesReqID,
		TradingSessionID:      tradingSessionID,
		TradingSessionSubID:   tradingSessionSubID,
		MarketID:              marketID,
		SubscriptionRequest:   subscriptionRequest,
		SecurityID:            securityID,
		SecurityIDSource:      securityIDSource,
		Symbol:                symbol,
		SecurityExchange:      securityExchange,
		MarketSegmentID:       marketSegmentID,
		TradSesReqType:        tradSesReqType,
		TradSesSubReqType:     tradSesSubReqType,
		TradSesMode:           tradSesMode,
		TradingSessionDate:    tradingSessionDate,
		TradingSessionTime:    tradingSessionTime,
		TradingSessionSubTime: tradingSessionSubTime,
		ExpirationDate:        expirationDate,
	}
}

func (msg *MsgTradingSessionStatusRequest) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionStatusRequest) Type() string {
	return TypeMsgTradingSessionStatusRequest
}

func (msg *MsgTradingSessionStatusRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionStatusRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionStatusRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradingSessionStatus = "trading_session_status"

var _ sdk.Msg = &MsgTradingSessionStatus{}

func NewMsgTradingSessionStatus(creator string, sessionID string, tradSesReqID string, tradingSessionID string, tradSesStatus int32, tradSesStatusRejReason int32, tradSesStartTime string, tradSesOpenTime string, tradSesPreCloseTime string, tradSesCloseTime string, tradSesEndTime string, totalVolumeTraded int32, tradSesHighPx string, tradSesLowPx string, securityID string, securityIDSource string, symbol string, securityExchange string, marketSegmentID string, marketID string) *MsgTradingSessionStatus {
	return &MsgTradingSessionStatus{
		Creator:                creator,
		SessionID:              sessionID,
		TradSesReqID:           tradSesReqID,
		TradingSessionID:       tradingSessionID,
		TradSesStatus:          tradSesStatus,
		TradSesStatusRejReason: tradSesStatusRejReason,
		TradSesStartTime:       tradSesStartTime,
		TradSesOpenTime:        tradSesOpenTime,
		TradSesPreCloseTime:    tradSesPreCloseTime,
		TradSesCloseTime:       tradSesCloseTime,
		TradSesEndTime:         tradSesEndTime,
		TotalVolumeTraded:      totalVolumeTraded,
		TradSesHighPx:          tradSesHighPx,
		TradSesLowPx:           tradSesLowPx,
		SecurityID:             securityID,
		SecurityIDSource:       securityIDSource,
		Symbol:                 symbol,
		SecurityExchange:       securityExchange,
		MarketSegmentID:        marketSegmentID,
		MarketID:               marketID,
	}
}

func (msg *MsgTradingSessionStatus) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionStatus) Type() string {
	return TypeMsgTradingSessionStatus
}

func (msg *MsgTradingSessionStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradingSessionStatusRequestReject = "trading_session_status_request_reject"

var _ sdk.Msg = &MsgTradingSessionStatusRequestReject{}

func NewMsgTradingSessionStatusRequestReject(creator string, sessionID string, refSeqNum string, refMsgType string, sessionRejectReason int32, text string) *MsgTradingSessionStatusRequestReject {
	return &MsgTradingSessionStatusRequestReject{
		Creator:             creator,
		SessionID:           sessionID,
		RefSeqNum:           refSeqNum,
		RefMsgType:          refMsgType,
		SessionRejectReason: sessionRejectReason,
		Text:                text,
	}
}

func (msg *MsgTradingSessionStatusRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgTradingSessionStatusRequestReject) Type() string {
	return TypeMsgTradingSessionStatusRequestReject
}

func (msg *MsgTradingSessionStatusRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradingSessionStatusRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradingSessionStatusRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
