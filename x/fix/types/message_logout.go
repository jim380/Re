package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogoutInitiator = "logout_initiator"

var _ sdk.Msg = &MsgLogoutInitiator{}

func NewMsgLogoutInitiator(initiatorAddress string, sessionID string, text SessionLogoutInitiator) *MsgLogoutInitiator {
	return &MsgLogoutInitiator{
		InitiatorAddress:       initiatorAddress,
		SessionID:              sessionID,
		SessionLogoutInitiator: &text,
	}
}

func (msg *MsgLogoutInitiator) Route() string {
	return RouterKey
}

func (msg *MsgLogoutInitiator) Type() string {
	return TypeMsgLogoutInitiator
}

func (msg *MsgLogoutInitiator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogoutInitiator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogoutInitiator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgLogoutAcceptor = "logout_acceptor"

var _ sdk.Msg = &MsgLogoutAcceptor{}

func NewMsgLogoutAcceptor(acceptorAddress string, sessionID string, text SessionLogoutAcceptor) *MsgLogoutAcceptor {
	return &MsgLogoutAcceptor{
		AcceptorAddress:       acceptorAddress,
		SessionID:             sessionID,
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
