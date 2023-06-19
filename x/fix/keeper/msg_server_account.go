package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// CreateAccount creates account for users of Re Protocol
func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	for _, accounts := range k.GetAllAccount(ctx) {

		// check for if the provided company name is not taken
		if accounts.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}

		// check for if the provided website is not taken
		if accounts.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}

		// check for if account address is not used already
		if accounts.Address == msg.Address {
			return nil, sdkerrors.Wrapf(types.ErrAccountIsTaken, "Account: %s", msg.Address)
		}
	}

	newAccount := types.Account{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
	}

	// set new account data to store
	k.SetAccount(ctx, msg.Address, newAccount)

	return &types.MsgCreateAccountResponse{}, nil
}

// UpdateAccount updates account for users of Re Protocol
func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	for _, accounts := range k.GetAllAccount(ctx) {
		// Checks if an account exists
		if accounts.Empty() {
			return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
		}
		if accounts.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}
		if accounts.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}
		// check for if account address is not used already
		if accounts.Address != msg.Creator {
			return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Creator)
		}
	}

	editedAccount := types.Account{
		Address:          msg.Address,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
	}

	// set edited account data to store
	k.SetAccount(ctx, msg.Address, editedAccount)

	return &types.MsgUpdateAccountResponse{}, nil
}

// DeleteAccount deletes an account for users of Re Protocol
func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// Checks if an account exists
	existingAccount := k.GetAccount(ctx, msg.Address)
	if existingAccount.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Address)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Address != existingAccount.Address {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Address)
	}

	k.RemoveAccount(ctx, msg.Address)

	return &types.MsgDeleteAccountResponse{}, nil
}
