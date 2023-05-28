package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradingSessionStatusRequest(goCtx context.Context, msg *types.MsgTradingSessionStatusRequest) (*types.MsgTradingSessionStatusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionStatusRequestResponse{}, nil
}

func (k msgServer) TradingSessionStatus(goCtx context.Context, msg *types.MsgTradingSessionStatus) (*types.MsgTradingSessionStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionStatusResponse{}, nil
}

func (k msgServer) TradingSessionStatusRequestReject(goCtx context.Context, msg *types.MsgTradingSessionStatusRequestReject) (*types.MsgTradingSessionStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionStatusRequestRejectResponse{}, nil
}
