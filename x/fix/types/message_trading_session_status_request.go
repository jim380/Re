package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradingSessionStatusRequest = "trading_session_status_request"

var _ sdk.Msg = &MsgTradingSessionStatusRequest{}

func NewMsgTradingSessionStatusRequest(creator string, tradSesReqID string, tradingSessionID string, tradingSessionSubID string, marketID string, subscriptionRequest string, securityID string, securityIDSource string, symbol string, securityExchange string, marketSegmentID string, tradSesReqType int32, tradSesSubReqType int32, tradSesMode int32, tradingSessionDate string, tradingSessionTime string, tradingSessionSubTime string, expirationDate string) *MsgTradingSessionStatusRequest {
	return &MsgTradingSessionStatusRequest{
		Creator:               creator,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
