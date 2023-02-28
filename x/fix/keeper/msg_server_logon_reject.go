package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// LogonReject enables an Acceptor to reject the logon session
func (k msgServer) LogonReject(goCtx context.Context, msg *types.MsgLogonReject) (*types.MsgLogonRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionName)
	}

	if session.Status != "logon-request" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d Logon status is not in request state", &session.Status))
	}

	/*
		//TODO
		if session.LogonInitiator.Header.TargetCompID != msg.Header.SenderCompID {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong SenderCompID", msg.Header.SenderCompID))
		}
		if session.InitiatorAddress == msg.AcceptorAddress {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong Account Address", msg.AcceptorAddress))
		}
	*/

	// pass TargetCompID through GetAccount to ensure only the owner can reject session
	getAcc := k.GetAccount(ctx, session.LogonInitiator.Header.TargetCompID)
	if getAcc.Creator != msg.AcceptorAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong Account Address", msg.AcceptorAddress))
	}

	var sessionReject = types.SessionReject{
		AcceptorAddress: msg.AcceptorAddress,
		SessionName:     msg.SessionName,
		Header:          msg.Header,
		Text:            msg.Text,
		Trailer:         msg.Trailer,
	}

	//set Header and Trailer to session Reject
	sessionReject.Header = session.LogonInitiator.Header
	sessionReject.Trailer = session.LogonInitiator.Trailer

	//set msgType to 3, [msgType 35 = 3] for sesssion rejection
	sessionReject.Header.MsgType = "3"

	//set session status to rejected
	session.Status = "rejected"

	//set rejected session to sore
	k.SetSessions(ctx, msg.SessionName, session)
	k.SetSessionReject(ctx, msg.SessionName, sessionReject)

	return &types.MsgLogonRejectResponse{}, nil
}
