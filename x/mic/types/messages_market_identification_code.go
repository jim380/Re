package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRegisterMarketIdentificationCode = "register_market_identification_code"
	TypeMsgUpdateMarketIdentificationCode   = "update_market_identification_code"
	TypeMsgDeleteMarketIdentificationCode   = "delete_market_identification_code"
)

var _ sdk.Msg = &MsgRegisterMarketIdentificationCode{}

func NewMsgRegisterMarketIdentificationCode(creator string, mic string, operating_MIC string, OPRT_SGMT string, market_name string, legal_entity_name string, legal_entity_identifier string, market_category string, acronym string, ISO_country_code string, city string, website string, status string, creation_date string, last_update_date string, last_validation_date string, expiry_date string, comments string) *MsgRegisterMarketIdentificationCode {
	return &MsgRegisterMarketIdentificationCode{
		Creator:               creator,
		MIC:                   mic,
		Operating_MIC:         operating_MIC,
		OPRT_SGMT:             operating_MIC,
		MarketName:            market_name,
		LegalEntityName:       legal_entity_name,
		LegalEntityIdentifier: legal_entity_identifier,
		MarketCategory:        market_category,
		Acronym:               acronym,
		ISOCountryCode:        ISO_country_code,
		City:                  city,
		Website:               website,
		Status:                status,
		CreationDate:          creation_date,
		LastUpdateDate:        last_update_date,
		LastValidationDate:    last_validation_date,
		ExpiryDate:            expiry_date,
		Comments:              comments,
	}
}

func (msg *MsgRegisterMarketIdentificationCode) Route() string {
	return RouterKey
}

func (msg *MsgRegisterMarketIdentificationCode) Type() string {
	return TypeMsgRegisterMarketIdentificationCode
}

func (msg *MsgRegisterMarketIdentificationCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterMarketIdentificationCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterMarketIdentificationCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMarketIdentificationCode{}

func NewMsgUpdateMarketIdentificationCode(creator string, mic string, operating_MIC string, OPRT_SGMT string, market_name string, legal_entity_name string, legal_entity_identifier string, market_category string, acronym string, ISO_country_code string, city string, website string, status string, creation_date string, last_update_date string, last_validation_date string, expiry_date string, comments string) *MsgUpdateMarketIdentificationCode {
	return &MsgUpdateMarketIdentificationCode{
		Creator:               creator,
		MIC:                   mic,
		Operating_MIC:         operating_MIC,
		OPRT_SGMT:             operating_MIC,
		MarketName:            market_name,
		LegalEntityName:       legal_entity_name,
		LegalEntityIdentifier: legal_entity_identifier,
		MarketCategory:        market_category,
		Acronym:               acronym,
		ISOCountryCode:        ISO_country_code,
		City:                  city,
		Website:               website,
		Status:                status,
		CreationDate:          creation_date,
		LastUpdateDate:        last_update_date,
		LastValidationDate:    last_validation_date,
		ExpiryDate:            expiry_date,
		Comments:              comments,
	}
}

func (msg *MsgUpdateMarketIdentificationCode) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMarketIdentificationCode) Type() string {
	return TypeMsgUpdateMarketIdentificationCode
}

func (msg *MsgUpdateMarketIdentificationCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMarketIdentificationCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMarketIdentificationCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMarketIdentificationCode{}

func NewMsgDeleteMarketIdentificationCode(creator string, mic string) *MsgDeleteMarketIdentificationCode {
	return &MsgDeleteMarketIdentificationCode{
		MIC:     mic,
		Creator: creator,
	}
}
func (msg *MsgDeleteMarketIdentificationCode) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMarketIdentificationCode) Type() string {
	return TypeMsgDeleteMarketIdentificationCode
}

func (msg *MsgDeleteMarketIdentificationCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMarketIdentificationCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMarketIdentificationCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
