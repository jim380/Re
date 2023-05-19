package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityDefinitionRequest = "security_definition_request"

var _ sdk.Msg = &MsgSecurityDefinitionRequest{}

func NewMsgSecurityDefinitionRequest(creator string, securityReqID string, securityRequestType string, symbol string, securityExchange string, issuer string, securityDesc string, securityType string, currency string) *MsgSecurityDefinitionRequest {
	return &MsgSecurityDefinitionRequest{
		Creator:             creator,
		SecurityReqID:       securityReqID,
		SecurityRequestType: securityRequestType,
		Symbol:              symbol,
		SecurityExchange:    securityExchange,
		Issuer:              issuer,
		SecurityDesc:        securityDesc,
		SecurityType:        securityType,
		Currency:            currency,
	}
}

func (msg *MsgSecurityDefinitionRequest) Route() string {
	return RouterKey
}

func (msg *MsgSecurityDefinitionRequest) Type() string {
	return TypeMsgSecurityDefinitionRequest
}

func (msg *MsgSecurityDefinitionRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityDefinitionRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityDefinitionRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
