package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogonInitiator = "logon_initiator"

var _ sdk.Msg = &MsgLogonInitiator{}

func NewMsgLogonInitiator(initiatorAddress string, sessionName string) *MsgLogonInitiator {
	return &MsgLogonInitiator{
		InitiatorAddress: initiatorAddress,
		SessionName:      sessionName,
	}
}

func (msg *MsgLogonInitiator) Route() string {
	return RouterKey
}

func (msg *MsgLogonInitiator) Type() string {
	return TypeMsgLogonInitiator
}

func (msg *MsgLogonInitiator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogonInitiator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogonInitiator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
