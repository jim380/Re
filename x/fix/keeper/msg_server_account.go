package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// RegisterAccount creates account for FIX users of Re Protocol
func (k msgServer) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check that the address provided must be the user's address
	if msg.Address != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrAddressNotMatched, "Address: %s", msg.Creator)
	}

	// Check if an account exists
	_, found := k.GetAccountRegistration(ctx, msg.Address)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsTaken, "Account: %s", msg.Address)
	}

	newAccount := types.AccountRegistration{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        constants.CreatedAt,
	}

	// set new account data to store
	k.SetAccountRegistration(ctx, msg.Address, newAccount)

	return &types.MsgRegisterAccountResponse{}, nil
}

// UpdateAccount updates account for FIX users of Re Protocol
func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check that the address provided must be the user's address
	if msg.Address != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrAddressNotMatched, "Address: %s", msg.Creator)
	}

	account, found := k.GetAccountRegistration(ctx, msg.Address)
	// Check if an account exists
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
	}

	// check that it is only the owner that can update account information
	if account.Address != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Creator)
	}

	editedAccount := types.AccountRegistration{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        constants.CreatedAt,
	}

	// set edited account data to store
	k.SetAccountRegistration(ctx, msg.Address, editedAccount)

	return &types.MsgUpdateAccountResponse{}, nil
}

// DeleteAccount deletes an account for FIX users of Re Protocol
func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check that the address provided must be the user's address
	if msg.Address != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrAddressNotMatched, "Address: %s", msg.Creator)
	}

	// Check if an account exists
	account, found := k.GetAccountRegistration(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
	}

	// Check if the msg creator is the same as the current owner
	if msg.Creator != account.Address {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Address)
	}

	k.RemoveAccountRegistration(ctx, msg.Address)

	return &types.MsgDeleteAccountResponse{}, nil
}
