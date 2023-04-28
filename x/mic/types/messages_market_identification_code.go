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

func NewMsgRegisterMarketIdentificationCode(creator string, mic string, operatingMIC string, oprtSGMT string, marketName string, legalEntityName string, legalEntityIdentifier string, marketCategory string, acronym string, isoCountryCode string, city string, website string, status string, creationDate string, lastUpdateDate string, lastValidationDate string, expiryDate string, comments string) *MsgRegisterMarketIdentificationCode {
	return &MsgRegisterMarketIdentificationCode{
		Creator:               creator,
		MIC:                   mic,
		Operating_MIC:         operatingMIC,
		OPRT_SGMT:             oprtSGMT,
		MarketName:            marketName,
		LegalEntityName:       legalEntityName,
		LegalEntityIdentifier: legalEntityIdentifier,
		MarketCategory:        marketCategory,
		Acronym:               acronym,
		ISOCountryCode:        isoCountryCode,
		City:                  city,
		Website:               website,
		Status:                status,
		CreationDate:          creationDate,
		LastUpdateDate:        lastUpdateDate,
		LastValidationDate:    lastValidationDate,
		ExpiryDate:            expiryDate,
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

func NewMsgUpdateMarketIdentificationCode(creator string, oldMic string, newMic string, operatingMIC string, oprtSGMT string, marketName string, legalEntityName string, legalEntityIdentifier string, marketCategory string, acronym string, isoCountryCode string, city string, website string, status string, creationDate string, lastUpdateDate string, lastValidationDate string, expiryDate string, comments string) *MsgUpdateMarketIdentificationCode {
	return &MsgUpdateMarketIdentificationCode{
		Creator:               creator,
		Old_MIC:               oldMic,
		New_MIC:               newMic,
		Operating_MIC:         operatingMIC,
		OPRT_SGMT:             oprtSGMT,
		MarketName:            marketName,
		LegalEntityName:       legalEntityName,
		LegalEntityIdentifier: legalEntityIdentifier,
		MarketCategory:        marketCategory,
		Acronym:               acronym,
		ISOCountryCode:        isoCountryCode,
		City:                  city,
		Website:               website,
		Status:                status,
		CreationDate:          creationDate,
		LastUpdateDate:        lastUpdateDate,
		LastValidationDate:    lastValidationDate,
		ExpiryDate:            expiryDate,
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
