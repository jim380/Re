package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) OrderMassStatusRequest(goCtx context.Context, msg *types.MsgOrderMassStatusRequest) (*types.MsgOrderMassStatusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgOrderMassStatusRequestResponse{}, nil
}

func (k msgServer) OrderMassStatusReport(goCtx context.Context, msg *types.MsgOrderMassStatusReport) (*types.MsgOrderMassStatusReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgOrderMassStatusReportResponse{}, nil
}

func (k msgServer) OrderMassStatusRequestReject(goCtx context.Context, msg *types.MsgOrderMassStatusRequestReject) (*types.MsgOrderMassStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgOrderMassStatusRequestRejectResponse{}, nil
}