package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTerminateLogon = "terminate_logon"

var _ sdk.Msg = &MsgTerminateLogon{}

func NewMsgTerminateLogon(initiatorAddress string, sessionID string, did string) *MsgTerminateLogon {
	return &MsgTerminateLogon{
		InitiatorAddress: initiatorAddress,
		SessionID:      sessionID,
		Did:              did,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
