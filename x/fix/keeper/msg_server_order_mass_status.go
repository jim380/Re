package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) OrderMassStatusRequest(goCtx context.Context, msg *types.MsgOrderMassStatusRequest) (*types.MsgOrderMassStatusRequestResponse, error) {
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
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to send Order Mass Status
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Order
	order, found := k.GetOrders(ctx, msg.ClOrdID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsNotFound, "Order: %s", &order)
	}

	// check that an order has not been executed
	// order mass status request should only pass if the order has not been executed
	ordersExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.ClOrdID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrOrderIsExecutedAlready, "Order Execution Report: %s", &ordersExecutionReport)
	}

	// check that the owner of the order matches the creator of order mass status request
	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Order Mass Status Creator: %s", msg.Creator)
	}

	// check that any of these mandatory fields is not empty
	if order.Symbol != msg.Symbol {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "Symbol: %s", msg.Symbol)
	}
	if msg.MassStatusReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "MassStatusReqID: %s", msg.MassStatusReqID)
	}
	if msg.MassStatusReqType == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "MassStatusReqType: %s", msg.MassStatusReqType)
	}
	if msg.Account == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "Account: %s", msg.Account)
	}
	if msg.SecurityID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "SecurityID: %s", msg.SecurityID)
	}
	if msg.TradingSessionID == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "TradingSessionID : %s", msg.TradingSessionID)
	}

	// order mass status request
	orderMassStatusRequest := types.OrderMassStatus{
		SessionID: msg.SessionID,
		OrderMassStatusRequest: &types.OrderMassStatusRequest{
			MassStatusReqID:   msg.MassStatusReqID,
			MassStatusReqType: msg.MassStatusReqType,
			ClOrdID:           msg.ClOrdID,
			Account:           msg.Account,
			Symbol:            msg.Symbol,
			SecurityID:        msg.SecurityID,
			TradingSessionID:  msg.TradingSessionID,
			Creator:           msg.Creator,
		},
	}

	// set header from the existing order
	// since order mass status request is created by the same account
	orderMassStatusRequest.OrderMassStatusRequest.Header = order.Header

	// TODO
	// Recalculate the bodyLength, msgSeqNum in the header

	// set the msgType to order mass status request
	orderMassStatusRequest.OrderMassStatusRequest.Header.MsgType = "AF"

	// set sending time
	orderMassStatusRequest.OrderMassStatusRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// set Trailer from the existing order
	// TODO
	// checksum should be recalculated
	orderMassStatusRequest.OrderMassStatusRequest.Trailer = order.Trailer

	// set order mass status request to store
	k.SetOrderMassStatus(ctx, msg.MassStatusReqID, orderMassStatusRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderMassStatusRequestResponse{}, err
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
