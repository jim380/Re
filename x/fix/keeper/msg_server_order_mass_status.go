package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// OrderMassStatusRequest Creates Order Mass Status Request
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
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to send Order Mass Status Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
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
	if order.Header.SenderCompID != msg.Creator {
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
	orderMassStatusRequest.OrderMassStatusRequest.Header.SendingTime = constants.SendingTime

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

// OrderMassStatusReport creates Order Mass Status Report
func (k msgServer) OrderMassStatusReport(goCtx context.Context, msg *types.MsgOrderMassStatusReport) (*types.MsgOrderMassStatusReportResponse, error) {
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

	// check that the sessionID is active between both parties
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to send Order Mass Status Report
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Order Mass Status Request
	orderMassStatus, found := k.GetOrderMassStatus(ctx, msg.MassStatusReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusIsNotFound, "Order: %s", &orderMassStatus)
	}

	// the account that created order mass status request is not allowed to respond to order mass status report
	if orderMassStatus.OrderMassStatusRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusCreatorIsWrong, "Order Mass Status Report Creator: %s", msg.Creator)
	}

	// check that Order Mass Status Request has not been responded to
	if orderMassStatus.OrderMassStatusReport != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusRequestIsAcknowledged, "Order mass status Report: %s", orderMassStatus.OrderMassStatusReport)
	}

	// check that Order Mass Status Request has not been rejected
	if orderMassStatus.OrderMassStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusRequestIsRejected, "Order mass status Request Reject: %s", orderMassStatus.OrderMassStatusRequestReject)
	}

	// check that the mandatory fields match the values from Order Mass Status Request
	if orderMassStatus.OrderMassStatusRequest.ClOrdID != msg.ClOrdID {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "ClOrdID : %s", msg.ClOrdID)
	}
	if orderMassStatus.OrderMassStatusRequest.Symbol != msg.Symbol {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "Symbol: %s", msg.Symbol)
	}
	if orderMassStatus.OrderMassStatusRequest.Account == msg.Account {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "Account: %s", msg.Account)
	}
	if orderMassStatus.OrderMassStatusRequest.SecurityID == msg.SecurityID {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "SecurityID: %s", msg.SecurityID)
	}
	if orderMassStatus.OrderMassStatusRequest.TradingSessionID == msg.TradingSessionID {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusMismatchField, "TradingSessionID : %s", msg.TradingSessionID)
	}

	// check that any of these mandatory fields is not empty
	if msg.OrdStatus == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "OrdStatus: %s", msg.OrdStatus)
	}
	if msg.ExecType == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "ExecType: %s", msg.ExecType)
	}
	if msg.OrdQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "OrdQty: %s", msg.OrdQty)
	}
	if msg.LastPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "LastPx: %s", msg.LastPx)
	}
	if msg.CumQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "CumQty: %s", msg.CumQty)
	}
	if msg.AvgPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "AvgPx: %s", msg.AvgPx)
	}
	if msg.LeavesQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "LeavesQty: %s", msg.LeavesQty)
	}

	// order mass status report
	orderMassStatusReport := types.OrderMassStatus{
		SessionID:              msg.SessionID,
		OrderMassStatusRequest: orderMassStatus.OrderMassStatusRequest,
		OrderMassStatusReport: &types.OrderMassStatusReport{
			ClOrdID:          msg.ClOrdID,
			MassStatusReqID:  msg.MassStatusReqID,
			Account:          msg.Account,
			Symbol:           msg.Symbol,
			SecurityID:       msg.SecurityID,
			TradingSessionID: msg.TradingSessionID,
			OrdStatus:        msg.OrdStatus,
			ExecType:         msg.ExecType,
			OrdQty:           msg.OrdQty,
			LastPx:           msg.LastPx,
			CumQty:           msg.CumQty,
			AvgPx:            msg.AvgPx,
			LeavesQty:        msg.LeavesQty,
		},
	}

	// set header from order mass status request
	orderMassStatusReport.OrderMassStatusReport.Header = orderMassStatus.OrderMassStatusRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *orderMassStatusReport.OrderMassStatusReport.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = orderMassStatus.OrderMassStatusRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = orderMassStatus.OrderMassStatusRequest.Header.MsgSeqNum

	// set the msgType to Order Mass Status Report (AN)
	newHeader.MsgType = "AN"

	// switch senderCompID and targetCompID between Order Mass Status Request and Order Mass Status Report
	// set senderCompID of Order Mass Status Report to the targetCompID of Order Mass Status Request in the header
	newHeader.SenderCompID = orderMassStatus.OrderMassStatusRequest.Header.TargetCompID

	// set targetCompID of Order Mass Status Report to the senderCompID of Order Mass Status Request in the header
	newHeader.TargetCompID = orderMassStatus.OrderMassStatusRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	orderMassStatusReport.OrderMassStatusReport.Header = newHeader

	// set Trailer from the existing Order Mass Status Request
	// TODO
	// checksum should be recalcualted
	orderMassStatusReport.OrderMassStatusReport.Trailer = orderMassStatus.OrderMassStatusRequest.Trailer

	// set order mass status report to store
	k.SetOrderMassStatus(ctx, msg.MassStatusReqID, orderMassStatusReport)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderMassStatusReportResponse{}, err
}

// OrderMassStatusRequestReject creates Order Mass Status Request Reject
func (k msgServer) OrderMassStatusRequestReject(goCtx context.Context, msg *types.MsgOrderMassStatusRequestReject) (*types.MsgOrderMassStatusRequestRejectResponse, error) {
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

	// check that the sessionID is active between both parties
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to send Order Mass Status Request Reject
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Order Mass Status Request
	orderMassStatus, found := k.GetOrderMassStatus(ctx, msg.RefSeqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusIsNotFound, "Order: %s", &orderMassStatus)
	}

	// the account that created order mass status request is not allowed to respond to order mass status request reject
	if orderMassStatus.OrderMassStatusRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusCreatorIsWrong, "Order Mass Status Request Reject Creator: %s", msg.Creator)
	}

	// check that Order Mass Status Request has not been responded to
	if orderMassStatus.OrderMassStatusReport != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusRequestIsAcknowledged, "Order mass status Report: %s", orderMassStatus.OrderMassStatusReport)
	}

	// check that Order Mass Status Request has not been rejected
	if orderMassStatus.OrderMassStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusRequestIsRejected, "Order mass status Request Reject: %s", orderMassStatus.OrderMassStatusRequestReject)
	}

	// check that any of these mandatory fields is not empty
	if msg.RejCode == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "RejCode: %s", msg.RejCode)
	}
	if msg.Text == "" {
		return nil, sdkerrors.Wrapf(types.ErrOrderMassStatusEmptyField, "Text: %s", msg.Text)
	}

	// order mass status request reject
	orderMassStatusRequestReject := types.OrderMassStatus{
		SessionID:              msg.SessionID,
		OrderMassStatusRequest: orderMassStatus.OrderMassStatusRequest,
		OrderMassStatusReport:  orderMassStatus.OrderMassStatusReport,
		OrderMassStatusRequestReject: &types.OrderMassStatusRequestReject{
			RefSeqID: msg.RefSeqID,
			RejCode:  msg.RejCode,
			Text:     msg.Text,
		},
	}

	// set header from order mass status request
	orderMassStatusRequestReject.OrderMassStatusRequestReject.Header = orderMassStatus.OrderMassStatusRequest.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *orderMassStatusRequestReject.OrderMassStatusRequestReject.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = orderMassStatus.OrderMassStatusRequest.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = orderMassStatus.OrderMassStatusRequest.Header.MsgSeqNum

	// set the msgType to Order Mass Status Request Reject (AR)
	newHeader.MsgType = "AR"

	// switch senderCompID and targetCompID between Order Mass Status Request and Order Mass Status Request Reject
	// set senderCompID of Order Mass Status Request Reject to the targetCompID of Order Mass Status Request in the header
	newHeader.SenderCompID = orderMassStatus.OrderMassStatusRequest.Header.TargetCompID

	// set targetCompID of Order Mass Status Request Reject to the senderCompID of Order Mass Status Request in the header
	newHeader.TargetCompID = orderMassStatus.OrderMassStatusRequest.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	orderMassStatusRequestReject.OrderMassStatusRequestReject.Header = newHeader

	// set Trailer from the existing Order Mass Status Request
	// TODO
	// checksum should be recalcualted
	orderMassStatusRequestReject.OrderMassStatusRequestReject.Trailer = orderMassStatus.OrderMassStatusRequest.Trailer

	// set order mass status request reject to store
	k.SetOrderMassStatus(ctx, msg.RefSeqID, orderMassStatusRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOrderMassStatusRequestRejectResponse{}, err
}
