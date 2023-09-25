package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBitcoinTxs = "bitcoin_txs"

var _ sdk.Msg = &MsgBitcoinTxs{}

func NewMsgBitcoinTxs(creator string, oracleID string, address string) *MsgBitcoinTxs {
	return &MsgBitcoinTxs{
		Creator:  creator,
		OracleID: oracleID,
		Address:  address,
	}
}

func (msg *MsgBitcoinTxs) Route() string {
	return RouterKey
}

func (msg *MsgBitcoinTxs) Type() string {
	return TypeMsgBitcoinTxs
}

func (msg *MsgBitcoinTxs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBitcoinTxs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBitcoinTxs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
