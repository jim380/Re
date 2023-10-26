package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestLogoutInitiator() {
	type args struct {
		session            types.Sessions
		msgLogoutInitiator types.MsgLogoutInitiator
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			"Valid Logout Request",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					Status:     constants.LoggedInStatus,
					IsAccepted: true,
				},
				msgLogoutInitiator: *types.NewMsgLogoutInitiator(suite.address[0].String(), "sessionID", types.SessionLogoutInitiator{Text: "Request to Logout from the current session"}),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Logout Request Initiator's Account Address is not found the session",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					Status:     constants.LoggedInStatus,
					IsAccepted: true,
				},
				msgLogoutInitiator: *types.NewMsgLogoutInitiator(suite.address[2].String(), "sessionID", types.SessionLogoutInitiator{Text: "Request to Logout from the current session"}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"SessionID does not exist",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					LogonAcceptor: &types.LogonAcceptor{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					Status:     constants.LoggedInStatus,
					IsAccepted: true,
				},
				msgLogoutInitiator: *types.NewMsgLogoutInitiator(suite.address[0].String(), "sessionIDxyz", types.SessionLogoutInitiator{Text: "Request to Logout from the current session"}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// set session
			suite.fixKeeper.SetSessions(suite.ctx, tc.args.session.SessionID, tc.args.session)

			// call LogoutInitiator method
			res, err := suite.msgServer.LogoutInitiator(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogoutInitiator)
			_, found := suite.fixKeeper.GetSessionLogout(suite.ctx, tc.args.msgLogoutInitiator.SessionID)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Nil(res)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestLogoutAcceptor() {
	type args struct {
		sessionLogout     types.SessionLogout
		msgLogoutAcceptor types.MsgLogoutAcceptor
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			"Logout Request can be accepted when the session is not logged out already",
			args{
				sessionLogout: types.SessionLogout{
					SessionID: "sessionID",
					SessionLogoutInitiator: &types.SessionLogoutInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "5",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						Text: "Logout Request",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SessionLogoutAcceptor: nil,
				},
				msgLogoutAcceptor: *types.NewMsgLogoutAcceptor(suite.address[1].String(), "sessionID", types.SessionLogoutAcceptor{Text: "Logout confirmed"}),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Logout Acceptor Address is wrong",
			args{
				sessionLogout: types.SessionLogout{
					SessionID: "sessionID",
					SessionLogoutInitiator: &types.SessionLogoutInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "5",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						Text: "Logout Request",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SessionLogoutAcceptor: nil,
				},
				msgLogoutAcceptor: *types.NewMsgLogoutAcceptor(suite.address[0].String(), "sessionID", types.SessionLogoutAcceptor{Text: "Logout confirmed"}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logout Acceptor address is not found in the session logout",
			args{
				sessionLogout: types.SessionLogout{
					SessionID: "sessionID",
					SessionLogoutInitiator: &types.SessionLogoutInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "5",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						Text: "Logout Request",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SessionLogoutAcceptor: nil,
				},
				msgLogoutAcceptor: *types.NewMsgLogoutAcceptor(suite.address[2].String(), "sessionID", types.SessionLogoutAcceptor{Text: "Logout confirmed"}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logout Request has been accepted",
			args{
				sessionLogout: types.SessionLogout{
					SessionID: "sessionID",
					SessionLogoutInitiator: &types.SessionLogoutInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "5",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "8",
						},
						Text: "Logout Request",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SessionLogoutAcceptor: &types.SessionLogoutAcceptor{},
				},
				msgLogoutAcceptor: *types.NewMsgLogoutAcceptor(suite.address[1].String(), "sessionID", types.SessionLogoutAcceptor{Text: "Logout confirmed"}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// set session logout
			suite.fixKeeper.SetSessionLogout(suite.ctx, tc.args.sessionLogout.SessionID, tc.args.sessionLogout)

			// call LogoutAcceptor method
			res, err := suite.msgServer.LogoutAcceptor(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogoutAcceptor)
			_, found := suite.fixKeeper.GetSessionLogout(suite.ctx, tc.args.msgLogoutAcceptor.SessionID)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Nil(res)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}
