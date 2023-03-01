package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOrderCancelReject = "order_cancel_reject"

var _ sdk.Msg = &MsgOrderCancelReject{}

func NewMsgOrderCancelReject(creator string, sessionID string,  orderID string, origClOrdID string, clOrdID string, cxlRejReason int64, cxlRejResponseTo int64) *MsgOrderCancelReject {
	return &MsgOrderCancelReject{
		Creator:   creator,
		SessionID: sessionID,
		OrderID: orderID,
		OrigClOrdID: origClOrdID,
		ClOrdID: clOrdID,
		CxlRejReason: cxlRejReason,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
