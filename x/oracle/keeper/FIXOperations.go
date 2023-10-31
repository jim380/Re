package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fixKeeper "github.com/jim380/Re/x/fix/keeper"
	fixTypes "github.com/jim380/Re/x/fix/types"
)

// SessionOperation creates and sets the sessions for both logon initiator and acceptor
func SessionOperation(ctx sdk.Context, k fixKeeper.Keeper, sessionID, senderCompID, targetCompID string, msgSeqNum int64, sendingTime, chainID string) {

	// session setup logic

	session := fixTypes.Sessions{
		SessionID: sessionID,
		LogonInitiator: &fixTypes.LogonInitiator{
			Header: &fixTypes.Header{
				BeginString:  "fix4.4",
				MsgType:      "A",
				SenderCompID: senderCompID,
				TargetCompID: targetCompID,
				MsgSeqNum:    msgSeqNum,
				SendingTime:  sendingTime,
				ChainID:      chainID,
			},
			Trailer: &fixTypes.Trailer{
				CheckSum: 0,
			},
		},
		LogonAcceptor: &fixTypes.LogonAcceptor{
			Header: &fixTypes.Header{
				BeginString:  "fix4.4",
				MsgType:      "A",
				SenderCompID: targetCompID,
				TargetCompID: senderCompID,
				MsgSeqNum:    msgSeqNum,
				SendingTime:  sendingTime,
				ChainID:      chainID,
			},
			Trailer: &fixTypes.Trailer{
				CheckSum: 0,
			},
		},
		Status:     "loggedIn",
		IsAccepted: true,
	}

	k.SetSessions(ctx, sessionID, session)
}
