package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTradeCaptureReport = "trade_capture_report"

var _ sdk.Msg = &MsgTradeCaptureReport{}

func NewMsgTradeCaptureReport(creator string, sessionID string, tradeReportID string, tradeReportTransType string, tradeReportType string, trdType string, trdSubType string, side string, orderQty string, lastQty string, lastPx string, grossTradeAmt string, execID string, orderID string, tradeID string, origTradeID string, symbol string, securityID string, securityIDSource string, tradeDate string, transactTime string, settlType string, settlDate string) *MsgTradeCaptureReport {
	return &MsgTradeCaptureReport{
		Creator:              creator,
		SessionID:            sessionID,
		TradeReportID:        tradeReportID,
		TradeReportTransType: tradeReportTransType,
		TradeReportType:      tradeReportType,
		TrdType:              trdType,
		TrdSubType:           trdSubType,
		Side:                 side,
		OrderQty:             orderQty,
		LastQty:              lastQty,
		LastPx:               lastPx,
		GrossTradeAmt:        grossTradeAmt,
		ExecID:               execID,
		OrderID:              orderID,
		TradeID:              tradeID,
		OrigTradeID:          origTradeID,
		Symbol:               symbol,
		SecurityID:           securityID,
		SecurityIDSource:     securityIDSource,
		TradeDate:            tradeDate,
		TransactTime:         transactTime,
		SettlType:            settlType,
		SettlDate:            settlDate,
	}
}

func (msg *MsgTradeCaptureReport) Route() string {
	return RouterKey
}

func (msg *MsgTradeCaptureReport) Type() string {
	return TypeMsgTradeCaptureReport
}

func (msg *MsgTradeCaptureReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradeCaptureReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradeCaptureReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradeCaptureReportAcknowledgement = "trade_capture_report_acknowledgement"

var _ sdk.Msg = &MsgTradeCaptureReportAcknowledgement{}

func NewMsgTradeCaptureReportAcknowledgement(creator string, sessionID string, tradeReportID string, tradeID string, secondaryTradeID string, tradeReportType string, trdType string, trdSubType string, execType string, tradeReportRefID string, secondaryTradeReportID string, tradeReportStatus string, tradeTransType string, tradeReportRejectReason int64, text string) *MsgTradeCaptureReportAcknowledgement {
	return &MsgTradeCaptureReportAcknowledgement{
		Creator:                 creator,
		SessionID:               sessionID,
		TradeReportID:           tradeReportID,
		TradeID:                 tradeID,
		SecondaryTradeID:        secondaryTradeID,
		TradeReportType:         tradeReportType,
		TrdType:                 trdType,
		TrdSubType:              trdSubType,
		ExecType:                execType,
		TradeReportRefID:        tradeReportRefID,
		SecondaryTradeReportID:  secondaryTradeReportID,
		TradeReportStatus:       tradeReportStatus,
		TradeTransType:          tradeTransType,
		TradeReportRejectReason: tradeReportRejectReason,
		Text:                    text,
	}
}

func (msg *MsgTradeCaptureReportAcknowledgement) Route() string {
	return RouterKey
}

func (msg *MsgTradeCaptureReportAcknowledgement) Type() string {
	return TypeMsgTradeCaptureReportAcknowledgement
}

func (msg *MsgTradeCaptureReportAcknowledgement) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradeCaptureReportAcknowledgement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradeCaptureReportAcknowledgement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTradeCaptureReportRejection = "trade_capture_report_rejection"

var _ sdk.Msg = &MsgTradeCaptureReportRejection{}

func NewMsgTradeCaptureReportRejection(creator string, sessionID string, tradeReportID string, tradeReportRejectReason int64, tradeReportRejectRefID string, text string) *MsgTradeCaptureReportRejection {
	return &MsgTradeCaptureReportRejection{
		Creator:                 creator,
		SessionID:               sessionID,
		TradeReportID:           tradeReportID,
		TradeReportRejectReason: tradeReportRejectReason,
		TradeReportRejectRefID:  tradeReportRejectRefID,
		Text:                    text,
	}
}

func (msg *MsgTradeCaptureReportRejection) Route() string {
	return RouterKey
}

func (msg *MsgTradeCaptureReportRejection) Type() string {
	return TypeMsgTradeCaptureReportRejection
}

func (msg *MsgTradeCaptureReportRejection) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTradeCaptureReportRejection) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTradeCaptureReportRejection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
