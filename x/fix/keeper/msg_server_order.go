package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// NewOrderSingle creates a New Single Order in Re Protocol
func (k msgServer) NewOrderSingle(goCtx context.Context, msg *types.MsgNewOrderSingle) (*types.MsgNewOrderSingleResponse, error) {
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
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to an Order
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check if order exists
	order, found := k.GetOrders(ctx, msg.ClOrdID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order exist already", &order))
	}

	// check that these mandatory fields are not empty
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Symbol: %s", msg.Symbol)
	}
	if msg.OrderQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrderQty: %s", msg.OrderQty)
	}
	if msg.Price == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Price: %s", msg.Price)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.Side), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Side: %v", msg.Side)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.OrdType), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrdType: %v", msg.OrdType)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.TimeInForce), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "TimeInForce: %v", msg.TimeInForce)
	}

	newOrder := types.Orders{
		SessionID:    session.SessionID,
		ClOrdID:      msg.ClOrdID,
		Symbol:       msg.Symbol,
		Side:         msg.Side,
		OrderQty:     msg.OrderQty,
		OrdType:      msg.OrdType,
		Price:        msg.Price,
		TimeInForce:  msg.TimeInForce,
		Text:         msg.Text,
		TransactTime: msg.TransactTime,
	}

	// fetch Header from existing session
	// In the FIX Protocol, New Single Order message can be sent by either the initiator or the acceptor in the FIX session.
	// Determine whether it is the initiator or acceptor
	var header *types.Header
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, D is the message type for New Single Order
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating New Single Order
	newOrder.Header = header
	newOrder.Header.MsgType = "D"
	newOrder.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session
	var trailer *types.Trailer
	if session.LogonInitiator.Header.SenderCompID == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	newOrder.Trailer = trailer

	// set the new order single to store
	k.SetOrders(ctx, msg.ClOrdID, newOrder)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgNewOrderSingleResponse{}, err
}

// OrderCancelRequest requests for the cancellation of New Single Order
func (k msgServer) OrderCancelRequest(goCtx context.Context, msg *types.MsgOrderCancelRequest) (*types.MsgOrderCancelRequestResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Order Cancel Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check if order exists
	order, found := k.GetOrders(ctx, msg.OrigClOrdID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order doesn't exist", &order))
	}
	if msg.Creator != order.Header.SenderCompID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account address is not the creator of the order", &order))
	}

	// check that this order cancellation has not been requested already
	orderCancelRequest, found := k.GetOrdersCancelRequest(ctx, msg.OrigClOrdID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s The creator has requested to cancel this order already", &orderCancelRequest))
	}

	// An OrderCancelRequest can only be made if the New Single Order has not been executed or rejected
	orderExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.OrigClOrdID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsExecutedAlready, "Order Execution Report: %s", &orderExecutionReport)
	}

	orderCancelReject, found := k.GetOrdersCancelReject(ctx, msg.OrigClOrdID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order is Rejected already", &orderCancelReject))
	}

	// set order cancel request
	newOrderCancelRequest := types.OrdersCancelRequest{
		SessionID:   msg.SessionID,
		Header:      order.Header,
		OrigClOrdID: msg.OrigClOrdID,
		ClOrdID:     msg.ClOrdID,
		Trailer:     order.Trailer,
	}

	// msgType = F
	newOrderCancelRequest.Header.MsgType = "F"

	// set order cancel request to store
	k.SetOrdersCancelRequest(ctx, msg.OrigClOrdID, newOrderCancelRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderCancelRequestResponse{}, err
}

// OrderExecutionReport creates Order Execution Report
func (k msgServer) OrderExecutionReport(goCtx context.Context, msg *types.MsgOrderExecutionReport) (*types.MsgOrderExecutionReportResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Order Execution Report
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check if order exists
	order, found := k.GetOrders(ctx, msg.ClOrdID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order doesn't exist", &order))
	}

	// same account address can not used for creating New Single Order and Execution Report with the same ClOrdID
	if msg.Creator == order.Header.SenderCompID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create execution report", msg.Creator))
	}

	// An OrderExecutionReport can only be made if the New Single Order has not been executed or rejected
	orderExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.ClOrdID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsExecutedAlready, "Order Execution Report: %s", &orderExecutionReport)
	}

	orderCancelReject, found := k.GetOrdersCancelReject(ctx, msg.ClOrdID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order is Rejected already", &orderCancelReject))
	}

	// check that these mandatory fields are not empty
	if msg.OrderID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrderID: %s", msg.OrderID)
	}
	if msg.ExecID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "msg.ExecID: %s", msg.ExecID)
	}
	if msg.OrdStatus == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrdStatus: %s", msg.OrdStatus)
	}
	if msg.ExecType == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "ExecType: %s", msg.ExecType)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Symbol: %s", msg.Symbol)
	}
	if msg.OrderQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrderQty: %s", msg.OrderQty)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.Side), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Side: %v", msg.Side)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.TimeInForce), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "TimeInForce: %v", msg.TimeInForce)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.Price), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "Price: %v", msg.Price)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.LastPx), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "LastPx: %v", msg.LastPx)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.LastQty), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "LastQty: %v", msg.LastQty)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.LeavesQty), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "LeavesQty: %v", msg.LeavesQty)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.CumQty), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "CumQty: %v", msg.CumQty)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.AvgPx), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "AvgPx: %v", msg.AvgPx)
	}

	newOrderExecutionReport := types.OrdersExecutionReport{
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
	}

	// set header from the existing New Single Order
	newOrderExecutionReport.Header = order.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *newOrderExecutionReport.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = order.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = order.Header.MsgSeqNum

	// set the msgType to Order Execution Report
	newHeader.MsgType = "8"

	// switch senderCompID and targetCompID between New Single Order and Order Execution Report
	// set senderCompID of Order Execution Report to the targetCompID of New Single Order in the header
	newHeader.SenderCompID = order.Header.TargetCompID

	// set targetCompID of Order Execution Report to the senderCompID of New Single Order in the header
	newHeader.TargetCompID = order.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	newOrderExecutionReport.Header = newHeader

	// set Trailer from the existing New Single Order
	// TODO
	// checksum should be recalcualted
	newOrderExecutionReport.Trailer = order.Trailer

	// set order execution report to store
	k.SetOrdersExecutionReport(ctx, msg.ClOrdID, newOrderExecutionReport)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderExecutionReportResponse{}, err
}

// OrderCancelReject creates Order Cancel Reject
func (k msgServer) OrderCancelReject(goCtx context.Context, msg *types.MsgOrderCancelReject) (*types.MsgOrderCancelRejectResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Order Cancel Reject
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check if order exists
	order, found := k.GetOrders(ctx, msg.OrigClOrdID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order doesn't exist", &order))
	}

	// same account can not used to create New Single Order and to Reject the order with the same ClOrdID
	if msg.Creator == order.Header.SenderCompID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to Reject this order", &order))
	}

	// An OrderCancelRequest can only be made if the New Single Order has not been executed or rejected
	orderExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.OrigClOrdID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsExecutedAlready, "Order Execution Report: %s", &orderExecutionReport)
	}

	orderCancelReject, found := k.GetOrdersCancelReject(ctx, msg.OrigClOrdID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Order is Rejected already", &orderCancelReject))
	}

	// check that these mandatory fields are not empty
	if msg.OrderID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "OrderID: %s", msg.OrderID)
	}
	if _, err := strconv.ParseInt(fmt.Sprint(msg.CxlRejReason), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderEmptyField, "CxlRejReason: %v", msg.CxlRejReason)
	}

	newOrderCancelReject := types.OrdersCancelReject{
		SessionID:        msg.SessionID,
		OrderID:          msg.OrderID,
		OrigClOrdID:      msg.OrigClOrdID,
		ClOrdID:          msg.ClOrdID,
		CxlRejReason:     msg.CxlRejReason,
		CxlRejResponseTo: msg.CxlRejResponseTo,
		TransactTime:     msg.TransactTime,
	}

	// set header from the existing New Single Order
	newOrderCancelReject.Header = order.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *newOrderCancelReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = order.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = order.Header.MsgSeqNum

	// set the msgType to Order Execution Report
	newHeader.MsgType = "9"

	// switch senderCompID and targetCompID between New Single Order and Order Cancel Reject
	// set senderCompID of Order Cancel Reject to the targetCompID of New Single Order in the header
	newHeader.SenderCompID = order.Header.TargetCompID

	// set targetCompID of Order Cancel Reject to the senderCompID of New Single Order in the header
	newHeader.TargetCompID = order.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// pass all the edited values to the newHeader
	newOrderCancelReject.Header = newHeader

	// set Trailer from the existing New Single Order
	// TODO
	// checksum should be recalcualted
	newOrderCancelReject.Trailer = order.Trailer

	k.SetOrdersCancelReject(ctx, msg.OrigClOrdID, newOrderCancelReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderCancelRejectResponse{}, err
}
