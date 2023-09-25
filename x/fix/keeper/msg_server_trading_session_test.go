package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestTradingSessionStatusRequest() {
	type args struct {
		session                        types.Sessions
		msgTradingSessionStatusRequest types.MsgTradingSessionStatusRequest
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
			"Valid Trading Session Status Request",
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
				msgTradingSessionStatusRequest: *types.NewMsgTradingSessionStatusRequest(suite.address[0].String(), "sessionID", "tradSesReqID", "tradingSessionID", "tradingSessionSubID", "marketID", "subscriptionRequest", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", 2, 2, 1, "2023-09-25", "1:09:05", "1:10:05", "2023-09-26"),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"Logon is not establishlished between both parties",
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
					Status:     "Logged Out",
					IsAccepted: true,
				},
				msgTradingSessionStatusRequest: *types.NewMsgTradingSessionStatusRequest(suite.address[0].String(), "sessionID", "tradSesReqID", "tradingSessionID", "tradingSessionSubID", "marketID", "subscriptionRequest", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", 2, 2, 1, "2023-09-25", "1:09:05", "1:10:05", "2023-09-26"),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Parties involved in the session are not accepted",
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
				msgTradingSessionStatusRequest: *types.NewMsgTradingSessionStatusRequest(suite.address[2].String(), "sessionID", "tradSesReqID", "tradingSessionID", "tradingSessionSubID", "marketID", "subscriptionRequest", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", 2, 2, 1, "2023-09-25", "1:09:05", "1:10:05", "2023-09-26"),
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
				msgTradingSessionStatusRequest: *types.NewMsgTradingSessionStatusRequest(suite.address[0].String(), "sessionIDxnxn4", "tradSesReqID", "tradingSessionID", "tradingSessionSubID", "marketID", "subscriptionRequest", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", 2, 2, 1, "2023-09-25", "1:09:05", "1:10:05", "2023-09-26"),
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

			// call Trading Session Status Request
			res, err := suite.msgServer.TradingSessionStatusRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionStatusRequest)

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

func (suite *KeeperTestSuite) TestTradingSessionStatus() {
	type args struct {
		session                 types.Sessions
		msgTradingSessionStatus types.MsgTradingSessionStatus
		tradingSession          types.TradingSession
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
			"Valid Trading Session Status",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[1].String(), "sessionID", "TradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
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
			"Invalid Session Id",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[1].String(), "sessionIDdnvdnvdn", "TradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
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
			" User responding is not the recipient of the Trading Session Status Request",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[2].String(), "sessionID", "TradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
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
			"Trading Session Status Request is acknowledged",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[1].String(), "sessionID", "TradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionStatusRequestReject: &types.TradingSessionStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						RefSeqNum:           "RefSeqNum",
						RefMsgType:          "RefMsgType",
						SessionRejectReason: 1,
						Text:                "Text",
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
			"The account cannot used for creating Trading Session Status Request and Trading Session Status with the same TradSesReqID",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[0].String(), "sessionID", "TradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
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
			"Trading Session Status Request is rejected already",
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
				msgTradingSessionStatus: *types.NewMsgTradingSessionStatus(suite.address[1].String(), "sessionID", "tradSesReqID", "tradingSessionID", 2, 2, "TradSesStartTime", "tradSesOpenTime", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "tradSesEndTime", 6, "tradSesHighPx", "tradSesLowPx", "securityID", "securityIDSource", "symbol", "securityExchange", "marketSegmentID", "marketID"),
				tradingSession: types.TradingSession{
					SessionID: "sessionID",
					TradingSessionStatusRequest: &types.TradingSessionStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "TradSesReqID",
						MarketID:              "MarketID",
						SubscriptionRequest:   "SubscriptionRequest",
						SecurityID:            "SecurityID",
						SecurityIDSource:      "SecurityIDSource",
						Symbol:                "Symbol",
						SecurityExchange:      "SecurityExchange",
						MarketSegmentID:       "MarketSegmentID",
						TradSesReqType:        1,
						TradSesSubReqType:     2,
						TradSesMode:           1,
						TradingSessionDate:    "2023-09-25",
						TradingSessionTime:    time.Now().Format("2006-01-02 15:04:05"),
						TradingSessionSubTime: time.Now().Format("2006-01-02 15:04:05"),
						ExpirationDate:        "2023-09-25",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionStatusRequestReject: &types.TradingSessionStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						RefSeqNum:           "RefSeqNum",
						RefMsgType:          "RefMsgType",
						SessionRejectReason: 1,
						Text:                "Text",
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
			// set order
			suite.fixKeeper.SetTradingSession(suite.ctx, tc.args.tradingSession.TradingSessionStatusRequest.TradSesReqID, tc.args.tradingSession)

			// call Order cancel request method
			res, err := suite.msgServer.TradingSessionStatus(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionStatus)

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

/*func (suite *KeeperTestSuite) TestOrderExecutionReport() {
	type args struct {
		session                 types.Sessions
		MsgOrderExecutionReport types.MsgOrderExecutionReport
		order                   types.Orders
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
			"Valid Order Execution Report",
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
				MsgOrderExecutionReport: *types.NewMsgOrderExecutionReport(suite.address[1].String(), "sessionID", "clOrdID", "orderID", "execID", "ordStatus", "execType", "symbol", 1, "500", "700", 5, 2, 1, 1, 2, 1, "Order execution report"),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "clOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			"Same account address can not used for creating New Single Order and Execution Report with the same ClOrdID",
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
				MsgOrderExecutionReport: *types.NewMsgOrderExecutionReport(suite.address[0].String(), "sessionID", "clOrdID", "orderID", "execID", "ordStatus", "execType", "symbol", 1, "500", "700", 5, 2, 1, 1, 2, 1, "Order execution report"),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "clOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			"Account address is not the creator of the order execution report",
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
				MsgOrderExecutionReport: *types.NewMsgOrderExecutionReport(suite.address[2].String(), "sessionID", "clOrdID", "orderID", "execID", "ordStatus", "execType", "symbol", 1, "500", "700", 5, 2, 1, 1, 2, 1, "Order execution report"),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "clOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
				MsgOrderExecutionReport: *types.NewMsgOrderExecutionReport(suite.address[1].String(), "sessionID575757", "clOrdID", "orderID", "execID", "ordStatus", "execType", "symbol", 1, "500", "700", 5, 2, 1, 1, 2, 1, "Order execution report"),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "clOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			// set order
			suite.fixKeeper.SetOrders(suite.ctx, tc.args.order.ClOrdID, tc.args.order)

			// call Order execution report method
			res, err := suite.msgServer.OrderExecutionReport(sdk.WrapSDKContext(suite.ctx), &tc.args.MsgOrderExecutionReport)

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

func (suite *KeeperTestSuite) TestOrderCancelReject() {
	type args struct {
		session              types.Sessions
		msgOrderCancelReject types.MsgOrderCancelReject
		order                types.Orders
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
			"Valid Order Cancel Reject",
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
				msgOrderCancelReject: *types.NewMsgOrderCancelReject(suite.address[1].String(), "sessionID", "orderID", "OrgclOrdID", "clOrdID", 2, 2),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "OrgclOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			"Same account address can not used for creating New Single Order and Execution Report with the same ClOrdID",
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
				msgOrderCancelReject: *types.NewMsgOrderCancelReject(suite.address[0].String(), "sessionID", "orderID", "OrgclOrdID", "clOrdID", 2, 2),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "OrgclOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			"Account address is not the creator of the order cancel reject",
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
				msgOrderCancelReject: *types.NewMsgOrderCancelReject(suite.address[2].String(), "sessionID", "orderID", "OrgclOrdID", "clOrdID", 2, 2),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "clOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
				msgOrderCancelReject: *types.NewMsgOrderCancelReject(suite.address[0].String(), "sessionID8778787", "OrgclOrdID", "orgClOrdID", "clOrdID", 2, 2),
				order: types.Orders{
					SessionID: "sessionID",
					Header: &types.Header{
						BeginString:  "FIX4.4",
						BodyLength:   10,
						MsgType:      "A",
						SenderCompID: suite.address[0].String(),
						TargetCompID: suite.address[1].String(),
						MsgSeqNum:    1,
						SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
					},
					ClOrdID:      "OrgclOrdID",
					Symbol:       "TEXT_SYMBOL",
					Side:         1,
					OrderQty:     "10",
					OrdType:      1,
					Price:        "500",
					TimeInForce:  5,
					Text:         "This is an order",
					TransactTime: time.Now().Format("2006-01-02 15:04:05"),
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
			// set order
			suite.fixKeeper.SetOrders(suite.ctx, tc.args.order.ClOrdID, tc.args.order)

			// call Order cancel reject method
			res, err := suite.msgServer.OrderCancelReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgOrderCancelReject)

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
*/
