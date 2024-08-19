package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

// LogoutInitiator requests for logout from the session
func (k msgServer) LogoutInitiator(goCtx context.Context, msg *types.MsgLogoutInitiator) (*types.MsgLogoutInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.InitiatorAddress)
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
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.AcceptorAddress)
	}

	// get logout session, if not found, then logout is not initiated
	initiatedLogoutSesssion, found := k.GetSessionLogout(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrSessionLogoutIsNotInitiated, "Session Logout: %s", initiatedLogoutSesssion.SessionID)
	}

	// check that the address accepting the logout participated in the session
	if initiatedLogoutSesssion.SessionLogoutInitiator.Header.SenderCompID != msg.AcceptorAddress && initiatedLogoutSesssion.SessionLogoutInitiator.Header.TargetCompID != msg.AcceptorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Logout Acceptor: %s", msg.AcceptorAddress)
	}

	// check that the session logout has not been accepted
	if initiatedLogoutSesssion.SessionLogoutAcceptor != nil {
		return nil, sdkerrors.Wrapf(types.ErrSessionLogoutIsAccepted, "Session Logout: %s", initiatedLogoutSesssion.SessionLogoutAcceptor)
	}

	// check that the initiator and the acceptor address are not the same
	if initiatedLogoutSesssion.SessionLogoutInitiator.Header.SenderCompID == msg.AcceptorAddress {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrKeyNotFound, fmt.Sprintf("key %s Wrong Account Address", msg.AcceptorAddress))
	}

	sessionLogoutAcceptor := types.SessionLogout{
		SessionID:              msg.SessionID,
		SessionLogoutAcceptor:  msg.SessionLogoutAcceptor,
		SessionLogoutInitiator: initiatedLogoutSesssion.SessionLogoutInitiator,
	}

	// set header from the existing session logout initiator
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header = initiatedLogoutSesssion.SessionLogoutInitiator.Header

	// create a copy of the Header
	newHeader := new(types.Header)
	*newHeader = *sessionLogoutAcceptor.SessionLogoutAcceptor.Header

	// set bodyLength
	// TODO
	// Recalculate the bodyLength in the header
	newHeader.BodyLength = initiatedLogoutSesssion.SessionLogoutInitiator.Header.BodyLength

	// set msgSeqNum
	newHeader.MsgSeqNum = initiatedLogoutSesssion.SessionLogoutInitiator.Header.MsgSeqNum

	// set the msgType to session logout acceptor
	newHeader.MsgType = "5"

	// switch senderCompID and targetCompID between Session Logout Acceptor and Session Logout Initiator
	// set senderCompID of Session Logout Acceptor to the targetCompID of Session Logout Initiator in the header
	newHeader.SenderCompID = initiatedLogoutSesssion.SessionLogoutInitiator.Header.TargetCompID

	// set targetCompID of Session Logout Acceptor to the senderCompID of Session Logout Initiator in the header
	newHeader.TargetCompID = initiatedLogoutSesssion.SessionLogoutInitiator.Header.SenderCompID

	// set sending time
	newHeader.SendingTime = constants.SendingTime

	// pass all the edited values to the newHeader
	sessionLogoutAcceptor.SessionLogoutAcceptor.Header = newHeader

	// set Trailer from the existing session logout initiator
	sessionLogoutAcceptor.SessionLogoutAcceptor.Trailer = initiatedLogoutSesssion.SessionLogoutInitiator.Trailer

	// set session logout acceptor to store
	k.SetSessionLogout(ctx, msg.SessionID, sessionLogoutAcceptor)

	// remove session from store
	k.RemoveSessions(ctx, msg.SessionID)

	return &types.MsgLogoutAcceptorResponse{}, nil
}
