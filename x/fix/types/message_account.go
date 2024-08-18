package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterAccount = "register_account"

var _ sdk.Msg = &MsgRegisterAccount{}

func NewMsgRegisterAccount(creator string, address string, companyName string, website string, socialMediaLinks string) *MsgRegisterAccount {
	return &MsgRegisterAccount{
		Creator:          creator,
		Address:          address,
		CompanyName:      companyName,
		Website:          website,
		SocialMediaLinks: socialMediaLinks,
	}
}

func (msg *MsgRegisterAccount) Route() string {
	return RouterKey
}

func (msg *MsgRegisterAccount) Type() string {
	return TypeMsgRegisterAccount
}

func (msg *MsgRegisterAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgUpdateAccount = "update_account"

var _ sdk.Msg = &MsgUpdateAccount{}

func NewMsgUpdateAccount(creator string, address string, companyName string, website string, socialMediaLinks string) *MsgUpdateAccount {
	return &MsgUpdateAccount{
		Creator:          creator,
		Address:          address,
		CompanyName:      companyName,
		Website:          website,
		SocialMediaLinks: socialMediaLinks,
	}
}

func (msg *MsgUpdateAccount) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAccount) Type() string {
	return TypeMsgUpdateAccount
}

func (msg *MsgUpdateAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgDeleteAccount = "delete_account"

var _ sdk.Msg = &MsgDeleteAccount{}

func NewMsgDeleteAccount(creator string, address string) *MsgDeleteAccount {
	return &MsgDeleteAccount{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgDeleteAccount) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAccount) Type() string {
	return TypeMsgDeleteAccount
}

func (msg *MsgDeleteAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func EmptyAddress(address string) bool {
	return address == ""
}

func (a AccountRegistration) Empty() bool {
	return EmptyAddress(a.Address)
}
