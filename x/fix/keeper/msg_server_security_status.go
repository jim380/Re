package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// SecurityStatusRequest creates Security Status Request
func (k msgServer) SecurityStatusRequest(goCtx context.Context, msg *types.MsgSecurityStatusRequest) (*types.MsgSecurityStatusRequestResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Security Status Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Security Status Request field is not empty
	if msg.SubscriptionRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityStatusEmptyField, "SubscriptionRequestType: %s", msg.SubscriptionRequestType)
	}

	// Must be unique, or the ID of previous Security Status Request (e) to disable if SubscriptionRequestType (263) = Disable previous Snapshot + Updates Request (2)
	if msg.SecurityStatusReqID != "" {
		// check that the provided SecurityStatusReqID exists from the Security Status
		securityStatus, found := k.GetSecurityStatus(ctx, msg.SecurityStatusReqID)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrSecurityStatusIsNotFound, ": %s", securityStatus.SecurityStatusRequest)
		}
	}

	securityStatusRequest := types.SecurityStatus{
		SessionID: msg.SessionID,
		SecurityStatusRequest: &types.SecurityStatusRequest{
			SecurityStatusReqID:     msg.SecurityStatusReqID,
			Instrument:              msg.Instrument,
			NoUnderlyings:           msg.NoUnderlyings,
			UnderlyingInstrument:    msg.UnderlyingInstrument,
			NoLegs:                  msg.NoLegs,
			InstrumentLeg:           msg.InstrumentLeg,
			SubscriptionRequestType: msg.SubscriptionRequestType,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
			Creator:                 msg.Creator,
		},
	}

	// fetch Header from existing session
	// Determine whether it is the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, e is the message type for Security Status Request
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security Status Request
	securityStatusRequest.SecurityStatusRequest.Header = header
	securityStatusRequest.SecurityStatusRequest.Header.MsgType = "x"
	securityStatusRequest.SecurityStatusRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityStatusRequest.SecurityStatusRequest.Trailer = trailer

	// set Security Status Request to store
	k.SetSecurityStatus(ctx, msg.SecurityStatusReqID, securityStatusRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityStatusRequestResponse{}, err
}

func (k msgServer) SecurityStatusResponse(goCtx context.Context, msg *types.MsgSecurityStatusResponse) (*types.MsgSecurityStatusResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusResponseResponse{}, nil
}

func (k msgServer) SecurityStatusRequestReject(goCtx context.Context, msg *types.MsgSecurityStatusRequestReject) (*types.MsgSecurityStatusRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityStatusRequestRejectResponse{}, nil
}
