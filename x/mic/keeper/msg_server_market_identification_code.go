package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/mic/types"
)

// RegisterMarketIdentificationCode registers user's market identification code issued by ISO standard 10383 on the chain
func (k msgServer) RegisterMarketIdentificationCode(goCtx context.Context, msg *types.MsgRegisterMarketIdentificationCode) (*types.MsgRegisterMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that MIC doesn't exist already
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrMICExistsAlready, "MIC: %s", msg.MIC)
	}

	// check if the msg creator is not the same as the owner of any existing MIC
	if msg.Creator == val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsTaken, "Creator: %s", msg.Creator)
	}

	// check these fields are not empty
    if msg.MIC == "" {

	}
	if msg.Name == "" {

	}
	if msg.Location == "" {

	}
	if msg.AssetClass == "" {

	}
    if msg.Currency == "" {

	}

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

	// set MIC Data to store
	k.SetMarketIdentificationCode(ctx, marketIdentificationCode)

	return &types.MsgRegisterMarketIdentificationCodeResponse{}, nil
}

// UpdateMarketIdentificationCode updates user's market identification code issued by ISO standard 10383 on the chain
func (k msgServer) UpdateMarketIdentificationCode(goCtx context.Context, msg *types.MsgUpdateMarketIdentificationCode) (*types.MsgUpdateMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICIsEmpty, "MIC: %s", msg.MIC)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsWrong, "Creator: %s", msg.Creator)
	}

	// TODO
	// Include other checks

	var editedMarketIdentificationCode = types.MarketIdentificationCode{
		Creator:             msg.Creator,
		MIC:                 msg.MIC,
		Name:                msg.Name,
		Location:            msg.Location,
		AssetClass:          msg.AssetClass,
		Currency:            msg.Currency,
		RegulatoryAuthority: msg.RegulatoryAuthority,
		Status:              msg.Status,
	}

	// set edited MIC Data to store
	k.SetMarketIdentificationCode(ctx, editedMarketIdentificationCode)

	return &types.MsgUpdateMarketIdentificationCodeResponse{}, nil
}

// DeleteMarketIdentificationCode deletes user's market identification code from the store
func (k msgServer) DeleteMarketIdentificationCode(goCtx context.Context, msg *types.MsgDeleteMarketIdentificationCode) (*types.MsgDeleteMarketIdentificationCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMarketIdentificationCode(ctx, msg.MIC)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMICIsEmpty, "MIC: %s", msg.MIC)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrapf(types.ErrMICCreatorIsWrong, "Creator: %s", msg.Creator)
	}

	//remove MIC from store
	k.RemoveMarketIdentificationCode(ctx, msg.MIC)

	return &types.MsgDeleteMarketIdentificationCodeResponse{}, nil
}
