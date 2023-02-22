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

	var account = types.Account{
		Creator:          msg.Creator,
		Did:              msg.Did,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		CreatedAt:        int32(time.Now().Unix()),
	}

	//creator of DID Document should be same with creator of Account
	if getDidDocument.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrAccountUserIsNotTheSame, "Account: %s", msg.Creator)
	}

	//getAcc := k.GetAccount(ctx, msg.Did)
	//if getAcc.Did == account.Did {
	//	return nil, sdkerrors.Wrapf(types.ErrDIDIsTaken, "DID: %s", msg.Did)
	//}

	//check for if the provided company name is not taken
	getCompanyName := k.GetAccountCompanyName(ctx, msg.CompanyName)
	if getCompanyName == account.CompanyName {
		return nil, sdkerrors.Wrapf(types.ErrCompanyNameIsTaken, "Company Name: %s", msg.CompanyName)
	}

	// check for if the provided website is not taken
	getWebsite := k.GetAccountWebsite(ctx, msg.Website)
	if getWebsite == account.Website {
		return nil, sdkerrors.Wrapf(types.ErrWebsiteIstaken, "Website: %s", msg.Website)
	}

	// get creator of the account
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	getAddr := k.GetAccountCreator(ctx)
	if getAddr.Equals(creator) {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsTaken, "Account: %s", msg.Creator)
	}

	k.SetAccount(ctx, msg.Did, account)
	k.SetAccountCompanyName(ctx, msg.CompanyName, account)
	k.SetAccountWebsite(ctx, msg.Website, account)
	k.SetAccountCreator(ctx, creator)

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
