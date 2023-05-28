package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradingSessionStatus = "trading_session_status"

var _ sdk.Msg = &MsgTradingSessionStatus{}

func NewMsgTradingSessionStatus(creator string, tradSesReqID string, tradingSessionID string, tradSesStatus int32, tradSesStatusRejReason int32, tradSesStartTime string, tradSesOpenTime string, tradSesPreCloseTime string, tradSesCloseTime string, tradSesEndTime string, totalVolumeTraded int32, tradSesHighPx string, tradSesLowPx string, securityID string, securityIDSource string, symbol string, securityExchange string, marketSegmentID string, marketID string) *MsgTradingSessionStatus {
	return &MsgTradingSessionStatus{
		Creator:                creator,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
