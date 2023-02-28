package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgLogonInitiator) FIXVersionByInitiator() string {
	msg.LogonInitiator.Header.BeginString = fixVersion
	return msg.LogonInitiator.Header.BeginString
}

func NewHeader(bodyLength int64, msgType string, senderCompID string, targetCompID string, msgSeqNum int64, sendingTime string) Header {
	return Header{
		BodyLength:   bodyLength,
		MsgType:      msgType,
		SenderCompID: senderCompID,
		TargetCompID: targetCompID,
		MsgSeqNum:    msgSeqNum,
		SendingTime:  sendingTime,
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
