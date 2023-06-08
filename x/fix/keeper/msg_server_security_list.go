package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// SecurityListRequest creates Security List Request
func (k msgServer) SecurityListRequest(goCtx context.Context, msg *types.MsgSecurityListRequest) (*types.MsgSecurityListRequestResponse, error) {
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

	// check that the parties involved in a session are the ones using the sessionID and are able to create Security List Request
	if session.InitiatorAddress != msg.Creator && session.AcceptorAddress != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Session Creator: %s", msg.Creator)
	}

	// check that the mandatory Security list Request field is not empty
	if msg.SecurityListRequestType == "" {
		return nil, sdkerrors.Wrapf(types.ErrSecurityListEmptyField, "SecurityListRequestType: %s", msg.SecurityListRequestType)
	}

	// security list request
	securityListRequest := types.SecurityList{
		SessionID: msg.SessionID,
		SecurityListRequest: &types.SecurityListRequest{
			SecurityReqID:           msg.SecurityReqID,
			SecurityListRequestType: msg.SecurityListRequestType,
			NoUnderlyings:           msg.NoUnderlyings,
			NoLegs:                  msg.NoLegs,
			Currency:                msg.Currency,
			Text:                    msg.Text,
			EncodedTextLen:          msg.EncodedTextLen,
			EncodedText:             msg.EncodedText,
			TradingSessionID:        msg.TradingSessionID,
			TradingSessionSubID:     msg.TradingSessionSubID,
			SubscriptionRequestType: msg.SubscriptionRequestType,
		},
	}

	// fetch Header from existing session
	// In the FIX Protocol, a Security List Request message can be sent by either the initiator or the acceptor of the FIX session.
	// Determine whether we are the initiator or acceptor
	var header *types.Header
	if session.InitiatorAddress == msg.Creator {
		header = session.LogonInitiator.Header
	} else {
		header = session.LogonAcceptor.Header
	}

	// set the header and make changes to the header
	// calculate and include all changes to the header
	// Message type, y is the message type for Security List Requesst
	// BodyLength should be calculated using the BodyLength function
	// set sending time to current time at creating Security List Request
	securityListRequest.SecurityListRequest.Header = header
	securityListRequest.SecurityListRequest.Header.MsgType = "y"
	securityListRequest.SecurityListRequest.Header.SendingTime = time.Now().UTC().Format("20060102-15:04:05.000")

	// fetch Trailer from existing session
	var trailer *types.Trailer
	if session.InitiatorAddress == msg.Creator {
		trailer = session.LogonInitiator.Trailer
	} else {
		trailer = session.LogonAcceptor.Trailer
	}

	// set the Trailer and make changes to the trailer
	// checksum in the trailer can be recalculated using CalculateChecksum function
	securityListRequest.SecurityListRequest.Trailer = trailer

	// set Security List Request to store
	k.SetSecurityList(ctx, msg.SecurityReqID, securityListRequest)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSecurityListRequestResponse{}, err
}

func (k msgServer) SecurityListResponse(goCtx context.Context, msg *types.MsgSecurityListResponse) (*types.MsgSecurityListResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityListResponseResponse{}, nil
}

func (k msgServer) SecurityListRequestReject(goCtx context.Context, msg *types.MsgSecurityListRequestReject) (*types.MsgSecurityListRequestRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSecurityListRequestRejectResponse{}, nil
}
