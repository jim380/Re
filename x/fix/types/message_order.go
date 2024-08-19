package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewOrderSingle = "new_order_single"

var _ sdk.Msg = &MsgNewOrderSingle{}

func NewMsgNewOrderSingle(creator string, sessionID string, clOrdID string, symbol string, side int64, orderQty string, ordType int64, price string, timeInForce int64, text string) *MsgNewOrderSingle {
	return &MsgNewOrderSingle{
		Creator:     creator,
		SessionID:   sessionID,
		ClOrdID:     clOrdID,
		Symbol:      symbol,
		Side:        side,
		OrderQty:    orderQty,
		OrdType:     ordType,
		Price:       price,
		TimeInForce: timeInForce,
		Text:        text,
	}
}

func (msg *MsgNewOrderSingle) Route() string {
	return RouterKey
}

func (msg *MsgNewOrderSingle) Type() string {
	return TypeMsgNewOrderSingle
}

func (msg *MsgNewOrderSingle) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewOrderSingle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewOrderSingle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgOrderCancelRequest = "order_cancel_request"

var _ sdk.Msg = &MsgOrderCancelRequest{}

func NewMsgOrderCancelRequest(creator string, sessionID string, origClOrdID string, clOrdID string) *MsgOrderCancelRequest {
	return &MsgOrderCancelRequest{
		Creator:     creator,
		SessionID:   sessionID,
		OrigClOrdID: origClOrdID,
		ClOrdID:     clOrdID,
	}
}

func (msg *MsgOrderCancelRequest) Route() string {
	return RouterKey
}

func (msg *MsgOrderCancelRequest) Type() string {
	return TypeMsgOrderCancelRequest
}

func (msg *MsgOrderCancelRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderCancelRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderCancelRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgOrderExecutionReport = "order_execution_report"

var _ sdk.Msg = &MsgOrderExecutionReport{}

func NewMsgOrderExecutionReport(creator string, sessionID string, clOrdID string, orderID string, execID string, ordStatus string, execType string, symbol string, side int64, orderQty string, price string, timeInForce int64, lastPx int64, lastQty int64, leavesQty int64, cumQty int64, avgPx int64, text string) *MsgOrderExecutionReport {
	return &MsgOrderExecutionReport{
		Creator:     creator,
		SessionID:   sessionID,
		ClOrdID:     clOrdID,
		OrderID:     orderID,
		ExecID:      execID,
		OrdStatus:   ordStatus,
		ExecType:    execType,
		Symbol:      symbol,
		Side:        side,
		OrderQty:    orderQty,
		Price:       price,
		TimeInForce: timeInForce,
		LastPx:      lastPx,
		LastQty:     lastQty,
		LeavesQty:   leavesQty,
		CumQty:      cumQty,
		AvgPx:       avgPx,
		Text:        text,
	}
}

func (msg *MsgOrderExecutionReport) Route() string {
	return RouterKey
}

func (msg *MsgOrderExecutionReport) Type() string {
	return TypeMsgOrderExecutionReport
}

func (msg *MsgOrderExecutionReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderExecutionReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderExecutionReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgOrderCancelReject = "order_cancel_reject"

var _ sdk.Msg = &MsgOrderCancelReject{}

func NewMsgOrderCancelReject(creator string, sessionID string, orderID string, origClOrdID string, clOrdID string, cxlRejReason int64, cxlRejResponseTo int64) *MsgOrderCancelReject {
	return &MsgOrderCancelReject{
		Creator:          creator,
		SessionID:        sessionID,
		OrderID:          orderID,
		OrigClOrdID:      origClOrdID,
		ClOrdID:          clOrdID,
		CxlRejReason:     cxlRejReason,
		CxlRejResponseTo: cxlRejResponseTo,
	}
}

func (msg *MsgOrderCancelReject) Route() string {
	return RouterKey
}

func (msg *MsgOrderCancelReject) Type() string {
	return TypeMsgOrderCancelReject
}

func (msg *MsgOrderCancelReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderCancelReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderCancelReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
