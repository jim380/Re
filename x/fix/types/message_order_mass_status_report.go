package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderMassStatusReport = "order_mass_status_report"

var _ sdk.Msg = &MsgOrderMassStatusReport{}

func NewMsgOrderMassStatusReport(creator string, sessionID string, clOrdID string, massStatusReqID string, account string, symbol string, securityID string, tradingSessionID string, ordStatus string, execType string, ordQty string, lastPx string, cumQty string, avgPx string, leavesQty string) *MsgOrderMassStatusReport {
	return &MsgOrderMassStatusReport{
		Creator:          creator,
		SessionID:        sessionID,
		ClOrdID:          clOrdID,
		MassStatusReqID:  massStatusReqID,
		Account:          account,
		Symbol:           symbol,
		SecurityID:       securityID,
		TradingSessionID: tradingSessionID,
		OrdStatus:        ordStatus,
		ExecType:         execType,
		OrdQty:           ordQty,
		LastPx:           lastPx,
		CumQty:           cumQty,
		AvgPx:            avgPx,
		LeavesQty:        leavesQty,
	}
}

func (msg *MsgOrderMassStatusReport) Route() string {
	return RouterKey
}

func (msg *MsgOrderMassStatusReport) Type() string {
	return TypeMsgOrderMassStatusReport
}

func (msg *MsgOrderMassStatusReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOrderMassStatusReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOrderMassStatusReport) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
