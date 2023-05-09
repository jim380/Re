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

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != types.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrMarketDataSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Trade Capture Report
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Order Execution Report
	ordersExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.OrderID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsNotFound, "Order Execution Report: %s", msg.OrderID)
	}

	// check that the owner of the order execution report(the executing party) matches the creator of trade capture report
	if ordersExecutionReport.Creator != msg.Creator {
		// handle err
	}

	// check that these fields from order execution report match the provided values
	if ordersExecutionReport.OrderID != msg.OrderID {
		// handle err
	}
	if ordersExecutionReport.ExecID != msg.ExecID {
		// handle err
	}

	

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
