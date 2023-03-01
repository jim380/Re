package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderExecutionReport = "order_execution_report"

var _ sdk.Msg = &MsgOrderExecutionReport{}

func NewMsgOrderExecutionReport(creator string, sessionID string, clOrdID string, orderID string, execID string, ordStatus string, execType string, symbol string, side int64, orderQty int64, price int64, timeInForce int64, lastPx int64, lastQty int64, leavesQty int64, cumQty int64, avgPx int64, text string) *MsgOrderExecutionReport {
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
