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
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
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
		Creator:      msg.Creator,
	}

	// fetch Header from existing session
	// In the FIX Protocol, New Single Order message can be sent by either the initiator or the acceptor in the FIX session.
	// Determine whether it is the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
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
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	newOrder.Trailer = trailer

	// set the new order single to store
	k.SetOrders(ctx, msg.SessionID, newOrder)

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
