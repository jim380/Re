package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	for _, accounts := range k.GetAllAccountRegistration(ctx) {

		// check for if account address is not used already
		if accounts.Address == msg.Address {
			return nil, sdkerrors.Wrapf(types.ErrAccountIsTaken, "Account: %s", msg.Address)
		}

		// check for if the provided company name is not taken
		if accounts.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}

		// check for if the provided website is not taken
		if accounts.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}
	}

	newAccount := types.AccountRegistration{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
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

	for _, accounts := range k.GetAllAccountRegistration(ctx) {

		// Checks if an account exists
		if accounts.Empty() {
			return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
		}

		// check that it is only the owner that can update account information
		if accounts.Address != msg.Creator {
			return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Creator)
		}

		if accounts.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}
		if accounts.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}
	}

	editedAccount := types.AccountRegistration{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
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

	// Checks if an account exists
	existingAccount, found := k.GetAccountRegistration(ctx, msg.Address)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != existingAccount.Address {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Address)
	}

	k.RemoveAccountRegistration(ctx, msg.Address)

	return &types.MsgDeleteAccountResponse{}, nil
}
