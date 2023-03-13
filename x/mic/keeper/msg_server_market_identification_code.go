package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/mic/types"
)

func (k msgServer) CreateMarketIdentificationCode(goCtx context.Context, msg *types.MsgCreateMarketIdentificationCode) (*types.MsgCreateMarketIdentificationCodeResponse, error) {
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

	id := k.AppendMarketIdentificationCode(
		ctx,
		marketIdentificationCode,
	)

	return &types.MsgCreateMarketIdentificationCodeResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMarketIdentificationCode(goCtx context.Context, msg *types.MsgUpdateMarketIdentificationCode) (*types.MsgUpdateMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var marketIdentificationCode = types.MarketIdentificationCode{
		Creator:             msg.Creator,
		Id:                  msg.Id,
		MIC:                 msg.MIC,
		Name:                msg.Name,
		Location:            msg.Location,
		AssetClass:          msg.AssetClass,
		Currency:            msg.Currency,
		RegulatoryAuthority: msg.RegulatoryAuthority,
		Status:              msg.Status,
	}

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMarketIdentificationCode(ctx, marketIdentificationCode)

	return &types.MsgUpdateMarketIdentificationCodeResponse{}, nil
}

func (k msgServer) DeleteMarketIdentificationCode(goCtx context.Context, msg *types.MsgDeleteMarketIdentificationCode) (*types.MsgDeleteMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMarketIdentificationCode(ctx, msg.Id)

	return &types.MsgDeleteMarketIdentificationCodeResponse{}, nil
}
