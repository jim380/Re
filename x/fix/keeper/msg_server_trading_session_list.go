package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) TradingSessionListRequest(goCtx context.Context, msg *types.MsgTradingSessionListRequest) (*types.MsgTradingSessionListRequestResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create trading session list request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Trading Session list Request field is not empty
	if msg.SubscriptionRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrTradingSessionListEmptyField, "SubscriptionRequestType: %s", msg.SubscriptionRequestType)
	}

	// when TradSesReqID is from the previous Trading Session Status Request (g)
	if msg.TradSesReqID != "" {
		// check that the provided TradSesReqID exists from the Trading Session
		tradingSession, found := k.GetTradingSession(ctx, msg.TradSesReqID)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrTradingSessionIsNotFound, ": %s", tradingSession.TradingSessionStatusRequest)
		}
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
	// set sending time to current time at creating Trading Session Status Request
	tradingSessionListRequest.TradingSessionListRequestReject.Header = session.LogonInitiator.Header
	tradingSessionListRequest.TradingSessionListRequestReject.Header.MsgType = "BI"
	tradingSessionListRequest.TradingSessionListRequestReject.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	// TODO
	// checksum in the trailer can be recalculated using CalculateChecksum function
	tradingSessionListRequest.TradingSessionListRequestReject.Trailer = session.LogonInitiator.Trailer

	// set Trade Session List Request to store
	k.SetTradingSessionList(ctx, msg.TradSesReqID, tradingSessionListRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionListRequestResponse{}, err
}

func (k msgServer) TradingSessionListResponse(goCtx context.Context, msg *types.MsgTradingSessionListResponse) (*types.MsgTradingSessionListResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionListResponseResponse{}, nil
}

func (k msgServer) TradingSessionListRequestReject(goCtx context.Context, msg *types.MsgTradingSessionListRequestReject) (*types.MsgTradingSessionListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionListRequestRejectResponse{}, nil
}
