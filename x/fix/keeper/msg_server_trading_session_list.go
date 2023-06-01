package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradingSessionListRequest(goCtx context.Context, msg *types.MsgTradingSessionListRequest) (*types.MsgTradingSessionListRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionListRequestResponse{}, nil
}


func (k msgServer) TradingSessionListResponse(goCtx context.Context, msg *types.MsgTradingSessionListResponse) (*types.MsgTradingSessionListResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionListResponseResponse{}, nil
}

func (k msgServer) TradingSessionListRequestReject(goCtx context.Context, msg *types.MsgTradingSessionListRequestReject) (*types.MsgTradingSessionListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionListRequestRejectResponse{}, nil
}
