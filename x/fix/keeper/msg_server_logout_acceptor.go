package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// LogoutAcceptor accepts to logout from the session
func (k msgServer) LogoutAcceptor(goCtx context.Context, msg *types.MsgLogoutAcceptor) (*types.MsgLogoutAcceptorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	//get logout session, if not found, then logout is not initiated
	initiatedLogoutSesssion, found := k.GetSessionLogout(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Logout was not initiated", msg.SessionID))
	}

	/*
		if initiatedLogoutSesssion.InitiatorAddress == msg.AcceptorAddress {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong account address", msg.AcceptorAddress))
		}
	*/

	//TODO
	//Include checks
	//when logout is accpted, session name should be removed from the store
	//check that the address accepting the logout participated in the session

	sessionLogoutAcceptor := types.SessionLogout{
		AcceptorAddress:        msg.AcceptorAddress,
		SessionID:              msg.SessionID,
		SessionLogoutAcceptor:  msg.SessionLogoutAcceptor,
		InitiatorAddress:       initiatedLogoutSesssion.InitiatorAddress,
		SessionLogoutInitiator: initiatedLogoutSesssion.SessionLogoutInitiator,
	}

	//set the session logout to store
	sessionLogoutAcceptor.AcceptorAddress = session.AcceptorAddress
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header = session.LogonAcceptor.Header
	sessionLogoutAcceptor.SessionLogoutAcceptor.Trailer = session.LogonAcceptor.Trailer
	sessionLogoutAcceptor.SessionLogoutAcceptor.Text = msg.SessionLogoutAcceptor.Text
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header.MsgType = "5"

	k.SetSessionLogout(ctx, msg.SessionID, sessionLogoutAcceptor)

	//remove session from store
	k.RemoveSessions(ctx, msg.SessionID)

	return &types.MsgLogoutAcceptorResponse{}, nil
}
