package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityListRequestReject = "security_list_request_reject"

var _ sdk.Msg = &MsgSecurityListRequestReject{}

func NewMsgSecurityListRequestReject(creator string, sessionID string, securityReqID string, securityListRequestType string, securityRequestResult string, text string, encodedTextLen string, encodedText string) *MsgSecurityListRequestReject {
	return &MsgSecurityListRequestReject{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityReqID:           securityReqID,
		SecurityListRequestType: securityListRequestType,
		SecurityRequestResult:   securityRequestResult,
		Text:                    text,
		EncodedTextLen:          encodedTextLen,
		EncodedText:             encodedText,
	}
}

func (msg *MsgSecurityListRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityListRequestReject) Type() string {
	return TypeMsgSecurityListRequestReject
}

func (msg *MsgSecurityListRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityListRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityListRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
