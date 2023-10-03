package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEthereumTxs = "ethereum_txs"

var _ sdk.Msg = &MsgEthereumTxs{}

func NewMsgEthereumTxs(creator string, oracleID string, address string) *MsgEthereumTxs {
	return &MsgEthereumTxs{
		Creator:  creator,
		OracleID: oracleID,
		Address:  address,
	}
}

func (msg *MsgEthereumTxs) Route() string {
	return RouterKey
}

func (msg *MsgEthereumTxs) Type() string {
	return TypeMsgEthereumTxs
}

func (msg *MsgEthereumTxs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEthereumTxs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEthereumTxs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
