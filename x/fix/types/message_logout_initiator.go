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
		SessionID:            sessionID,
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
