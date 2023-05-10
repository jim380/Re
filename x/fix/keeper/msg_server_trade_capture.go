package keeper

import (
	"context"
	"time"

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

	// trade capture report
	tradeCaptureReport := types.TradeCapture{
		SessionID: msg.SessionID,
		TradeCaptureReport: &types.TradeCaptureReport{
			TradeReportID:        msg.TradeReportID,
			TradeReportTransType: msg.TradeReportTransType,
			TradeReportType:      msg.TradeReportType,
			TradeRequestID:       msg.TradeRequestID,
			TrdType:              msg.TrdType,
			TrdSubType:           msg.TrdSubType,
			Side:                 msg.Side,
			OrderQty:             msg.OrderQty,
			LastQty:              msg.LastQty,
			LastPx:               msg.LastPx,
			GrossTradeAmt:        msg.GrossTradeAmt,
			ExecID:               msg.ExecID,
			OrderID:              msg.OrderID,
			TradeID:              msg.TradeID,
			OrigTradeID:          msg.OrigTradeID,
			Symbol:               msg.Symbol,
			SecurityID:           msg.SecurityID,
			SecurityIDSource:     msg.SecurityIDSource,
			TradeDate:            msg.TradeDate,
			TransactTime:         msg.TransactTime,
			SettlType:            msg.SettlType,
			SettlDate:            msg.SettlDate,
		},
	}

	// set header from the existing order execution report
	// since trade capture report is created by the same account
	tradeCaptureReport.TradeCaptureReport.Header = ordersExecutionReport.Header

	// TODO
	// Recalculate the bodyLength, msgSeqNum in the header

	// set the msgType to trade capture report
	tradeCaptureReport.TradeCaptureReport.Header.MsgType = "AE"

	// set sending time
	tradeCaptureReport.TradeCaptureReport.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// set Trailer from the existing order execution report
	// checksum should be recalculated
	tradeCaptureReport.TradeCaptureReport.Trailer = ordersExecutionReport.Trailer

	// set Trade Capture Report to store
	k.SetTradeCapture(ctx, msg.TradeReportID, tradeCaptureReport)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradeCaptureReportResponse{}, err
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
