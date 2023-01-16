package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetAccount(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
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
