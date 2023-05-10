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
		return nil, sdkerrors.Wrapf(types.ErrOrdersExecutionReportIsNotFound, "Order Execution Report: %s", msg.OrderID)
	}

	// check that the owner of the order execution report(the executing party) matches the creator of trade capture report
	if ordersExecutionReport.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Order Execution Report Creator: %s", msg.Creator)
	}

	// check that these fields from order execution report match the provided values
	if ordersExecutionReport.OrderID != msg.OrderID {
		return nil, sdkerrors.Wrapf(types.ErrOrderIDIsNotFound, "OrderID: %s", msg.OrderID)
	}
	if ordersExecutionReport.ExecID != msg.ExecID {
		return nil, sdkerrors.Wrapf(types.ErrExecIDIsNotFound, "ExecID: %s", msg.ExecID)
	}

	// check that mandatory Trade Capture Report fields are not empty
	if msg.TradeReportID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeReportIDIsNotFound, "TradeReportID: %s", msg.TradeReportID)
	}
	if msg.TradeReportTransType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeReportTransTypeIsNotFound, "TradeReportTransType: %s", msg.TradeReportTransType)
	}
	if msg.TrdType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTrdTypeIsNotFound, "TrdType: %s", msg.TrdType)
	}
	if msg.Side == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSideIsNotFound, "Side: %s", msg.Side)
	}
	if msg.OrderQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureOrderQtyIsNotFound, "OrderQty: %s", msg.OrderQty)
	}
	if msg.LastQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureLastQtyIsNotFound, "LastQty: %s", msg.LastQty)
	}
	if msg.LastPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureLastPxIsNotFound, "LastPx: %s", msg.LastPx)
	}
	if msg.GrossTradeAmt == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureGrossTradeAmtIsNotFound, "GrossTradeAmt: %s", msg.GrossTradeAmt)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSymbolIsNotFound, "Symbol: %s", msg.Symbol)
	}
	if msg.SecurityID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSecurityIDIsNotFound, "SecurityID: %s", msg.SecurityID)
	}
	if msg.TradeDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureTradeDateIsNotFound, "TradeDate: %s", msg.TradeDate)
	}
	if msg.TransactTime == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureTransactTimeIsNotFound, "TransactTime: %s", msg.TransactTime)
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
