package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/x/fix/types"
)

const (
	sessionNameLength = 10
)

func (k msgServer) LogonInitiator(goCtx context.Context, msg *types.MsgLogonInitiator) (*types.MsgLogonInitiatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// DID will be used senderCompID and targetCompID for both parties
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

	// same DID can not be used for intiating and accepting in the same session
	//if senderCompID.Did == targetCompID.Did {

	//}

	//check for if this session Name exists already
	_, found := k.GetSessions(ctx, msg.SessionName)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrSessionNameFound, "Session Name: %s", msg.SessionName)
	}

	//get the sending time
	sendingTime := time.Now().UTC().Format("20060102-15:04:05.000")
	msg.LogonInitiator.Header.SendingTime = sendingTime

	//get MsgSeqNum from  GetSessionsCount
	msgSeqNum := k.GetSessionsCount(ctx)
	msg.LogonInitiator.Header.MsgSeqNum = int64(msgSeqNum) + 1

	//body length of logon message
	msgBody := msg.SessionName + msg.InitiatorAddress + strconv.FormatInt(msg.LogonInitiator.EncryptMethod, 10) + strconv.FormatInt(msg.LogonInitiator.HeartBtInt, 10)
	bodyLength := msg.BodyLength(msgBody)
	msg.LogonInitiator.Header.BodyLength = bodyLength

	//set the standard header
	header := types.NewHeader(msg.LogonInitiator.Header.MsgType, senderCompID.Did, targetCompID.Did)
	//header.BeginString = msg.FIXVersion()

	// set the standard trailer
	//trailer := types.NewTrailer(bodyLength)

	// set the logon initiator message
	logonInitiator := types.NewLogonInitiator(header, msg.LogonInitiator.EncryptMethod, msg.LogonInitiator.HeartBtInt)

	var newSession = types.Sessions{
		SessionName:      msg.SessionName,
		LogonInitiator:   &logonInitiator,
		IsLoggedIn:       false,
		IsAccepted:       false,
		InitiatorAddress: msg.InitiatorAddress,
	}

	//set new logon session to store
	k.SetSessions(ctx, msg.SessionName, newSession)

	return &types.MsgLogonInitiatorResponse{}, nil
}
