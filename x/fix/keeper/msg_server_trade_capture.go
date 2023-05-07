package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradeCaptureReport(goCtx context.Context, msg *types.MsgTradeCaptureReport) (*types.MsgTradeCaptureReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportResponse{}, nil
}

func (k msgServer) TradeCaptureReportAcknowledgement(goCtx context.Context, msg *types.MsgTradeCaptureReportAcknowledgement) (*types.MsgTradeCaptureReportAcknowledgementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportAcknowledgementResponse{}, nil
}

func (k msgServer) TradeCaptureReportRejection(goCtx context.Context, msg *types.MsgTradeCaptureReportRejection) (*types.MsgTradeCaptureReportRejectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradeCaptureReportRejectionResponse{}, nil
}
