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

// TradingSessionStatusRequest creates Trading Session Status Request
func (k msgServer) TradingSessionStatusRequest(goCtx context.Context, msg *types.MsgTradingSessionStatusRequest) (*types.MsgTradingSessionStatusRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != types.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create trading session status request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that mandatory Trading Session Status Request fields are not empty
	if msg.TradSesReqID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionEmptyField, "TradSesReqID: %s", msg.TradSesReqID)
	}
	if msg.TradingSessionID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionEmptyField, "TradingSessionID: %s", msg.TradingSessionID)
	}
	if msg.TradingSessionSubID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionEmptyField, "TradingSessionSubID: %s", msg.TradingSessionSubID)
	}

	// trading session status request
	tradingSessionStatusRequest := types.TradingSession{
		SessionID: msg.SessionID,
		TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
			TradSesReqID:          msg.TradSesReqID,
			TradingSessionID:      msg.TradingSessionID,
			TradingSessionSubID:   msg.TradingSessionSubID,
			MarketID:              msg.MarketID,
			SubscriptionRequest:   msg.SubscriptionRequest,
			SecurityID:            msg.SecurityID,
			SecurityIDSource:      msg.SecurityIDSource,
			Symbol:                msg.Symbol,
			SecurityExchange:      msg.SecurityExchange,
			MarketSegmentID:       msg.MarketSegmentID,
			TradSesReqType:        msg.TradSesReqType,
			TradSesSubReqType:     msg.TradSesSubReqType,
			TradSesMode:           msg.TradSesMode,
			TradingSessionDate:    msg.TradingSessionDate,
			TradingSessionTime:    msg.TradingSessionTime,
			TradingSessionSubTime: msg.TradingSessionSubTime,
			ExpirationDate:        msg.ExpirationDate,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session Status Request is sent by the client and the client is the party that initiates the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "g" to indicate a Trading Session Status Request.
	// SenderCompID (49) contains the ID of the sender (client).
	// TargetCompID (56) contains the ID of the target (trading venue).
	// TradingSessionID (336) may be used to specify the trading session for which the status is requested.
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session Status Request
	tradingSessionStatusRequest.TradingSessionStatusRequest.Header = session.LogonInitiator.Header
	tradingSessionStatusRequest.TradingSessionStatusRequest.Header.MsgType = "g"
	tradingSessionStatusRequest.TradingSessionStatusRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionStatusRequest.TradingSessionStatusRequest.Trailer = session.LogonInitiator.Trailer

	// set Trading Session Status Request to store
	k.SetTradingSession(ctx, msg.TradSesReqID, tradingSessionStatusRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionStatusRequestResponse{}, err
}

// TradingSessionStatus creates Trading Session Status
func (k msgServer) TradingSessionStatus(goCtx context.Context, msg *types.MsgTradingSessionStatus) (*types.MsgTradingSessionStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Trading Session Status matches with the sessionID for Trading Session Status Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Trading Session Status Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Trading Session Status Creator: %s", msg.Creator)
	}

	// get Trading Session
	tradingSession, found := k.GetTradingSession(ctx, msg.TradSesReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionIsNotFound, ": %s", tradingSession.TradingSessionStatusRequest)
	}

	// same account can not used for creating Trading Session Status Request and Trading Session Status with the same TradSesReqID
	if tradingSession.TradingSessionStatusRequest.Creator == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Trading Session Status", msg.Creator))
	}

	// check that the Trading Session Status Request is not rejected already
	if tradingSession.TradingSessionStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionStatusRequestIsRejected, "Trading Session: %s", tradingSession.TradingSessionStatusRequestReject)
	}

	// check that the Trading Session Status Request is not acknowledged already
	if tradingSession.TradingSessionStatus != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionStatusRequestIsAcknowledged, "Trading Session: %s", tradingSession.TradingSessionStatus)
	}

	// check that the mandatory field matches the value from Trading Session Status Request
	if tradingSession.TradingSessionStatusRequest.TradingSessionID != msg.TradingSessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionMismatchField, "TradingSessionID: %s", msg.TradingSessionID)
	}

	// check that the mandatory Trading Session Status field is not empty
	if _, err := strconv.ParseInt(fmt.Sprint(msg.TradSesStatus), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionEmptyField, "TradSesStatus: %v", msg.TradSesStatus)
	}

	// Trading Session Status
	tradingSessionStatus := types.TradingSession{
		SessionID:                   msg.SessionID,
		TradingSessionStatusRequest: tradingSession.TradingSessionStatusRequest,
		TradingSessionStatus: &types.TradingSessionStatus{
			TradSesReqID:           msg.TradSesReqID,
			TradingSessionID:       msg.TradingSessionID,
			TradSesStatus:          msg.TradSesStatus,
			TradSesStatusRejReason: msg.TradSesStatusRejReason,
			TradSesStartTime:       msg.TradSesStartTime,
			TradSesOpenTime:        msg.TradSesOpenTime,
			TradSesPreCloseTime:    msg.TradSesPreCloseTime,
			TradSesCloseTime:       msg.TradSesCloseTime,
			TradSesEndTime:         msg.TradSesEndTime,
			TotalVolumeTraded:      msg.TotalVolumeTraded,
			TradSesHighPx:          msg.TradSesHighPx,
			TradSesLowPx:           msg.TradSesLowPx,
			SecurityID:             msg.SecurityID,
			SecurityIDSource:       msg.SecurityIDSource,
			Symbol:                 msg.Symbol,
			SecurityExchange:       msg.SecurityExchange,
			MarketSegmentID:        msg.MarketSegmentID,
			MarketID:               msg.MarketID,
			Creator:                msg.Creator,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session Status is sent by the server(trading venue) and the server is the party that accepts the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "h" to indicate a Trading Session Status.
	// SenderCompID (49) contains the ID of the sender (trading venue).
	// TargetCompID (56) contains the ID of the target (client).
	// TradingSessionID (336) may be used to specify the trading session for which the status is requested.
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session Status
	tradingSessionStatus.TradingSessionStatus.Header = session.LogonAcceptor.Header
	tradingSessionStatus.TradingSessionStatus.Header.MsgType = "h"
	tradingSessionStatus.TradingSessionStatus.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionStatus.TradingSessionStatus.Trailer = session.LogonAcceptor.Trailer

	// set new Trading Session Status to store
	k.SetTradingSession(ctx, msg.TradSesReqID, tradingSessionStatus)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionStatusResponse{}, err
}

func (k msgServer) TradingSessionStatusRequestReject(goCtx context.Context, msg *types.MsgTradingSessionStatusRequestReject) (*types.MsgTradingSessionStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Trading Session Status Request Reject matches with the sessionID for Trading Session Status Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Trading Session Status Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Trading Session Status Creator: %s", msg.Creator)
	}

	// get Trading Session
	tradingSession, found := k.GetTradingSession(ctx, msg.RefSeqNum)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionIsNotFound, ": %s", tradingSession.TradingSessionStatusRequest)
	}

	// same account can not used for creating Trading Session Status Request and Trading Session Status Request Reject with the same TradSesReqID
	if tradingSession.TradingSessionStatusRequest.Creator == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Trading Session Status Request Reject", msg.Creator))
	}

	// check that the Trading Session Status Request is not rejected already
	if tradingSession.TradingSessionStatusRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionStatusRequestIsRejected, "Trading Session: %s", tradingSession.TradingSessionStatusRequestReject)
	}

	// check that the Trading Session Status Request is not acknowledged already
	if tradingSession.TradingSessionStatus != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionStatusRequestIsAcknowledged, "Trading Session: %s", tradingSession.TradingSessionStatus)
	}

	// check that the mandatory Trading Session Status Request Reject field is not empty
	if _, err := strconv.ParseInt(fmt.Sprint(msg.SessionRejectReason), 10, 64); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionEmptyField, "SessionRejectReason: %v", msg.SessionRejectReason)
	}

	// trading session status request reject
	tradingSessionStatusRequestReject := types.TradingSession{
		SessionID:                   msg.SessionID,
		TradingSessionStatusRequest: tradingSession.TradingSessionStatusRequest,
		TradingSessionStatus:        tradingSession.TradingSessionStatus,
		TradingSessionStatusRequestReject: &types.TradingSessionStatusRequestReject{
			RefSeqNum:           msg.RefSeqNum,
			RefMsgType:          msg.RefMsgType,
			SessionRejectReason: msg.SessionRejectReason,
			Text:                msg.Text,
			Creator:             msg.Creator,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session Status is sent by the server(trading venue) and the server is the party that accepts the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "j" to indicate a Trading Session Status Request Reject.
	// SenderCompID (49) contains the ID of the sender (trading venue).
	// TargetCompID (56) contains the ID of the target (client).
	// TradingSessionID (336) may be used to specify the trading session for which the status is requested.
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session Status Request Reject
	tradingSessionStatusRequestReject.TradingSessionStatusRequestReject.Header = session.LogonAcceptor.Header
	tradingSessionStatusRequestReject.TradingSessionStatusRequestReject.Header.MsgType = "j"
	tradingSessionStatusRequestReject.TradingSessionStatusRequestReject.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// for now copy trailer from session, it should be re-calculated
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionStatusRequestReject.TradingSessionStatusRequestReject.Trailer = session.LogonAcceptor.Trailer

	// set new Trading Session Status to store
	k.SetTradingSession(ctx, msg.RefSeqNum, tradingSessionStatusRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionStatusRequestRejectResponse{}, err
}
