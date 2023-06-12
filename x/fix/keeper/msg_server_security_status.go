package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) SecurityStatusRequest(goCtx context.Context, msg *types.MsgSecurityStatusRequest) (*types.MsgSecurityStatusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusRequestResponse{}, nil
}

func (k msgServer) SecurityStatusResponse(goCtx context.Context, msg *types.MsgSecurityStatusResponse) (*types.MsgSecurityStatusResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusResponseResponse{}, nil
}

func (k msgServer) SecurityStatusRequestReject(goCtx context.Context, msg *types.MsgSecurityStatusRequestReject) (*types.MsgSecurityStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusRequestRejectResponse{}, nil
}
