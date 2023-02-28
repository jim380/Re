package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewOrderSingle = "new_order_single"

var _ sdk.Msg = &MsgNewOrderSingle{}

func NewMsgNewOrderSingle(creator string, sessionName string) *MsgNewOrderSingle {
	return &MsgNewOrderSingle{
		Creator:     creator,
		SessionName: sessionName,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}