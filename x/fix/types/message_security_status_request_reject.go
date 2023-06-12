package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityStatusRequestReject = "security_status_request_reject"

var _ sdk.Msg = &MsgSecurityStatusRequestReject{}

func NewMsgSecurityStatusRequestReject(creator string, sessionID string, securityStatusReqID string, securityRejectReason string, text string) *MsgSecurityStatusRequestReject {
	return &MsgSecurityStatusRequestReject{
		Creator:              creator,
		SessionID:            sessionID,
		SecurityStatusReqID:  securityStatusReqID,
		SecurityRejectReason: securityRejectReason,
		Text:                 text,
	}
}

func (msg *MsgSecurityStatusRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityStatusRequestReject) Type() string {
	return TypeMsgSecurityStatusRequestReject
}

func (msg *MsgSecurityStatusRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityStatusRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityStatusRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
