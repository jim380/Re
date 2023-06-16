package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityTypesRequestReject = "security_types_request_reject"

var _ sdk.Msg = &MsgSecurityTypesRequestReject{}

func NewMsgSecurityTypesRequestReject(creator string, sessionID string, securityReqID string, rejectReason string, text string) *MsgSecurityTypesRequestReject {
	return &MsgSecurityTypesRequestReject{
		Creator:       creator,
		SessionID:     sessionID,
		SecurityReqID: securityReqID,
		RejectReason:  rejectReason,
		Text:          text,
	}
}

func (msg *MsgSecurityTypesRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityTypesRequestReject) Type() string {
	return TypeMsgSecurityTypesRequestReject
}

func (msg *MsgSecurityTypesRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityTypesRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityTypesRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
