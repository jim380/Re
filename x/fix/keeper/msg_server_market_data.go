package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) MarketDataRequest(goCtx context.Context, msg *types.MsgMarketDataRequest) (*types.MsgMarketDataRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataRequestResponse{}, nil
}

func (k msgServer) MarketDataSnapshotFullRefresh(goCtx context.Context, msg *types.MsgMarketDataSnapshotFullRefresh) (*types.MsgMarketDataSnapshotFullRefreshResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataSnapshotFullRefreshResponse{}, nil
}

func (k msgServer) MarketDataIncremental(goCtx context.Context, msg *types.MsgMarketDataIncremental) (*types.MsgMarketDataIncrementalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataIncrementalResponse{}, nil
}

func (k msgServer) MarketDataRequestReject(goCtx context.Context, msg *types.MsgMarketDataRequestReject) (*types.MsgMarketDataRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMarketDataRequestRejectResponse{}, nil
}
