package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

// LogonInitiator initiates the session
func (k msgServer) LogonInitiator(goCtx context.Context, msg *types.MsgLogonInitiator) (*types.MsgLogonInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.InitiatorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.InitiatorAddress)
	}

	// address will be used as senderCompID and targetCompID for both parties
	// get address from registered accounts
	senderCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.SenderCompID)
	if senderCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "senderCompID: %s", msg.LogonInitiator.Header.SenderCompID)
	}
	if senderCompID.Address != msg.InitiatorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "senderCompID: %s", msg.InitiatorAddress)
	}

	targetCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.TargetCompID)
	if targetCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "targetCompID: %s", msg.LogonInitiator.Header.TargetCompID)
	}

	// same address can not be used for intiating and accepting in the same FIX session
	if senderCompID.Address == targetCompID.Address {
		return nil, sdkerrors.Wrapf(types.ErrSessionSameAddress, "Address: %s", msg.LogonInitiator.Header.TargetCompID)
	}

	// check for if this session Name exists already
	_, found := k.GetSessions(ctx, msg.SessionID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrSessionIDFound, "SessionID: %s", msg.SessionID)
	}

	// set the standard header
	header := types.NewHeader(msg.LogonInitiator.Header.BodyLength, msg.LogonInitiator.Header.MsgType, senderCompID.Address, targetCompID.Address, msg.LogonInitiator.Header.MsgSeqNum, msg.LogonInitiator.Header.SendingTime)

	// set the FIX Version
	header.BeginString = msg.FIXVersionByInitiator()

	// set the standard trailer
	trailer := types.NewTrailer(msg.LogonInitiator.Trailer.CheckSum)

	// set the logon initiator message
	logonInitiator := types.NewLogonInitiator(header, msg.LogonInitiator.EncryptMethod, msg.LogonInitiator.HeartBtInt, trailer)

	newInitiatorSession := types.Sessions{
		SessionID:        msg.SessionID,
		LogonInitiator:   &logonInitiator,
		Status:           "logon-request",
		IsAccepted:       false,
	}

	// set new Initiator logon session to store
	k.SetSessions(ctx, msg.SessionID, newInitiatorSession)

	return &types.MsgLogonInitiatorResponse{}, nil
}

// LogonAcceptor accepts the session
func (k msgServer) LogonAcceptor(goCtx context.Context, msg *types.MsgLogonAcceptor) (*types.MsgLogonAcceptorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.AcceptorAddress)
	}

	// check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	// checks that the provided session name matches with the existing session name in store
	if session.SessionID != msg.SessionID {
		return nil, sdkerrors.Wrapf(types.ErrWrongSession, "SessionID: %s", msg.SessionID)
	}

	// checks that address provided matches with the address in the session
	// at logon acceptor, targetCompID becomes senderCompID and senderCompID becomes targetCompID
	if session.LogonInitiator.Header.TargetCompID != msg.LogonAcceptor.Header.SenderCompID {
		return nil, sdkerrors.Wrapf(types.ErrIncorrectAddress, "SessionID: %s", session.LogonInitiator.Header.TargetCompID)
	}
	if session.LogonInitiator.Header.SenderCompID != msg.LogonAcceptor.Header.TargetCompID {
		return nil, sdkerrors.Wrapf(types.ErrIncorrectAddress, "SessionID: %s", session.LogonInitiator.Header.TargetCompID)
	}

	// get address from GetAccount to have access to the account creator
	senderCompID := k.GetAccount(ctx, msg.LogonAcceptor.Header.SenderCompID)
	if senderCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrAccountIsEmpty, "Account: %s", msg.LogonAcceptor.Header.SenderCompID)
	}
	if senderCompID.Address != msg.AcceptorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.AcceptorAddress)
	}

	// set the standard header
	header := types.NewHeader(session.LogonInitiator.Header.BodyLength, msg.LogonAcceptor.Header.MsgType, msg.LogonAcceptor.Header.SenderCompID, msg.LogonAcceptor.Header.TargetCompID, session.LogonInitiator.Header.MsgSeqNum, msg.LogonAcceptor.Header.SendingTime)

	// set the FIX Version
	header.BeginString = msg.FIXVersionByAcceptor()

	// set the standard trailer
	trailer := types.NewTrailer(msg.LogonAcceptor.Trailer.CheckSum)

	// set the logon initiator message
	LogonAcceptor := types.NewLogonAcceptor(header, msg.LogonAcceptor.EncryptMethod, msg.LogonAcceptor.HeartBtInt, trailer)

	// TODO
	// prevent Txs from being sent once IsAccepted are set to true and check the status
	// if session.status != "logon-request" || session.status != "rejected"  && session.IsAccepted == false {
	//
	// }

	newAcceptorSession := types.Sessions{
		SessionID:        session.SessionID,
		LogonInitiator:   session.LogonInitiator,
		LogonAcceptor:    &LogonAcceptor,
		Status:           "loggedIn",
		IsAccepted:       true,
	}

	// set new Acceptor logon session to store
	k.SetSessions(ctx, msg.SessionID, newAcceptorSession)

	return &types.MsgLogonAcceptorResponse{}, nil
}

// LogonReject enables an Acceptor to reject the logon session
func (k msgServer) LogonReject(goCtx context.Context, msg *types.MsgLogonReject) (*types.MsgLogonRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.AcceptorAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.AcceptorAddress)
	}

	// check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID %s", msg.SessionID)
	}

	if session.Status != "logon-request" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d Logon status is not in request state", &session.Status))
	}

	if session.LogonInitiator.Header.TargetCompID != msg.Header.SenderCompID {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong SenderCompID", msg.Header.SenderCompID))
	}
	if session.LogonAcceptor.Header.SenderCompID == msg.AcceptorAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong Account Address", msg.AcceptorAddress))
	}

	// pass TargetCompID through GetAccount to ensure only the owner can reject session
	getAcc := k.GetAccount(ctx, session.LogonInitiator.Header.TargetCompID)
	if getAcc.Address != msg.AcceptorAddress {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s Wrong Account Address", msg.AcceptorAddress))
	}

	sessionReject := types.SessionReject{
		SessionID:       msg.SessionID,
		Header:          msg.Header,
		Text:            msg.Text,
		Trailer:         msg.Trailer,
	}

	// set Header and Trailer to session Reject
	sessionReject.Header = session.LogonInitiator.Header
	sessionReject.Trailer = session.LogonInitiator.Trailer

	// set msgType to 3, [msgType 35 = 3] for sesssion rejection
	sessionReject.Header.MsgType = "3"

	// set session status to rejected
	session.Status = "rejected"

	// set rejected session to sore
	k.SetSessions(ctx, msg.SessionID, session)
	k.SetSessionReject(ctx, msg.SessionID, sessionReject)

	return &types.MsgLogonRejectResponse{}, nil
}

// TerminateLogon terminates an initiated session by the initiator
func (k msgServer) TerminateLogon(goCtx context.Context, msg *types.MsgTerminateLogon) (*types.MsgTerminateLogonResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the session exits with the addresss
	session, found := k.GetSessions(ctx, msg.SessionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "SessionID: %s", msg.SessionID)
	}

	if session.LogonInitiator.Header.SenderCompID != msg.InitiatorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.InitiatorAddress)
	}

	// check that session is not accepted yet
	if session.Status != "logon-request" {
		return nil, sdkerrors.Wrapf(types.ErrSessionIsAccepted, "Account: %s", msg.SessionID)
	}

	// remove session from store
	k.RemoveSessions(ctx, msg.SessionID)

	return &types.MsgTerminateLogonResponse{}, nil
}
