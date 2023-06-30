package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestLogonInitiator() {
	type args struct {
		msgRegisterAccount1 types.MsgRegisterAccount
		msgRegisterAccount2 types.MsgRegisterAccount
		msgLogonInitiator   types.MsgLogonInitiator
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
			"Valid Logon Request",
			args{
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[0].String(),
					Address:          suite.address[0].String(),
					CompanyName:      "",
					Website:          "",
					SocialMediaLinks: "",
				},
				msgRegisterAccount2: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID1", types.LogonInitiator{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
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
				}),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Same Address can not used for SenderCompID & TargetCompID",
			args{
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[0].String(),
					Address:          suite.address[0].String(),
					CompanyName:      "",
					Website:          "",
					SocialMediaLinks: "",
				},
				msgRegisterAccount2: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID2", types.LogonInitiator{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[0].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					EncryptMethod: 1,
					HeartBtInt:    1,
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"InitiatorAddress is different from SenderCompID",
			args{
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[0].String(),
					Address:          suite.address[0].String(),
					CompanyName:      "",
					Website:          "",
					SocialMediaLinks: "",
				},
				msgRegisterAccount2: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonInitiator: *types.NewMsgLogonInitiator(suite.address[0].String(), "sessionID3", types.LogonInitiator{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[2].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					EncryptMethod: 1,
					HeartBtInt:    1,
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				}),
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

			// call RegisterAccount method
			_, _ = suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &tc.args.msgRegisterAccount1)
			_, _ = suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &tc.args.msgRegisterAccount2)

			// call LogonInitiator method
			res, err := suite.msgServer.LogonInitiator(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogonInitiator)
			session, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgLogonInitiator.SessionID)

			if tc.errArgs.shouldPass {
				suite.Require().True(found)
				suite.Require().NoError(err, tc.name)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(session)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().Nil(res)
				suite.Require().Empty(session)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestLogonAcceptor() {
	type args struct {
		session             types.Sessions
		msgRegisterAccount1 types.MsgRegisterAccount
		msgLogonAcceptor    types.MsgLogonAcceptor
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
			"Logon Request can be Accepted",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonAcceptor: *types.NewMsgLogonAcceptor(suite.address[1].String(), "sessionID1", types.LogonAcceptor{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
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
				}),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Same Address can not used for SenderCompID & TargetCompID",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonAcceptor: *types.NewMsgLogonAcceptor(suite.address[1].String(), "sessionID1", types.LogonAcceptor{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					EncryptMethod: 1,
					HeartBtInt:    1,
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Wrong SessionID",
			args{
				session: types.Sessions{
					SessionID: "sessionID2",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonAcceptor: *types.NewMsgLogonAcceptor(suite.address[1].String(), "sessionID1", types.LogonAcceptor{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
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
				}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been Accepted",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LoggedInStatus,
					IsAccepted:    true,
				},
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonAcceptor: *types.NewMsgLogonAcceptor(suite.address[1].String(), "sessionID1", types.LogonAcceptor{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
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
				}),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been rejected",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.RejectedStatus,
					IsAccepted:    false,
				},
				msgRegisterAccount1: types.MsgRegisterAccount{
					Creator:          suite.address[1].String(),
					Address:          suite.address[1].String(),
					CompanyName:      "cypherCore",
					Website:          "cypherCore.io",
					SocialMediaLinks: "@cypherCore",
				},
				msgLogonAcceptor: *types.NewMsgLogonAcceptor(suite.address[1].String(), "sessionID1", types.LogonAcceptor{
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
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
				}),
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

			// call RegisterAccount method
			_, _ = suite.msgServer.RegisterAccount(sdk.WrapSDKContext(suite.ctx), &tc.args.msgRegisterAccount1)

			// call LogonAcceptor method
			res, err := suite.msgServer.LogonAcceptor(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogonAcceptor)
			_, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgLogonAcceptor.SessionID)

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

func (suite *KeeperTestSuite) TestLogonReject() {
	type args struct {
		session        types.Sessions
		msgLogonReject types.MsgLogonReject
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
			"Logon Request can be Rejected when the request has not been accepted or rejected",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgLogonReject: types.MsgLogonReject{
					AcceptorAddress: suite.address[1].String(),
					SessionID:       "sessionID1",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[0].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					Text: "Not Ready",
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				},
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"SessionID in Logon Request does not match the sessionID provided by Acceptor",
			args{
				session: types.Sessions{
					SessionID: "sessionID1xyz",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgLogonReject: types.MsgLogonReject{
					AcceptorAddress: suite.address[1].String(),
					SessionID:       "sessionID1",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[0].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					Text: "Not Ready",
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Wrong Account Address",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgLogonReject: types.MsgLogonReject{
					AcceptorAddress: suite.address[0].String(),
					SessionID:       "sessionID1",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					Text: "Not Ready",
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been accepted",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LoggedInStatus,
					IsAccepted:    true,
				},
				msgLogonReject: types.MsgLogonReject{
					AcceptorAddress: suite.address[1].String(),
					SessionID:       "sessionID1",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[0].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					Text: "Not Ready",
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been Rejected",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.RejectedStatus,
					IsAccepted:    false,
				},
				msgLogonReject: types.MsgLogonReject{
					AcceptorAddress: suite.address[1].String(),
					SessionID:       "sessionID1",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[1].String(),
						TargetCompID: suite.address[0].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					Text: "Not Ready",
					Trailer: &types.Trailer{
						CheckSum: 10,
					},
				},
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

			// call LogonReject method
			res, err := suite.msgServer.LogonReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgLogonReject)
			_, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgLogonReject.SessionID)

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

func (suite *KeeperTestSuite) TestTerminateLogon() {
	type args struct {
		session           types.Sessions
		msgTerminateLogon types.MsgTerminateLogon
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
			"Logon Request can be Terminated when the request has not been accepted or rejected",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgTerminateLogon: types.MsgTerminateLogon{
					InitiatorAddress: suite.address[0].String(),
					SessionID:        "sessionID1",
					Address:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"SessionID is not found",
			args{
				session: types.Sessions{
					SessionID: "sessionID1xyz",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    false,
				},
				msgTerminateLogon: types.MsgTerminateLogon{
					InitiatorAddress: suite.address[0].String(),
					SessionID:        "sessionID1",
					Address:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Wrong Account Address",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LogonRequestStatus,
					IsAccepted:    true,
				},
				msgTerminateLogon: types.MsgTerminateLogon{
					InitiatorAddress: suite.address[1].String(),
					SessionID:        "sessionID1",
					Address:          suite.address[1].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been accepted",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.LoggedInStatus,
					IsAccepted:    false,
				},
				msgTerminateLogon: types.MsgTerminateLogon{
					InitiatorAddress: suite.address[0].String(),
					SessionID:        "sessionID1",
					Address:          suite.address[0].String(),
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Logon Request has been rejected",
			args{
				session: types.Sessions{
					SessionID: "sessionID1",
					LogonInitiator: &types.LogonInitiator{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
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
					LogonAcceptor: &types.LogonAcceptor{},
					Status:        constants.RejectedStatus,
					IsAccepted:    false,
				},
				msgTerminateLogon: types.MsgTerminateLogon{
					InitiatorAddress: suite.address[0].String(),
					SessionID:        "sessionID1",
					Address:          suite.address[0].String(),
				},
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

			// call TerminateLogon method
			res, err := suite.msgServer.TerminateLogon(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTerminateLogon)
			_, found := suite.fixKeeper.GetSessions(suite.ctx, tc.args.msgTerminateLogon.SessionID)

			if tc.errArgs.shouldPass {
				suite.Require().False(found)
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
