package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// LogoutInitiator requests for logout from the session
func (k msgServer) LogoutInitiator(goCtx context.Context, msg *types.MsgLogoutInitiator) (*types.MsgLogoutInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.InitiatorAddress)
	}

	// check for if this sessionID exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that logon is established between both parties and that logon status equals to "loggedIn"
	if session.Status != constants.LoggedInStatus {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsNotLoggedIn, "Status of Session: %s", session.Status)
	}

	// check that the address initiating a session logout belongs to the session
	if session.LogonInitiator.Header.SenderCompID != msg.InitiatorAddress && session.LogonAcceptor.Header.SenderCompID != msg.InitiatorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Logout Initiator: %s", msg.InitiatorAddress)
	}

	// check for if this sessionLogout exists already
	initiatedLogoutSesssion, found := k.GetSessionLogout(ctx, msg.SessionID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrSessionLogoutIsInitiated, "Session Logout: %s", initiatedLogoutSesssion.SessionID)
	}
	sessionLogoutInitiator := types.SessionLogout{
		SessionID:              msg.SessionID,
		SessionLogoutInitiator: msg.SessionLogoutInitiator,
	}

	sessionLogoutInitiator.SessionLogoutInitiator.Header = session.LogonInitiator.Header
	sessionLogoutInitiator.SessionLogoutInitiator.Header.SenderCompID = session.LogonInitiator.Header.SenderCompID
	sessionLogoutInitiator.SessionLogoutInitiator.Trailer = session.LogonInitiator.Trailer
	sessionLogoutInitiator.SessionLogoutInitiator.Header.MsgType = "5"
	sessionLogoutInitiator.SessionLogoutInitiator.Text = msg.SessionLogoutInitiator.Text

	// set session logout to store
	k.SetSessionLogout(ctx, msg.SessionID, sessionLogoutInitiator)

	return &types.MsgLogoutInitiatorResponse{}, nil
}

// LogoutAcceptor accepts to logout from the session
func (k msgServer) LogoutAcceptor(goCtx context.Context, msg *types.MsgLogoutAcceptor) (*types.MsgLogoutAcceptorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.AcceptorAddress)
	}

	// check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionID)
	}

	// check that the address accepting the logout participated in the session
	if session.LogonInitiator.Header.SenderCompID != msg.AcceptorAddress && session.LogonAcceptor.Header.SenderCompID != msg.AcceptorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Logout Acceptor: %s", msg.AcceptorAddress)
	}

	// get logout session, if not found, then logout is not initiated
	initiatedLogoutSesssion, found := k.GetSessionLogout(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSessionLogoutIsNotInitiated, "Session Logout: %s", initiatedLogoutSesssion.SessionID)
	}

	// check that the session logout has not been accepted
	if initiatedLogoutSesssion.SessionLogoutAcceptor != nil {
		return nil, sdkerrors.Wrapf(types.ErrSessionLogoutIsAccepted, "Session Logout: %s", initiatedLogoutSesssion.SessionLogoutAcceptor)
	}

	sessionLogoutAcceptor := types.SessionLogout{
		SessionID:              msg.SessionID,
		SessionLogoutAcceptor:  msg.SessionLogoutAcceptor,
		SessionLogoutInitiator: initiatedLogoutSesssion.SessionLogoutInitiator,
	}

	// set the session logout to store
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header = session.LogonAcceptor.Header
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header.SenderCompID = session.LogonAcceptor.Header.SenderCompID
	sessionLogoutAcceptor.SessionLogoutAcceptor.Trailer = session.LogonAcceptor.Trailer
	sessionLogoutAcceptor.SessionLogoutAcceptor.Text = msg.SessionLogoutAcceptor.Text
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header.MsgType = "5"

	// set session logout acceptor to store
	k.SetSessionLogout(ctx, msg.SessionID, sessionLogoutAcceptor)

	// remove session from store
	k.RemoveSessions(ctx, msg.SessionID)

	return &types.MsgLogoutAcceptorResponse{}, nil
}
