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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "cosmos",
						},
						TradSesReqID:           "TradSesReqID",
						TradingSessionID:       "TradingSessionID",
						TradSesStatus:          1,
						TradSesStatusRejReason: 1,
						TradSesStartTime:       "TradSesStartTime",
						TradSesOpenTime:        "TradSesOpenTime",
						TradSesPreCloseTime:    "TradSesPreCloseTime",
						TradSesCloseTime:       "TradSesCloseTime",
						TradSesEndTime:         "TradSesEndTime",
						TotalVolumeTraded:      1,
						TradSesHighPx:          "TradSesHighPx",
						TradSesLowPx:           "TradSesLowPx",
						SecurityID:             "SecurityID",
						SecurityIDSource:       "SecurityIDSource",
						Symbol:                 "Symbol",
						SecurityExchange:       "SecurityExchange",
						MarketSegmentID:        "MarketSegmentID",
						MarketID:               "MarketID",
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
			"The account cannot be used for creating Trading Session Status Request and Trading Session Status with the same TradSesReqID",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
			// set trading session
			suite.fixKeeper.SetTradingSession(suite.ctx, tc.args.tradingSession.TradingSessionStatusRequest.TradSesReqID, tc.args.tradingSession)

			// call trading session status method
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

func (suite *KeeperTestSuite) TestTradingSessionStatusRequestReject() {
	type args struct {
		session                              types.Sessions
		msgTradingSessionStatusRequestReject types.MsgTradingSessionStatusRequestReject
		tradingSession                       types.TradingSession
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
			"Valid Trading Session Status Request Reject",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[1].String(), "sessionID", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[1].String(), "sessionIDfeffed5", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
			"Account address is not the creator of the trading session status request reject",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[2].String(), "sessionID", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[1].String(), "sessionID", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[1].String(), "sessionID", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
							ChainID:      "cosmos",
						},
						TradSesReqID:           "TradSesReqID",
						TradingSessionID:       "TradingSessionID",
						TradSesStatus:          1,
						TradSesStatusRejReason: 1,
						TradSesStartTime:       "TradSesStartTime",
						TradSesOpenTime:        "TradSesOpenTime",
						TradSesPreCloseTime:    "TradSesPreCloseTime",
						TradSesCloseTime:       "TradSesCloseTime",
						TradSesEndTime:         "TradSesEndTime",
						TotalVolumeTraded:      1,
						TradSesHighPx:          "TradSesHighPx",
						TradSesLowPx:           "TradSesLowPx",
						SecurityID:             "SecurityID",
						SecurityIDSource:       "SecurityIDSource",
						Symbol:                 "Symbol",
						SecurityExchange:       "SecurityExchange",
						MarketSegmentID:        "MarketSegmentID",
						MarketID:               "MarketID",
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
			"Same account can not used for creating Trading Session Status Request Reject",
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
							ChainID:      "cosmos",
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
							ChainID:      "cosmos",
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
				msgTradingSessionStatusRequestReject: *types.NewMsgTradingSessionStatusRequestReject(suite.address[0].String(), "sessionID", "refSeqNum", "refMsgType", 2, "text"),
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
							ChainID:      "cosmos",
						},
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						TradSesReqID:          "refSeqNum",
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
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// set session
			suite.fixKeeper.SetSessions(suite.ctx, tc.args.session.SessionID, tc.args.session)
			// set trading session
			suite.fixKeeper.SetTradingSession(suite.ctx, tc.args.tradingSession.TradingSessionStatusRequest.TradSesReqID, tc.args.tradingSession)

			// call trading session request reject method
			res, err := suite.msgServer.TradingSessionStatusRequestReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionStatusRequestReject)

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
