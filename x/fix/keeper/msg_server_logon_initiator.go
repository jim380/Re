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

	// check for if this session Name exists already
	session, found := k.GetSessions(ctx, msg.SessionName)
	if !found {

	}
	if session.LogonInitiator.Header.MsgSeqNum == msg.LogonInitiator.Header.MsgSeqNum {

	}

	// DID will be used senderCompID and targetCompID for both parties
	// get DID from registered accounts
	senderCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.SenderCompID)
	if senderCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.LogonInitiator.Header.SenderCompID)
	}

	targetCompID := k.GetAccount(ctx, msg.LogonInitiator.Header.TargetCompID)
	if targetCompID.Empty() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDidDocument, "DID Document: %s", msg.LogonInitiator.Header.TargetCompID)
	}

	// generate new session name for each session
	sessionName, err := generateRandomString(sessionNameLength)
	if err != nil {
		panic(err)
	}

	// get the sending time
	sendingTime := time.Now().UTC().Format("20060102-15:04:05.000")

	// get MsgSeqNum from  GetSessionsCount
	msgSeqNum := k.GetSessionsCount(ctx)

	// body length of logon message
	body := msg.SessionName + msg.InitiatorAddress + strconv.FormatInt(msg.LogonInitiator.EncryptMethod, 10) + strconv.FormatInt(msg.LogonInitiator.HeartBtInt, 10)
	bodyLength := msg.BodyLength(body)

	//set the standard header
	header := types.Header{
		BeginString:  msg.FIXVersion(),
		BodyLength:   bodyLength,
		MsgType:      msg.LogonInitiator.Header.MsgType,
		SenderCompID: senderCompID.Did,
		TargetCompID: targetCompID.Did,
		MsgSeqNum:    int64(msgSeqNum) + 1,
		SendingTime:  sendingTime,
	}

	// set the standard trailer
	trailer := types.Trailer{
		CheckSum: msg.LogonInitiator.Trailer.CheckSum,
	}

	// set the logon initiator message
	logonInitiator := types.LogonInitiator{
		Header:        &header,
		EncryptMethod: msg.LogonInitiator.EncryptMethod,
		HeartBtInt:    msg.LogonInitiator.HeartBtInt,
		Trailer:       &trailer,
	}

	var newSession = types.Sessions{
		SessionName:      sessionName,
		LogonInitiator:   &logonInitiator,
		IsLoggedIn:       false,
		IsAccepted:       false,
		InitiatorAddress: msg.InitiatorAddress,
	}

	//set new logon session to store
	k.SetSessions(ctx, msg.SessionName, newSession)

	return &types.MsgLogonInitiatorResponse{}, nil
}
