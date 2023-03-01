package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogonAcceptor = "logon_acceptor"

var _ sdk.Msg = &MsgLogonAcceptor{}

func NewMsgLogonAcceptor(acceptorAddress string, sessionID string, LogonAcceptor LogonAcceptor) *MsgLogonAcceptor {
	return &MsgLogonAcceptor{
		AcceptorAddress: acceptorAddress,
		SessionID:     sessionID,
		LogonAcceptor:   &LogonAcceptor,
	}
}

func (msg *MsgLogonAcceptor) Route() string {
	return RouterKey
}

func (msg *MsgLogonAcceptor) Type() string {
	return TypeMsgLogonAcceptor
}

func (msg *MsgLogonAcceptor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogonAcceptor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogonAcceptor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgLogonAcceptor) FIXVersionByAcceptor() string {
	msg.LogonAcceptor.Header.BeginString = fixVersion
	return msg.LogonAcceptor.Header.BeginString
}

func NewLogonAcceptor(header Header, encryptMethod int64, heartBtInt int64, trailer Trailer) LogonAcceptor {
	return LogonAcceptor{
		Header:        &header,
		EncryptMethod: encryptMethod,
		HeartBtInt:    heartBtInt,
		Trailer:       &trailer,
	}
}
