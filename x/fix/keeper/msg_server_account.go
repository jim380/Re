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

	//get DID Document from the DID module && check if DID is exists
	getDidDocument := k.didKeeper.GetDIDDocument(ctx, msg.Did)
	if getDidDocument.Document.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.Did)
	}

	getAcc := k.GetAccount(ctx, msg.Did)
	if getAcc.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.Creator)
	}
	if getDidDocument.Creator != getAcc.Creator {
		return nil, sdkerrors.Wrapf(types.ErrAccountUserIsNotTheSame, "Account: %s", msg.Creator)
	}

	var account = types.Account{
		Creator:          msg.Creator,
		Did:              msg.Did,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        int32(time.Now().Unix()),
	}

	// check if DID is used already
	if getAcc.Did == account.Did {
		return nil, sdkerrors.Wrapf(types.ErrDIDIsTaken, "DID: %s", msg.Did)
	}

	// no account should be created with same COMPANY NAME
	if getAcc.CompanyName == account.CompanyName {
		return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
	}
 
	k.SetAccount(ctx, msg.Did, account)

	return &types.MsgCreateAccountResponse{}, nil
}

func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var account = types.Account{
		Creator:          msg.Creator,
		Did:              msg.Did,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        int32(time.Now().Unix()),
	}

	// Checks that the element exists
	val := k.GetAccount(ctx, msg.Did)
	//if !val {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	//}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAccount(ctx, msg.Did, account)

	return &types.MsgUpdateAccountResponse{}, nil
}

func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val := k.GetAccount(ctx, msg.Did)
	//if !found {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	//}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//k.RemoveAccount(ctx, msg.Did)

	return &types.MsgDeleteAccountResponse{}, nil
}
