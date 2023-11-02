package fixoperations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	fixTypes "github.com/jim380/Re/x/fix/types"
)

type SessionData struct {
	SessionID    string
	MsgType      string
	MsgType_     string
	SenderCompID string
	TargetCompID string
	MsgSeqNum    int64
	SendingTime  string
	ChainID      string
}

// SessionOperation creates and sets the sessions for both logon initiator and acceptor
func SessionOperation(ctx sdk.Context, k fixKeeper.Keeper, sessionData SessionData) {
	session := fixTypes.Sessions{
		SessionID: sessionData.SessionID,
		LogonInitiator: &fixTypes.LogonInitiator{
			Header: &fixTypes.Header{
				BeginString:  "fix4.4",
				MsgType:      sessionData.MsgType,
				SenderCompID: sessionData.SenderCompID,
				TargetCompID: sessionData.TargetCompID,
				MsgSeqNum:    sessionData.MsgSeqNum,
				SendingTime:  sessionData.SendingTime,
				ChainID:      sessionData.ChainID,
			},
			Trailer: &fixTypes.Trailer{
				CheckSum: 0,
			},
		},
		LogonAcceptor: &fixTypes.LogonAcceptor{
			Header: &fixTypes.Header{
				BeginString:  "fix4.4",
				MsgType:      sessionData.MsgType_,
				SenderCompID: sessionData.TargetCompID,
				TargetCompID: sessionData.SenderCompID,
				MsgSeqNum:    sessionData.MsgSeqNum,
				SendingTime:  sessionData.SendingTime,
				ChainID:      sessionData.ChainID,
			},
			Trailer: &fixTypes.Trailer{
				CheckSum: 0,
			},
		},
		Status:     "loggedIn",
		IsAccepted: true,
	}

	k.SetSessions(ctx, sessionData.SessionID, session)
}
