package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	fixVersion            = "FIX4.4"
	TypeMsgLogonInitiator = "logon_initiator"
)

var _ sdk.Msg = &MsgLogonInitiator{}

func NewMsgLogonInitiator(initiatorAddress string, sessionID string, logonInitator LogonInitiator) *MsgLogonInitiator {
	return &MsgLogonInitiator{
		InitiatorAddress: initiatorAddress,
		SessionID:        sessionID,
		LogonInitiator:   &logonInitator,
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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgLogonInitiator) FIXVersionByInitiator() string {
	msg.LogonInitiator.Header.BeginString = fixVersion
	return msg.LogonInitiator.Header.BeginString
}

func NewHeaderInitiator(bodyLength int64, msgType string, senderCompID string, targetCompID string, msgSeqNum int64, sendingTime string, chainID string) Header {
	return Header{
		BodyLength:   bodyLength,
		MsgType:      msgType,
		SenderCompID: senderCompID,
		TargetCompID: targetCompID,
		MsgSeqNum:    msgSeqNum,
		SendingTime:  sendingTime,
		ChainID:      chainID,
	}
}

func NewTrailer(checkSum int64) Trailer {
	return Trailer{
		CheckSum: checkSum,
	}
}

func NewLogonInitiator(header Header, encryptMethod int64, heartBtInt int64, trailer Trailer) LogonInitiator {
	return LogonInitiator{
		Header:        &header,
		EncryptMethod: encryptMethod,
		HeartBtInt:    heartBtInt,
		Trailer:       &trailer,
	}
}

const TypeMsgLogonAcceptor = "logon_acceptor"

var _ sdk.Msg = &MsgLogonAcceptor{}

func NewMsgLogonAcceptor(acceptorAddress string, sessionID string, logonAcceptor LogonAcceptor) *MsgLogonAcceptor {
	return &MsgLogonAcceptor{
		AcceptorAddress: acceptorAddress,
		SessionID:       sessionID,
		LogonAcceptor:   &logonAcceptor,
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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgLogonAcceptor) FIXVersionByAcceptor() string {
	msg.LogonAcceptor.Header.BeginString = fixVersion
	return msg.LogonAcceptor.Header.BeginString
}

func NewHeaderAcceptor(bodyLength int64, msgType string, senderCompID string, targetCompID string, msgSeqNum int64, sendingTime string) Header {
	return Header{
		BodyLength:   bodyLength,
		MsgType:      msgType,
		SenderCompID: senderCompID,
		TargetCompID: targetCompID,
		MsgSeqNum:    msgSeqNum,
		SendingTime:  sendingTime,
	}
}

func NewLogonAcceptor(header Header, encryptMethod int64, heartBtInt int64, trailer Trailer) LogonAcceptor {
	return LogonAcceptor{
		Header:        &header,
		EncryptMethod: encryptMethod,
		HeartBtInt:    heartBtInt,
		Trailer:       &trailer,
	}
}

const TypeMsgLogonReject = "logon_reject"

var _ sdk.Msg = &MsgLogonReject{}

func NewMsgLogonReject(acceptorAddress string, sessionID string, text string) *MsgLogonReject {
	return &MsgLogonReject{
		AcceptorAddress: acceptorAddress,
		SessionID:       sessionID,
		Text:            text,
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
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTerminateLogon = "terminate_logon"

var _ sdk.Msg = &MsgTerminateLogon{}

func NewMsgTerminateLogon(initiatorAddress string, sessionID string, address string) *MsgTerminateLogon {
	return &MsgTerminateLogon{
		InitiatorAddress: initiatorAddress,
		SessionID:        sessionID,
		Address:          address,
	}
}

func (msg *MsgTerminateLogon) Route() string {
	return RouterKey
}

func (msg *MsgTerminateLogon) Type() string {
	return TypeMsgTerminateLogon
}

func (msg *MsgTerminateLogon) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTerminateLogon) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTerminateLogon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
