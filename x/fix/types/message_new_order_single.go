package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
