package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogonReject = "logon_reject"

var _ sdk.Msg = &MsgLogonReject{}

func NewMsgLogonReject(acceptorAddress string, sessionName string, text string, did Header) *MsgLogonReject {
	return &MsgLogonReject{
		AcceptorAddress: acceptorAddress,
		SessionName:     sessionName,
		Text:            text,
		Header:          &did,
	}
}

func (msg *MsgLogonReject) Route() string {
	return RouterKey
}

func (msg *MsgLogonReject) Type() string {
	return TypeMsgLogonReject
}

func (msg *MsgLogonReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogonReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogonReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
