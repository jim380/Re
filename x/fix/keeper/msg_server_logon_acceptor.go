package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) LogonAcceptor(goCtx context.Context, msg *types.MsgLogonAcceptor) (*types.MsgLogonAcceptorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrEmptySession, "Session Name: %s", msg.SessionName)
	}

	//checks that the provided session name matches with the existing session name in store
	if session.SessionName != msg.SessionName {
		return nil, sdkerrors.Wrapf(types.ErrWrongSession, "Session Name: %s", msg.SessionName)
	}

	//checks that DID provided matches with the DID in the session
	//at logon acceptor, targetCompID becomes senderCompID and senderCompID becomes targetCompID
	if session.LogonInitiator.Header.TargetCompID != msg.LogonAcceptor.Header.SenderCompID {
		return nil, sdkerrors.Wrapf(types.ErrIncorrectDID, "Session Name: %s", session.LogonInitiator.Header.TargetCompID, msg.LogonAcceptor.Header.SenderCompID)
	}
	if session.LogonInitiator.Header.SenderCompID != msg.LogonAcceptor.Header.TargetCompID {
		return nil, sdkerrors.Wrapf(types.ErrIncorrectDID, "Session Name: %s", session.LogonInitiator.Header.TargetCompID, msg.LogonAcceptor.Header.SenderCompID)
	}

	//get DID from GetAccount to have access to the account creator
	senderCompID := k.GetAccount(ctx, msg.LogonAcceptor.Header.SenderCompID)
	if senderCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.LogonAcceptor.Header.SenderCompID)
	}
	if senderCompID.Creator != msg.AcceptorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.AcceptorAddress)
	}

	//set the standard header
	header := types.NewHeader(session.LogonInitiator.Header.BodyLength, msg.LogonAcceptor.Header.MsgType, msg.LogonAcceptor.Header.SenderCompID, msg.LogonAcceptor.Header.TargetCompID, session.LogonInitiator.Header.MsgSeqNum, msg.LogonAcceptor.Header.SendingTime)

	//set the FIX Version
	header.BeginString = msg.FIXVersionByAcceptor()

	//set the standard trailer
	trailer := types.NewTrailer(msg.LogonAcceptor.Trailer.CheckSum)

	// set the logon initiator message
	LogonAcceptor := types.NewLogonAcceptor(header, msg.LogonAcceptor.EncryptMethod, msg.LogonAcceptor.HeartBtInt, trailer)

	//TODO
	//prevent Txs from being sent once IsloggedIn and IsAccepted are set to true
	//if session.IsLoggedIn == false && session.IsAccepted == false {
	//	return nil, sdkerrors.Wrapf(types.ErrSessionIsAccepted, "SessionIsAccepted: %s", session.IsAccepted)
	//}

	var newAcceptorSession = types.Sessions{
		SessionName:      session.SessionName,
		LogonInitiator:   session.LogonInitiator,
		LogonAcceptor:    &LogonAcceptor,
		Status:           "loggedIn",
		IsAccepted:       true,
		InitiatorAddress: session.InitiatorAddress,
		AcceptorAddress:  msg.AcceptorAddress,
	}

	//set new Acceptor logon session to store
	k.SetSessions(ctx, msg.SessionName, newAcceptorSession)

	return &types.MsgLogonAcceptorResponse{}, nil
}
