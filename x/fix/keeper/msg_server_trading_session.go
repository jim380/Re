package keeper

import (
	"context"
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

	// check for if the provided session ID existss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
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

	// set Trade Capture Report to store
	k.SetTradingSession(ctx, msg.TradSesReqID, tradingSessionStatusRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgTradingSessionStatusRequestResponse{}, err
}

func (k msgServer) TradingSessionStatus(goCtx context.Context, msg *types.MsgTradingSessionStatus) (*types.MsgTradingSessionStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionStatusResponse{}, nil
}

func (k msgServer) TradingSessionStatusRequestReject(goCtx context.Context, msg *types.MsgTradingSessionStatusRequestReject) (*types.MsgTradingSessionStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTradingSessionStatusRequestRejectResponse{}, nil
}
