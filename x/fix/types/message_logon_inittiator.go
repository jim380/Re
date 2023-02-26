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

func (msg *MsgLogonInitiator) FIXVersion() string {
	msg.LogonInitiator.Header.BeginString = fixVersion
	return msg.LogonInitiator.Header.BeginString
}

func (msg *MsgLogonInitiator) BodyLength(msgLength string) int64 {
	//sample of the msg
	// "8=FIX.4.4|9=59|35=A|34=1|49=SenderCompID|52=20230219-10:30:00.000|56=TargetCompID|98=0|108=30|141=Y|10=118|"
	logonMsg := msgLength
	length := len(logonMsg)
	msg.LogonInitiator.Header.BodyLength = int64(length)
	return msg.LogonInitiator.Header.BodyLength
}
