package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCosmoshubTxs = "cosmoshub_txs"

var _ sdk.Msg = &MsgCosmoshubTxs{}

func NewMsgCosmoshubTxs(creator string, oracleID string, address string) *MsgCosmoshubTxs {
	return &MsgCosmoshubTxs{
		Creator:  creator,
		OracleID: oracleID,
		Address:  address,
	}
}

func (msg *MsgCosmoshubTxs) Route() string {
	return RouterKey
}

func (msg *MsgCosmoshubTxs) Type() string {
	return TypeMsgCosmoshubTxs
}

func (msg *MsgCosmoshubTxs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCosmoshubTxs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCosmoshubTxs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
