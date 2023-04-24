package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradeCaptureReport(goCtx context.Context, msg *types.MsgTradeCaptureReport) (*types.MsgTradeCaptureReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportResponse{}, nil
}

func (k msgServer) TradeCaptureReportAcknowledgement(goCtx context.Context, msg *types.MsgTradeCaptureReportAcknowledgement) (*types.MsgTradeCaptureReportAcknowledgementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportAcknowledgementResponse{}, nil
}

func (k msgServer) TradeCaptureReportRejection(goCtx context.Context, msg *types.MsgTradeCaptureReportRejection) (*types.MsgTradeCaptureReportRejectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportRejectionResponse{}, nil
}
