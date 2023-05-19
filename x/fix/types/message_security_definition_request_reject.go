package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityDefinitionRequestReject = "security_definition_request_reject"

var _ sdk.Msg = &MsgSecurityDefinitionRequestReject{}

func NewMsgSecurityDefinitionRequestReject(creator string, sessionID string, securityReqID string, securityRequestResult string, securityRequestError string, securityRequestErrorCode string, text string) *MsgSecurityDefinitionRequestReject {
	return &MsgSecurityDefinitionRequestReject{
		Creator:                  creator,
		SessionID:                sessionID,
		SecurityReqID:            securityReqID,
		SecurityRequestResult:    securityRequestResult,
		SecurityRequestError:     securityRequestError,
		SecurityRequestErrorCode: securityRequestErrorCode,
		Text:                     text,
	}
}

func (msg *MsgSecurityDefinitionRequestReject) Route() string {
	return RouterKey
}

func (msg *MsgSecurityDefinitionRequestReject) Type() string {
	return TypeMsgSecurityDefinitionRequestReject
}

func (msg *MsgSecurityDefinitionRequestReject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityDefinitionRequestReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityDefinitionRequestReject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
