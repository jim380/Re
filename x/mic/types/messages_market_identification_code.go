package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMarketIdentificationCode = "create_market_identification_code"
	TypeMsgUpdateMarketIdentificationCode = "update_market_identification_code"
	TypeMsgDeleteMarketIdentificationCode = "delete_market_identification_code"
)

var _ sdk.Msg = &MsgCreateMarketIdentificationCode{}

func NewMsgCreateMarketIdentificationCode(creator string, mIC string, name string, location string, assetClass string, currency string, regulatoryAuthority string, status string) *MsgCreateMarketIdentificationCode {
	return &MsgCreateMarketIdentificationCode{
		Creator:             creator,
		MIC:                 mIC,
		Name:                name,
		Location:            location,
		AssetClass:          assetClass,
		Currency:            currency,
		RegulatoryAuthority: regulatoryAuthority,
		Status:              status,
	}
}

func (msg *MsgCreateMarketIdentificationCode) Route() string {
	return RouterKey
}

func (msg *MsgCreateMarketIdentificationCode) Type() string {
	return TypeMsgCreateMarketIdentificationCode
}

func (msg *MsgCreateMarketIdentificationCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMarketIdentificationCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMarketIdentificationCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMarketIdentificationCode{}

func NewMsgUpdateMarketIdentificationCode(creator string, id uint64, mIC string, name string, location string, assetClass string, currency string, regulatoryAuthority string, status string) *MsgUpdateMarketIdentificationCode {
	return &MsgUpdateMarketIdentificationCode{
		Id:                  id,
		Creator:             creator,
		MIC:                 mIC,
		Name:                name,
		Location:            location,
		AssetClass:          assetClass,
		Currency:            currency,
		RegulatoryAuthority: regulatoryAuthority,
		Status:              status,
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

func NewMsgDeleteMarketIdentificationCode(creator string, id uint64) *MsgDeleteMarketIdentificationCode {
	return &MsgDeleteMarketIdentificationCode{
		Id:      id,
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
