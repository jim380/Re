package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSecurityDefinitionRequest = "security_definition_request"

var _ sdk.Msg = &MsgSecurityDefinitionRequest{}

func NewMsgSecurityDefinitionRequest(creator string, sessionID string, securityReqID string, securityRequestType string, symbol string, securityExchange string, issuer string, securityDesc string, securityType string, currency string) *MsgSecurityDefinitionRequest {
	return &MsgSecurityDefinitionRequest{
		Creator:             creator,
		SessionID:           sessionID,
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

const TypeMsgSecurityDefinition = "security_definition"

var _ sdk.Msg = &MsgSecurityDefinition{}

func NewMsgSecurityDefinition(creator string, sessionID string, securityReqID string, securityResponseID string, securityResponseType string, symbol string, securityExchange string, issuer string, securityDesc string, securityType string, currency string, contractMultiplier string, minPriceIncrement string, minPriceIncrementAmount string, unitOfMeasure string, priceUnitOfMeasure string, settlType string, settlDate string, maturityMonthYear string, couponRate string, factor string, creditRating string, securityExchangeID string) *MsgSecurityDefinition {
	return &MsgSecurityDefinition{
		Creator:                 creator,
		SessionID:               sessionID,
		SecurityReqID:           securityReqID,
		SecurityResponseID:      securityResponseID,
		SecurityResponseType:    securityResponseType,
		Symbol:                  symbol,
		SecurityExchange:        securityExchange,
		Issuer:                  issuer,
		SecurityDesc:            securityDesc,
		SecurityType:            securityType,
		Currency:                currency,
		ContractMultiplier:      contractMultiplier,
		MinPriceIncrement:       minPriceIncrement,
		MinPriceIncrementAmount: minPriceIncrementAmount,
		UnitOfMeasure:           unitOfMeasure,
		PriceUnitOfMeasure:      priceUnitOfMeasure,
		SettlType:               settlType,
		SettlDate:               settlDate,
		MaturityMonthYear:       maturityMonthYear,
		CouponRate:              couponRate,
		Factor:                  factor,
		CreditRating:            creditRating,
		SecurityExchangeID:      securityExchangeID,
	}
}

func (msg *MsgSecurityDefinition) Route() string {
	return RouterKey
}

func (msg *MsgSecurityDefinition) Type() string {
	return TypeMsgSecurityDefinition
}

func (msg *MsgSecurityDefinition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSecurityDefinition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSecurityDefinition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
