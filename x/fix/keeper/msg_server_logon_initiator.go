package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

func (k msgServer) LogonInitiator(goCtx context.Context, msg *types.MsgLogonInitiator) (*types.MsgLogonInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// DID will be used as senderCompID and targetCompID for both parties
	// get DID from registered accounts
	senderCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.SenderCompID)
	if senderCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.LogonInitiator.Header.SenderCompID)
	}
	if senderCompID.Creator != msg.InitiatorAddress {
		return nil, sdkerrors.Wrapf(types.ErrNotAccountCreator, "Account: %s", msg.InitiatorAddress)
	}

	targetCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.TargetCompID)
	if targetCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.LogonInitiator.Header.TargetCompID)
	}

	//same DID can not be used for intiating and accepting in the same FIX session
	if senderCompID.Did == targetCompID.Did {
		return nil, sdkerrors.Wrapf(types.ErrSessionSameDID, "DID: %s", msg.LogonInitiator.Header.TargetCompID, msg.LogonInitiator.Header.SenderCompID)
	}

	//check for if this session Name exists already
	_, found := k.GetSessions(ctx, msg.SessionName)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrSessionNameFound, "Session Name: %s", msg.SessionName)
	}

	//set the standard header
	header := types.NewHeader(msg.LogonInitiator.Header.BodyLength, msg.LogonInitiator.Header.MsgType, senderCompID.Did, targetCompID.Did, msg.LogonInitiator.Header.MsgSeqNum, msg.LogonInitiator.Header.SendingTime)

	//set the FIX Version
	header.BeginString = msg.FIXVersionByInitiator()

	//set the standard trailer
	trailer := types.NewTrailer(msg.LogonInitiator.Trailer.CheckSum)

	// set the logon initiator message
	logonInitiator := types.NewLogonInitiator(header, msg.LogonInitiator.EncryptMethod, msg.LogonInitiator.HeartBtInt, trailer)

	var newInitiatorSession = types.Sessions{
		SessionName:      msg.SessionName,
		LogonInitiator:   &logonInitiator,
		Status:           "logon-request",
		IsAccepted:       false,
		InitiatorAddress: msg.InitiatorAddress,
	}

	//set new Initiator logon session to store
	k.SetSessions(ctx, msg.SessionName, newInitiatorSession)

	return &types.MsgLogonInitiatorResponse{}, nil
}
