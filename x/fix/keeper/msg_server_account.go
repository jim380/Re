package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	accountExists, isFound := k.GetAccount(
		ctx,
		msg.Id,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var account = types.Account{
		Id:               msg.Id,
		Creator:          msg.Creator,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		DID:              msg.DID,
		CreatedAt:        int32(time.Now().Unix()),
	}

	if accountExists.Creator == account.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You Are Not Allowed to Have a double Account with the same Wallet Adddress")
	}

	if accountExists.CompanyName == account.CompanyName {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Account Name Already Exists")
	}

	id := k.AppendAccount(
		ctx,
		account,
	)

	return &types.MsgCreateAccountResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var account = types.Account{
		Creator:          msg.Creator,
		Id:               msg.Id,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		DID:              msg.DID,
	}

	// Checks that the element exists
	val, found := k.GetAccount(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAccount(ctx, account)

	return &types.MsgUpdateAccountResponse{}, nil
}

func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAccount(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAccount(ctx, msg.Id)

	return &types.MsgDeleteAccountResponse{}, nil
}
