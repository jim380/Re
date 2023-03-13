package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/mic/types"
)

func (k msgServer) RegisterMarketIdentificationCode(goCtx context.Context, msg *types.MsgRegisterMarketIdentificationCode) (*types.MsgRegisterMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//TODO
	//include checks

	var marketIdentificationCode = types.MarketIdentificationCode{
		Creator:             msg.Creator,
		MIC:                 msg.MIC,
		Name:                msg.Name,
		Location:            msg.Location,
		AssetClass:          msg.AssetClass,
		Currency:            msg.Currency,
		RegulatoryAuthority: msg.RegulatoryAuthority,
		Status:              msg.Status,
	}

  //set MIC Data to store
  k.SetMarketIdentificationCode(ctx, msg.MIC, marketIdentificationCode)

	return &types.MsgRegisterMarketIdentificationCodeResponse{}, nil
}

func (k msgServer) UpdateMarketIdentificationCode(goCtx context.Context, msg *types.MsgUpdateMarketIdentificationCode) (*types.MsgUpdateMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var marketIdentificationCode = types.MarketIdentificationCode{
		Creator:             msg.Creator,
		MIC:                 msg.MIC,
		Name:                msg.Name,
		Location:            msg.Location,
		AssetClass:          msg.AssetClass,
		Currency:            msg.Currency,
		RegulatoryAuthority: msg.RegulatoryAuthority,
		Status:              msg.Status,
	}

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.MIC))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMarketIdentificationCode(ctx, msg.MIC, marketIdentificationCode)

	return &types.MsgUpdateMarketIdentificationCodeResponse{}, nil
}

func (k msgServer) DeleteMarketIdentificationCode(goCtx context.Context, msg *types.MsgDeleteMarketIdentificationCode) (*types.MsgDeleteMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.MIC))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMarketIdentificationCode(ctx, msg.MIC)

	return &types.MsgDeleteMarketIdentificationCodeResponse{}, nil
}
