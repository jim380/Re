package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderExecutionReport = "order_execution_report"

var _ sdk.Msg = &MsgOrderExecutionReport{}

func NewMsgOrderExecutionReport(creator string, sessionID string) *MsgOrderExecutionReport {
	return &MsgOrderExecutionReport{
		Creator:   creator,
		SessionID: sessionID,
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
