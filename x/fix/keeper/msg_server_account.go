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

	// get DID Document from the DID module && check if DID is exists
	getDidDocument := k.didKeeper.GetDIDDocument(ctx, msg.Did)
	if getDidDocument.Document.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.Did)
	}

	// creator of DID Document should be same with creator of Account
	if getDidDocument.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotDIDCreator, "Account: %s", msg.Creator)
	}

	for _, did := range k.getDIDs(ctx) {

		existingAccount := k.GetAccount(ctx, did)

		if existingAccount.Did == msg.Did {
			return nil, sdkerrors.Wrapf(types.ErrDIDIsTaken, "DID: %s", msg.Did)
		}

		// check for if the provided company name is not taken
		if existingAccount.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}

		// check for if the provided website is not taken
		if existingAccount.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}

		// check for if account address is not used already
		if existingAccount.Creator == msg.Creator {
			return nil, sdkerrors.Wrapf(types.ErrAccountIsTaken, "Account: %s", msg.Creator)
		}
	}

	newAccount := types.Account{
		Creator:          msg.Creator,
		Did:              msg.Did,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
	}

	// set new account data to store
	k.SetAccount(ctx, msg.Did, newAccount)

	return &types.MsgCreateAccountResponse{}, nil
}

// UpdateAccount updates account for users of Re Protocol
func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get DID Document from the DID module && check if DID is exists
	getDidDocument := k.didKeeper.GetDIDDocument(ctx, msg.Did)
	if getDidDocument.Document.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.Did)
	}
	// creator of DID Document should be same with creator of Account
	if getDidDocument.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotDIDCreator, "Account: %s", msg.Creator)
	}

	for _, did := range k.getDIDs(ctx) {
		// Checks if an account exists
		existingAccount := k.GetAccount(ctx, did)
		if existingAccount.Empty() {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.Did)
		}
		if existingAccount.CompanyName == msg.CompanyName {
			return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
		}
		if existingAccount.Website == msg.Website {
			return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
		}
	}

	editedAccount := types.Account{
		Creator:          msg.Creator,
		Did:              msg.Did,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        time.Now().UTC().Format("20060102-15:04:05.000"),
	}

	// set edited account data to store
	k.SetAccount(ctx, msg.Did, editedAccount)

	return &types.MsgUpdateAccountResponse{}, nil
}

// DeleteAccount deletes an account for users of Re Protocol
func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks if an account exists
	existingAccount := k.GetAccount(ctx, msg.Did)
	if existingAccount.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.Did)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != existingAccount.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.Creator)
	}

	k.RemoveAccount(ctx, msg.Did)

	return &types.MsgDeleteAccountResponse{}, nil
}
