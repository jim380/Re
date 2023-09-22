package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestNewOrderSingle() {
	type args struct {
		session            types.Sessions
		msgNewOrderSingle types.MsgNewOrderSingle
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
			"Valid New Single Order",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "D",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
							MsgType:      "D",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[0].String(), "sessionID", "clOrdID", "TEXT_SYMBOL", 1, "100", 1,"500", 4, "New Order Single"),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Parties involved are not in the session",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "D",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
							MsgType:      "D",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[2].String(), "sessionID", "clOrdID", "TEXT_SYMBOL", 1, "100", 1,"500", 4, "New Order Single"),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"The sessionID does not exist",
			args{
				session: types.Sessions{
					SessionID: "sessionID",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "D",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
							MsgType:      "D",
							SenderCompID: suite.address[1].String(),
							TargetCompID: suite.address[0].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[2].String(), "sessionIDxyzs", "clOrdID", "TEXT_SYMBOL", 1, "100", 1,"500", 4, "New Order Single"),
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
			suite.fixKeeper.SetSessions(tc.args.session.SessionID, tc.args.session)
			// set order
			suite.fixKeeper.SetOrders(suite.ctx, tc.args.msgNewOrderSingle.ClOrdID, tc.args.msgNewOrderSingle)

			// call  New Single Order method
			res, err := suite.msgServer.NewOrderSingle(sdk.WrapSDKContext(suite.ctx), &tc.args.msgNewOrderSingle)

			if tc.errArgs.shouldPass {
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

/*func (suite *KeeperTestSuite) TestLogoutAcceptor() {
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
*/