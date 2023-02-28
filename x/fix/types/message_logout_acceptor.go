package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogoutAcceptor = "logout_acceptor"

var _ sdk.Msg = &MsgLogoutAcceptor{}

func NewMsgLogoutAcceptor(acceptorAddress string, sessionName string, text SessionLogoutAcceptor) *MsgLogoutAcceptor {
	return &MsgLogoutAcceptor{
		AcceptorAddress:       acceptorAddress,
		SessionName:           sessionName,
		SessionLogoutAcceptor: &text,
	}
}

func (msg *MsgLogoutAcceptor) Route() string {
	return RouterKey
}

func (msg *MsgLogoutAcceptor) Type() string {
	return TypeMsgLogoutAcceptor
}

func (msg *MsgLogoutAcceptor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogoutAcceptor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogoutAcceptor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
