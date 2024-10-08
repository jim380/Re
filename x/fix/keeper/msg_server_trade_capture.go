package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradeCaptureReport(goCtx context.Context, msg *types.MsgTradeCaptureReport) (*types.MsgTradeCaptureReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Trade Capture Report
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Order Execution Report
	ordersExecutionReport, found := k.GetOrdersExecutionReport(ctx, msg.OrderID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrdersExecutionReportIsNotFound, "Order Execution Report: %s", msg.OrderID)
	}

	// check that the owner of the order execution report(the executing party) matches the creator of trade capture report
	if ordersExecutionReport.Header.SenderCompID != msg.Creator {
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
		return nil, sdkerrors.Wrapf(types.ErrTradeReportIDIsEmpty, "TradeReportID: %s", msg.TradeReportID)
	}
	if msg.TradeReportTransType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeReportTransTypeIsEmpty, "TradeReportTransType: %s", msg.TradeReportTransType)
	}
	if msg.TrdType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTrdTypeIsEmpty, "TrdType: %s", msg.TrdType)
	}
	if msg.Side == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSideIsEmpty, "Side: %s", msg.Side)
	}
	if msg.OrderQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureOrderQtyIsEmpty, "OrderQty: %s", msg.OrderQty)
	}
	if msg.LastQty == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureLastQtyIsEmpty, "LastQty: %s", msg.LastQty)
	}
	if msg.LastPx == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureLastPxIsEmpty, "LastPx: %s", msg.LastPx)
	}
	if msg.GrossTradeAmt == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureGrossTradeAmtIsEmpty, "GrossTradeAmt: %s", msg.GrossTradeAmt)
	}
	if msg.Symbol == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSymbolIsEmpty, "Symbol: %s", msg.Symbol)
	}
	if msg.SecurityID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSecurityIDIsEmpty, "SecurityID: %s", msg.SecurityID)
	}
	if msg.TradeDate == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureTradeDateIsEmpty, "TradeDate: %s", msg.TradeDate)
	}
	if msg.TransactTime == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureTransactTimeIsEmpty, "TransactTime: %s", msg.TransactTime)
	}

	// trade capture report
	tradeCaptureReport := types.TradeCapture{
		SessionID: msg.SessionID,
		TradeCaptureReport: &types.TradeCaptureReport{
			TradeReportID:        msg.TradeReportID,
			TradeReportTransType: msg.TradeReportTransType,
			TradeReportType:      msg.TradeReportType,
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
	tradeCaptureReport.TradeCaptureReport.Header.SendingTime = constants.SendingTime

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
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided matches the trade capture report sessionID
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Trade Capture Report Acknowledgement
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Trade Capture
	tradeCapture, found := k.GetTradeCapture(ctx, msg.TradeReportID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureIsNotFound, ": %s", tradeCapture.TradeCaptureReport.TradeReportID)
	}

	// same account can not used for creating trade capture report and trade capture report acknowledgement with the same TradeReportID
	if tradeCapture.TradeCaptureReport.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create trade capture report acknowledgement", &tradeCapture))
	}

	// check that the mandatory fields match the values from Trade Capture Report
	if tradeCapture.TradeCaptureReport.TradeReportID != msg.TradeReportID {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TradeReportID: %s", msg.TradeReportID)
	}
	if tradeCapture.TradeCaptureReport.TradeReportType != msg.TradeReportType {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TradeReportType: %s", msg.TradeReportType)
	}
	if tradeCapture.TradeCaptureReport.TrdType != msg.TrdType {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TrdType: %s", msg.TrdType)
	}
	if tradeCapture.TradeCaptureReport.TrdSubType != msg.TrdSubType {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TrdSubType: %s", msg.TrdSubType)
	}
	if tradeCapture.TradeCaptureReport.TradeID != msg.TradeID {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TradeID: %s", msg.TradeID)
	}

	// check that the Trade Capture Report is not rejected already
	if tradeCapture.TradeCaptureReportRejection != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureReportIsRejected, "Trade Capture: %s", tradeCapture.TradeCaptureReportRejection)
	}

	// check that the Trade Capture Report is not acknowledged already
	if tradeCapture.TradeCaptureReportAcknowledgement != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureReportIsAcknowledged, "Trade Capture: %s", tradeCapture.TradeCaptureReportAcknowledgement)
	}

	// check that the mandatory Trade Capture Report Acknowledgement fields are not empty
	if msg.ExecType == "" {
		return nil, sdkerrors.Wrapf(types.ErrExecTypeIsEmpty, "ExecType: %s", msg.ExecType)
	}
	if msg.TradeReportStatus == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradeReportStatusIsEmpty, "TradeReportStatus : %s", msg.TradeReportStatus)
	}

	// trade capture report acknowledgement
	tradeCaptureReportAcknowledgement := types.TradeCapture{
		SessionID:          msg.SessionID,
		TradeCaptureReport: tradeCapture.TradeCaptureReport,
		TradeCaptureReportAcknowledgement: &types.TradeCaptureReportAcknowledgement{
			TradeReportID:           msg.TradeReportID,
			TradeID:                 msg.TradeID,
			SecondaryTradeID:        msg.SecondaryTradeID,
			TradeReportType:         msg.TradeReportType,
			TrdType:                 msg.TrdType,
			TrdSubType:              msg.TrdSubType,
			ExecType:                msg.ExecType,
			TradeReportRefID:        msg.TradeReportRefID,
			SecondaryTradeReportID:  msg.SecondaryTradeReportID,
			TradeReportStatus:       msg.TradeReportStatus,
			TradeTransType:          msg.TradeTransType,
			TradeReportRejectReason: msg.TradeReportRejectReason,
			Text:                    msg.Text,
		},
	}

	// set header from the trade capture report
	tradeCaptureReportAcknowledgement.TradeCaptureReportAcknowledgement.Header = tradeCapture.TradeCaptureReport.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *tradeCaptureReportAcknowledgement.TradeCaptureReportAcknowledgement.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = tradeCapture.TradeCaptureReport.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = tradeCapture.TradeCaptureReport.Header.MsgSeqNum

	// set the msgType to Trade Capture Report Acknowledgement
	newHeader.MsgType = "AR"

	// switch senderCompID and targetCompID between Trade Capture Report and Trade Capture Report Acknowledgement
	// set senderCompID of Trade Capture Report Acknowledgement to the targetCompID of Trade Capture Report in the header
	newHeader.SenderCompID = tradeCapture.TradeCaptureReport.Header.TargetCompID

	// set targetCompID of Trade Capture Report Acknowledgement to the senderCompID of Trade Capture Report in the header
	newHeader.TargetCompID = tradeCapture.TradeCaptureReport.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	tradeCaptureReportAcknowledgement.TradeCaptureReportAcknowledgement.Header = newHeader

	// set Trailer from the existing Trade Capture Report
	// checksum should be recalcualted
	tradeCaptureReportAcknowledgement.TradeCaptureReportAcknowledgement.Trailer = tradeCapture.TradeCaptureReport.Trailer

	// set Trade Capture Report Acknowledgement to store
	k.SetTradeCapture(ctx, msg.TradeReportID, tradeCaptureReportAcknowledgement)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradeCaptureReportAcknowledgementResponse{}, err
}

func (k msgServer) TradeCaptureReportRejection(goCtx context.Context, msg *types.MsgTradeCaptureReportRejection) (*types.MsgTradeCaptureReportRejectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided matches the trade capture report sessionID
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureSession, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create Trade Capture Report Rejection
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// get Trade Capture
	tradeCapture, found := k.GetTradeCapture(ctx, msg.TradeReportID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureIsNotFound, ": %s", tradeCapture.TradeCaptureReport.TradeReportID)
	}

	// same account can not used for creating trade capture report and trade capture report rejection with the same TradeReportID
	if tradeCapture.TradeCaptureReport.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create trade capture report acknowledgement", &tradeCapture))
	}

	// check that the mandatory field matches the value from Trade Capture Report
	if tradeCapture.TradeCaptureReport.TradeReportID != msg.TradeReportID {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureMismatchField, "TradeReportID: %s", msg.TradeReportID)
	}

	// check that the Trade Capture Report is not rejected already
	if tradeCapture.TradeCaptureReportRejection != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureReportIsRejected, "Trade Capture: %s", tradeCapture.TradeCaptureReportRejection)
	}

	// check that the Trade Capture Report is not acknowledged already
	if tradeCapture.TradeCaptureReportAcknowledgement != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradeCaptureReportIsAcknowledged, "Trade Capture: %s", tradeCapture.TradeCaptureReportAcknowledgement)
	}

	// check that the mandatory Trade Capture Report Acknowledgement fields are not empty
	if _, err := strconv.ParseInt(fmt.Sprint(msg.TradeReportRejectReason), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradeReportRejectReasonIsEmpty, "TradeReportRejectReason: %v", msg.TradeReportRejectReason)
	}

	// trade capture report rejection
	tradeCaptureReportRejection := types.TradeCapture{
		SessionID:                         msg.SessionID,
		TradeCaptureReport:                tradeCapture.TradeCaptureReport,
		TradeCaptureReportAcknowledgement: tradeCapture.TradeCaptureReportAcknowledgement,
		TradeCaptureReportRejection: &types.TradeCaptureReportRejection{
			TradeReportID:           msg.TradeReportID,
			TradeReportRejectReason: msg.TradeReportRejectReason,
			TradeReportRejectRefID:  msg.TradeReportRejectRefID,
			Text:                    msg.Text,
		},
	}

	// set header from the trade capture report
	tradeCaptureReportRejection.TradeCaptureReportRejection.Header = tradeCapture.TradeCaptureReport.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *tradeCaptureReportRejection.TradeCaptureReportRejection.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = tradeCapture.TradeCaptureReport.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = tradeCapture.TradeCaptureReport.Header.MsgSeqNum

	// set the msgType to Trade Capture Report Rejection
	newHeader.MsgType = "j"

	// switch senderCompID and targetCompID between Trade Capture Report and Trade Capture Report Rejection
	// set senderCompID of Trade Capture Report Rejection to the targetCompID of Trade Capture Report in the header
	newHeader.SenderCompID = tradeCapture.TradeCaptureReport.Header.TargetCompID

	// set targetCompID of Trade Capture Report Rejection to the senderCompID of Trade Capture Report in the header
	newHeader.TargetCompID = tradeCapture.TradeCaptureReport.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	tradeCaptureReportRejection.TradeCaptureReportRejection.Header = newHeader

	// set Trailer from the existing Trade Capture Report
	// checksum should be recalcualted
	tradeCaptureReportRejection.TradeCaptureReportRejection.Trailer = tradeCapture.TradeCaptureReport.Trailer

	// set Trade Capture Report Acknowledgement to store
	k.SetTradeCapture(ctx, msg.TradeReportID, tradeCaptureReportRejection)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradeCaptureReportRejectionResponse{}, err
}
