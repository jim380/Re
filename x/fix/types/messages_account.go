package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAccount = "create_account"
	TypeMsgUpdateAccount = "update_account"
	TypeMsgDeleteAccount = "delete_account"
)

var _ sdk.Msg = &MsgCreateAccount{}

func NewMsgCreateAccount(creator string, address string, companyName string, website string, socialMediaLinks string) *MsgCreateAccount {
	return &MsgCreateAccount{
		Creator:          creator,
		Address:          address,
		CompanyName:      companyName,
		Website:          website,
		SocialMediaLinks: socialMediaLinks,
	}
}

func (msg *MsgCreateAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateAccount) Type() string {
	return TypeMsgCreateAccount
}

func (msg *MsgCreateAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAccount{}

func NewMsgDeleteAccount(creator string, address string) *MsgDeleteAccount {
	return &MsgDeleteAccount{
		Address: address,
		Creator: creator,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func EmptyDID(did string) bool {
	return did == ""
}

func (a Account) Empty() bool {
	return EmptyDID(a.Address)
}
