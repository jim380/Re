package keeper

import (
	"context"
	"fmt"
	"log"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid Address")
	}

	//getAddr := k.GetAccountCreator(ctx)

	// Check if the value already exists
	//accountExists := k.GetAccount(ctx, msg.Did)

	//check if DID is valid from the FIX module

	getDID := k.didKeeper.GetDIDDocument(ctx, msg.Did)
	log.Println(getDID)

	var account = types.Account{
		Creator:          msg.Creator,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		DID:              msg.Did,
		CreatedAt:        int32(time.Now().Unix()),
	}

	/*if accountExists.DID == account.DID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DID already set")
	}

	if getAddr.Equals(addr) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "You Are Not Allowed to Have a double Account with the same Wallet Adddress")
	}

	if accountExists.CompanyName == account.CompanyName {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Account Name Already Exists")
	}
	*/

	//set account to store if all checks passes
	k.SetAccount(ctx, msg.Did, account)
	k.SetAccountCreator(ctx, addr)

	return &types.MsgCreateAccountResponse{}, nil
}

func (k msgServer) UpdateAccount(goCtx context.Context, msg *types.MsgUpdateAccount) (*types.MsgUpdateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var account = types.Account{
		Creator:          msg.Creator,
		CompanyName:      msg.CompanyName,
		Website:          msg.Website,
		SocialMediaLinks: msg.SocialMediaLinks,
		DID:              msg.Did,
	}

	// Checks that the element exists
	val := k.GetAccount(ctx, msg.Did)
	if val.Empty() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID %d doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAccount(ctx, val.DID, account)

	return &types.MsgUpdateAccountResponse{}, nil
}

func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	acc := k.GetAccount(ctx, msg.Did)
	if acc.Empty() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID %d doesn't exist", msg.Did))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != acc.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	//k.RemoveAccount(ctx, msg.Did)

	return &types.MsgDeleteAccountResponse{}, nil
}
