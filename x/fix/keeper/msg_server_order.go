package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) NewOrderSingle(goCtx context.Context, msg *types.MsgNewOrderSingle) (*types.MsgNewOrderSingleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// TODO
	// get session, check if session exists and if the status is set to loggedIn
	// check for account address
	// include checks for all the FIX Protocol messages
	// add more check cases

	// check for if this session Name exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	if session.Status != "loggedIn" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s state is not logged in", msg.SessionID))
	}

	// check if order exists
	order, found := k.GetOrders(ctx, msg.SessionID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order exist already", &order))
	}

	newOrder := types.Orders{
		SessionID:    session.SessionID,
		Header:       session.LogonInitiator.Header,
		ClOrdID:      msg.ClOrdID,
		Symbol:       msg.Symbol,
		Side:         msg.Side,
		OrderQty:     msg.OrderQty,
		OrdType:      msg.OrdType,
		Price:        msg.Price,
		TimeInForce:  msg.TimeInForce,
		Text:         msg.Text,
		TransactTime: msg.TransactTime,
		Trailer:      session.LogonInitiator.Trailer,
		Status:       msg.Status,
		Creator:      msg.Creator,
	}

	// set the msgType
	newOrder.Header.MsgType = "D"
	newOrder.TransactTime = time.Now().UTC().Format("20060102-15:04:05.000")
	newOrder.Status = "Requested"

	k.SetOrders(ctx, msg.SessionID, newOrder)

	return &types.MsgNewOrderSingleResponse{}, nil
}

func (k msgServer) OrderCancelRequest(goCtx context.Context, msg *types.MsgOrderCancelRequest) (*types.MsgOrderCancelRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message
	// Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator

	// set order cancel request data
	orderCancelRequest := types.OrdersCancelRequest{
		SessionID:   msg.SessionID,
		Header:      session.LogonInitiator.Header,
		OrigClOrdID: msg.OrigClOrdID,
		ClOrdID:     msg.ClOrdID,
		Trailer:     session.LogonInitiator.Trailer,
		Creator:     msg.Creator,
	}

	// msgType = F
	orderCancelRequest.Header.MsgType = "F"

	// set order cancel request to store
	k.SetOrdersCancelRequest(ctx, msg.SessionID, orderCancelRequest)

	return &types.MsgOrderCancelRequestResponse{}, nil
}

func (k msgServer) OrderExecutionReport(goCtx context.Context, msg *types.MsgOrderExecutionReport) (*types.MsgOrderExecutionReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message err
	// Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator
	// checks for the fields in the FIX Protocol

	orderExecutionReport := types.OrdersExecutionReport{
		SessionID:    msg.SessionID,
		Header:       session.LogonAcceptor.Header,
		ClOrdID:      msg.ClOrdID,
		OrderID:      msg.OrderID,
		ExecID:       msg.ExecID,
		OrdStatus:    msg.OrdStatus,
		ExecType:     msg.ExecType,
		Symbol:       msg.Symbol,
		Side:         msg.Side,
		OrderQty:     msg.OrderQty,
		Price:        msg.Price,
		TimeInForce:  msg.TimeInForce,
		LastPx:       msg.LastPx,
		LastQty:      msg.LastQty,
		LeavesQty:    msg.LeavesQty,
		CumQty:       msg.CumQty,
		AvgPx:        msg.AvgPx,
		Text:         msg.Text,
		TransactTime: msg.TransactTime,
		Trailer:      session.LogonAcceptor.Trailer,
		Creator:      msg.Creator,
	}

	orderExecutionReport.Header.MsgType = "8"
	orderExecutionReport.TransactTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// set order execution report to store
	k.SetOrdersExecutionReport(ctx, msg.SessionID, orderExecutionReport)

	return &types.MsgOrderExecutionReportResponse{}, nil
}

func (k msgServer) OrderCancelReject(goCtx context.Context, msg *types.MsgOrderCancelReject) (*types.MsgOrderCancelRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	session, _ := k.GetSessions(ctx, msg.SessionID)

	// TODO: Handling the message
	// Fetch session
	// check for existing origClOrdID
	// check that account address equals the creator
	// Handle all FIX protocol message

	orderCancelReject := types.OrdersCancelReject{
		SessionID:        msg.SessionID,
		Header:           session.LogonAcceptor.Header,
		OrderID:          msg.OrderID,
		OrigClOrdID:      msg.OrigClOrdID,
		ClOrdID:          msg.ClOrdID,
		CxlRejReason:     msg.CxlRejReason,
		CxlRejResponseTo: msg.CxlRejResponseTo,
		TransactTime:     msg.TransactTime,
		Trailer:          session.LogonAcceptor.Trailer,
		Creator:          msg.Creator,
	}

	// set msgType
	orderCancelReject.Header.MsgType = "9"
	orderCancelReject.TransactTime = time.Now().UTC().Format("20060102-15:04:05.000")

	k.SetOrdersCancelReject(ctx, msg.SessionID, orderCancelReject)

	return &types.MsgOrderCancelRejectResponse{}, nil
}
