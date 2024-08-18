package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/helpers"
	"github.com/jim380/Re/x/fix/types"
)

// TradingSessionListRequest creates Trading Session List Request
func (k msgServer) TradingSessionListRequest(goCtx context.Context, msg *types.MsgTradingSessionListRequest) (*types.MsgTradingSessionListRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", msg.SessionID)
	}

	// check that the parties involved in a session are the ones using the sessionID and are able to create trading session list request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Trading Session list Request field is not empty
	if msg.SubscriptionRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "SubscriptionRequestType: %s", msg.SubscriptionRequestType)
	}

	// when TradSesReqID is from the previous Trading Session Status Request (g)
	if msg.TradSesReqID == "" {
		// Generate a random TradSesReqID using helpers.GenerateRandomString
		tradSesReqID, _ := helpers.GenerateRandomString(constants.TradSesReqID)
		msg.TradSesReqID = tradSesReqID
	} else {
		// check that the provided TradSesReqID exists from the Trading Session
		tradingSession, found := k.GetTradingSession(ctx, msg.TradSesReqID)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrTradingSessionIsNotFound, ": %s", tradingSession.TradingSessionStatusRequest)
		}
		msg.TradSesReqID = tradingSession.TradingSessionStatusRequest.TradSesReqID
	}

	tradingSessionListRequest := types.TradingSessionList{
		SessionID: msg.SessionID,
		TradingSessionListRequest: &types.TradingSessionListRequest{
			TradSesReqID:            msg.TradSesReqID,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
			SecurityExchange:        msg.SecurityExchange,
			TradSesMethod:           msg.TradSesMethod,
			TradSesMode:             msg.TradSesMode,
			SubscriptionRequestType: msg.SubscriptionRequestType,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session List Request is sent by the client and the client is the party that initiates the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "BI" to indicate a Trading Session List Request.
	// SenderCompID (49) contains the ID of the sender (client).
	// TargetCompID (56) contains the ID of the target (trading venue).
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session List Request
	tradingSessionListRequest.TradingSessionListRequest.Header = session.LogonInitiator.Header
	tradingSessionListRequest.TradingSessionListRequest.Header.MsgType = "BI"
	tradingSessionListRequest.TradingSessionListRequest.Header.SendingTime = constants.SendingTime

	// fetch Trailer from existing session
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionListRequest.TradingSessionListRequest.Trailer = session.LogonInitiator.Trailer

	// set Trade Session List Request to store
	k.SetTradingSessionList(ctx, msg.TradSesReqID, tradingSessionListRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionListRequestResponse{}, err
}

// TradingSessionListResponse creates Trading Session List Response
// TradingSessionListResponse is used for responding to TradingSessionListRequest
func (k msgServer) TradingSessionListResponse(goCtx context.Context, msg *types.MsgTradingSessionListResponse) (*types.MsgTradingSessionListResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Trading Session List Response matches with the sessionID for Trading Session List Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Trading Session List Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Trading Session List Response Creator: %s", msg.Creator)
	}

	// get Trading Session List
	tradingSessionList, found := k.GetTradingSessionList(ctx, msg.TradSesReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListIsNotFound, ": %s", tradingSessionList.TradingSessionListRequest)
	}

	// same account can not used for creating Trading Session List Request and Trading Session List Response with the same TradSesReqID
	if tradingSessionList.TradingSessionListRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Trading Session List Response", msg.Creator))
	}

	// check that the Trading Session List Request is not rejected already
	if tradingSessionList.TradingSessionListRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListRequestIsRejected, "Trading Session List: %s", tradingSessionList.TradingSessionListRequestReject)
	}

	// check that the Trading Session List Request is not acknowledged already
	if tradingSessionList.TradingSessionListResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListRequestIsAcknowledged, "Trading Session List: %s", tradingSessionList.TradingSessionListResponse)
	}

	// check that the mandatory Trading Session list Response field are not empty
	if msg.NoTradingSessions == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "NoTradingSessions: %s", msg.NoTradingSessions)
	}
	if msg.TradingSessionID == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "TradingSessionID: %s", msg.TradingSessionID)
	}
	if msg.TradSesStatus == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "TradSesStatus: %s", msg.TradSesStatus)
	}

	tradingSessionListResponse := types.TradingSessionList{
		SessionID:                 msg.SessionID,
		TradingSessionListRequest: tradingSessionList.TradingSessionListRequest,
		TradingSessionListResponse: &types.TradingSessionListResponse{
			TradSesReqID:           msg.TradSesReqID,
			NoTradingSessions:      msg.NoTradingSessions,
			TradingSessionID:       msg.TradingSessionID,
			TradingSessionSubID:    msg.TradingSessionSubID,
			SecurityExchange:       msg.SecurityExchange,
			TradSesMethod:          msg.TradSesMethod,
			TradSesMode:            msg.TradSesMode,
			UnsolicitedIndicator:   msg.UnsolicitedIndicator,
			TradSesStatus:          msg.TradSesStatus,
			TradSesStatusRejReason: msg.TradSesStatusRejReason,
			TradSesStartTime:       msg.TradSesStartTime,
			TradSesOpenTime:        msg.TradSesOpenTime,
			TradSesPreCloseTime:    msg.TradSesPreCloseTime,
			TradSesCloseTime:       msg.TradSesCloseTime,
			TradSesEndTime:         msg.TradSesEndTime,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session List Response is sent by the server(trading venue) and the server is the party that accepts the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "BJ" to indicate a Trading Session List Response.
	// SenderCompID (49) contains the ID of the sender (trading venue).
	// TargetCompID (56) contains the ID of the target (client).
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session List Response
	tradingSessionListResponse.TradingSessionListResponse.Header = session.LogonAcceptor.Header
	tradingSessionListResponse.TradingSessionListResponse.Header.MsgType = "BJ"
	tradingSessionListResponse.TradingSessionListResponse.Header.SendingTime = constants.SendingTime

	// fetch Trailer from existing session
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionListResponse.TradingSessionListResponse.Trailer = session.LogonAcceptor.Trailer

	// set new Trading Session List Response to store
	k.SetTradingSessionList(ctx, msg.TradSesReqID, tradingSessionListResponse)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionListResponseResponse{}, err
}

// TradingSessionListRequestReject creates Trading Session List Request Reject
func (k msgServer) TradingSessionListRequestReject(goCtx context.Context, msg *types.MsgTradingSessionListRequestReject) (*types.MsgTradingSessionListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}

	// check for if the provided session ID exists
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// check that the sessionID provided by the creator of Trading Session List Request Reject matches with the sessionID for Trading Session List Request
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListSession, "SessionID: %s", msg.SessionID)
	}

	// check that the user responding is the recipient of the Trading Session List Request
	if session.LogonInitiator.Header.SenderCompID != msg.Creator && session.LogonAcceptor.Header.SenderCompID != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Trading Session List Request Reject Creator: %s", msg.Creator)
	}

	// get Trading Session List
	tradingSessionList, found := k.GetTradingSessionList(ctx, msg.TradSesReqID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListIsNotFound, ": %s", tradingSessionList.TradingSessionListRequest)
	}

	// same account can not used for creating Trading Session List Request and Trading Session List Request Reject with the same TradSesReqID
	if tradingSessionList.TradingSessionListRequest.Header.SenderCompID == msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s This account can not be used to create Trading Session List Request Reject", msg.Creator))
	}

	// check that the Trading Session List Request is not rejected already
	if tradingSessionList.TradingSessionListRequestReject != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListRequestIsRejected, "Trading Session List: %s", tradingSessionList.TradingSessionListRequestReject)
	}

	// check that the Trading Session List Request is not acknowledged already
	if tradingSessionList.TradingSessionListResponse != nil {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListRequestIsAcknowledged, "Trading Session List: %s", tradingSessionList.TradingSessionListResponse)
	}

	// check that the mandatory Trading Session Request Reject fields are not empty
	if msg.TradSesStatus == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "TradSesStatus: %s", msg.TradSesStatus)
	}
	if msg.TradSesStatusRejReason == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "TradSesStatusRejReason: %s", msg.TradSesStatusRejReason)
	}

	tradingSessionListRequestReject := types.TradingSessionList{
		SessionID:                  msg.SessionID,
		TradingSessionListRequest:  tradingSessionList.TradingSessionListRequest,
		TradingSessionListResponse: tradingSessionList.TradingSessionListResponse,
		TradingSessionListRequestReject: &types.TradingSessionListRequestReject{
			TradSesReqID:           msg.TradSesReqID,
			TradSesStatus:          msg.TradSesStatus,
			TradSesStatusRejReason: msg.TradSesStatusRejReason,
			Text:                   msg.Text,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, Trading Session List Request Reject is sent by the server(trading venue) and the server is the party that accepts the logon request
	// set the header and make changes to the header
	// calculate and include all changes to the header
	// MsgType (35) is set to "BK" to indicate a Trading Session List Request Reject.
	// SenderCompID (49) contains the ID of the sender (trading venue).
	// TargetCompID (56) contains the ID of the target (client).
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Trading Session List Request Reject
	tradingSessionListRequestReject.TradingSessionListRequestReject.Header = session.LogonAcceptor.Header
	tradingSessionListRequestReject.TradingSessionListRequestReject.Header.MsgType = "BK"
	tradingSessionListRequestReject.TradingSessionListRequestReject.Header.SendingTime = constants.SendingTime

	// fetch Trailer from existing session
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionListRequestReject.TradingSessionListRequestReject.Trailer = session.LogonAcceptor.Trailer

	// set new Trading Session List Request Reject to store
	k.SetTradingSessionList(ctx, msg.TradSesReqID, tradingSessionListRequestReject)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionListRequestRejectResponse{}, err
}
