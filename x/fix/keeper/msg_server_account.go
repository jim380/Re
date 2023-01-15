package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the Account already exists
	accountExists, isFound := k.GetAccount(
		ctx,
		msg.Index,
	)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "This Account Already Exists")
	}

	var account = types.Account{
		Creator:            msg.Creator,
		Index:              msg.Index,
		CompanyName:        msg.CompanyName,
		Address:            msg.Address,
		EmailAddress:       msg.EmailAddress,
		PhoneNumber:        msg.PhoneNumber,
		Website:            msg.Website,
		SocialMediaLinks:   msg.SocialMediaLinks,
		GovernmentIssuedId: msg.GovernmentIssuedId,
		CreatedAt:          msg.CreatedAt,
	}
    
	if accountExists.Creator == account.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You Are Not Allowed to Have a double Account with the same Wallet Adddress")
	}

	if accountExists.CompanyName == account.CompanyName {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Account Name Already Exists")
	}

	if accountExists.EmailAddress == account.EmailAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Email Address Already Exists")
	}

	if accountExists.PhoneNumber == account.PhoneNumber {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Phone Number Already Exists")
	}

	//To Do
	// include checks for GOVERNMENT ISSUED ID
	//Also check to see if all user's GOVERNMENT ISSUED ID can be passed through governance proposal mechanism for authentic verification of the documents by the community

	//set a new account if all checks passes
	k.SetAccount(
		ctx,
		account,
	)
	return &types.MsgCreateAccountResponse{}, nil
}

func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAccount(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var account = types.Account{
		Creator:            msg.Creator,
		Index:              msg.Index,
		CompanyName:        msg.CompanyName,
		Address:            msg.Address,
		EmailAddress:       msg.EmailAddress,
		PhoneNumber:        msg.PhoneNumber,
		Website:            msg.Website,
		SocialMediaLinks:   msg.SocialMediaLinks,
		GovernmentIssuedId: msg.GovernmentIssuedId,
		CreatedAt:          msg.CreatedAt,
	}

	k.SetAccount(ctx, account)

	return &types.MsgUpdateAccountResponse{}, nil
}

func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAccount(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAccount(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteAccountResponse{}, nil
}
