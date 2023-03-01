package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
