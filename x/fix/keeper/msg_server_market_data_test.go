package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestMarketDataRequest() {
	type args struct {
		session              types.Sessions
		msgMarketDataRequest types.MsgMarketDataRequest
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
			"Valid Market Data Response",
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
				msgMarketDataRequest: *types.NewMsgMarketDataRequest(suite.address[0].String(), "sessionID", "MdReqID", 1, 7, 5, 3, "symbol"),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Logon status is not logged in",
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
						},
						EncryptMethod: 1,
						HeartBtInt:    1,
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					Status:     "logged_out",
					IsAccepted: true,
				},
				msgMarketDataRequest: *types.NewMsgMarketDataRequest(suite.address[0].String(), "sessionID", "MdReqID", 1, 7, 5, 3, "symbol"),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"The parties involved in a session are the ones using the sessionID and are able to create Market Data Request",
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
				msgMarketDataRequest: *types.NewMsgMarketDataRequest(suite.address[2].String(), "sessionID", "MdReqID", 1, 7, 5, 3, "symbol"),
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
					LogonAcceptor: &types.LogonAcceptor{
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
					},
					Status:     constants.LoggedInStatus,
					IsAccepted: true,
				},
				msgMarketDataRequest: *types.NewMsgMarketDataRequest(suite.address[0].String(), "sessionIDgugugu", "MdReqID", 1, 7, 5, 3, "symbol"),
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
			// call  Market Data Request method
			res, err := suite.msgServer.MarketDataRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgMarketDataRequest)

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

func (suite *KeeperTestSuite) TestMarketDataSnapshotFullRefresh() {
	type args struct {
		session                          types.Sessions
		msgMarketDataSnapshotFullRefresh types.MsgMarketDataSnapshotFullRefresh
		marketData                       types.MarketData
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
			"Valid Security List Response",
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
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[1].String(), "sessionID", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
				},
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Market Data Request creator address is the same responding to the Market Data Request",
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
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[0].String(), "sessionID", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"The user responding is the recipient of the market data request",
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
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[2].String(), "sessionID", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Market Data Request to be acknowledged has been rejected",
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
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[1].String(), "sessionID", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataRequestReject: &types.MarketDataRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:        "MdReqID",
						MdReqRejReason: 2,
						Text:           "Text",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
				},
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Market Data Request is acknowledged already",
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
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[1].String(), "sessionID", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataSnapShotFullRefresh: &types.MarketDataSnapShotFullRefresh{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						Symbol:      "Symbol",
						NoMDEntries: 6,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 10,
								MdEntryType:    12,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 1,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
				},
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
					LogonAcceptor: &types.LogonAcceptor{
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
					},
					Status:     constants.LoggedInStatus,
					IsAccepted: true,
				},
				msgMarketDataSnapshotFullRefresh: *types.NewMsgMarketDataSnapshotFullRefresh(suite.address[1].String(), "sessionIDnfenfk", "MdReqID", "symbol", 7, []*types.MDEntry{{
					MdUpdateAction: 2,
					MdEntryType:    2,
					MdEntryPx:      "MdEntryPx",
					MdEntrySize:    "MdEntrySize",
				},
					{
						MdUpdateAction: 9,
						MdEntryType:    8,
						MdEntryPx:      "MdEntryPx",
						MdEntrySize:    "MdEntrySize",
					}}),
				marketData: types.MarketData{
					SessionID: "sessionID",
					MarketDataRequest: &types.MarketDataRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:                 "MdReqID",
						SubscriptionRequestType: 3,
						MarketDepth:             7,
						MdUpdateType:            6,
						NoRelatedSym:            7,
						Symbol:                  "Symbol",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					MarketDataIncremental: &types.MarketDataIncremental{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MdReqID:     "MdReqID",
						NoMDEntries: 7,
						MdEntries: []*types.MDEntry{
							{
								MdUpdateAction: 2,
								MdEntryType:    2,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
							{
								MdUpdateAction: 9,
								MdEntryType:    8,
								MdEntryPx:      "MdEntryPx",
								MdEntrySize:    "MdEntrySize",
							},
						},
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
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
			// set market data
			suite.fixKeeper.SetMarketData(suite.ctx, tc.args.marketData.MarketDataRequest.MdReqID, tc.args.marketData)
			// call market data snapshot full refresh method
			res, err := suite.msgServer.MarketDataSnapshotFullRefresh(sdk.WrapSDKContext(suite.ctx), &tc.args.msgMarketDataSnapshotFullRefresh)

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
