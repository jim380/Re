package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

/*func (suite *KeeperTestSuite) TestTradingSessionListRequest() {
	type args struct {
		session                      types.Sessions
		msgTradingSessionListRequest types.MsgTradingSessionListRequest
		tradingSession               types.TradingSession
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
			"Valid Trading Session List Request",
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
				msgTradingSessionListRequest: *types.NewMsgTradingSessionListRequest(suite.address[0].String(), "sessionID", "TradSesReqID", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "subscriptionRequestType"),
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"sessionID does not exist",
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
				msgTradingSessionListRequest: *types.NewMsgTradingSessionListRequest(suite.address[0].String(), "sessionIDgygygyg", "TradSesReqID", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "subscriptionRequestType"),
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
			"logon status not equals to loggedIn",
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
					Status:     "loggedOut",
					IsAccepted: true,
				},
				msgTradingSessionListRequest: *types.NewMsgTradingSessionListRequest(suite.address[0].String(), "sessionID", "TradSesReqID", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "subscriptionRequestType"),
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
			"Account address is not in session",
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
				msgTradingSessionListRequest: *types.NewMsgTradingSessionListRequest(suite.address[2].String(), "sessionID", "TradSesReqID", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "subscriptionRequestType"),
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
					TradingSessionStatus: &types.TradingSessionStatus{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
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
			// set trading session
			suite.fixKeeper.SetTradingSession(suite.ctx, tc.args.tradingSession.TradingSessionStatus.TradSesReqID, tc.args.tradingSession)
			// call  Trading session list request
			res, err := suite.msgServer.TradingSessionListRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionListRequest)

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
}*/

func (suite *KeeperTestSuite) TestTradingSessionListResponse() {
	type args struct {
		session                       types.Sessions
		msgTradingSessionListResponse types.MsgTradingSessionListResponse
		tradingSessionList            types.TradingSessionList
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
			"Valid Trading Session list response",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[1].String(), "sessionID", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"Account address is not in session",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[2].String(), "sessionID", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"Same account can not used for creating Trading Session List Request and Trading Session List Response with the same TradSesReqID",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[0].String(), "sessionID", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"Trading Session List Request is acknowledged already",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[1].String(), "sessionID", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionListResponse: &types.TradingSessionListResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:           "TradSesReqID",
						NoTradingSessions:      "NoTradingSessions",
						TradingSessionID:       "TradingSessionID",
						TradingSessionSubID:    "TradingSessionSubID",
						SecurityExchange:       "SecurityExchange",
						TradSesMethod:          "TradSesMethod",
						TradSesMode:            "TradSesMode",
						UnsolicitedIndicator:   "UnsolicitedIndicator",
						TradSesStatus:          "TradSesStatus",
						TradSesStatusRejReason: "TradSesStatusRejReason",
						TradSesStartTime:       "TradSesStartTime",
						TradSesOpenTime:        "TradSesOpenTime",
						TradSesPreCloseTime:    "TradSesPreCloseTime",
						TradSesCloseTime:       "TradSesCloseTime",
						TradSesEndTime:         "TradSesEndTime",
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
			"Trading Session List Request is rejected already",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[1].String(), "sessionID", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionListRequestReject: &types.TradingSessionListRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:           "TradSesReqID",
						TradSesStatus:          "TradSesStatus",
						TradSesStatusRejReason: "TradSesStatusRejReason",
						Text:                   "Text",
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
				msgTradingSessionListResponse: *types.NewMsgTradingSessionListResponse(suite.address[1].String(), "sessionIDnbbdjbd", "TradSesReqID", "noTradingSessions", "tradingSessionID", "tradingSessionSubID", "securityExchange", "tradSesMethod", "tradSesMode", "unsolicitedIndicator", "tradSesStatus", "tradSesStatusRejReason", "tradSesStartTime", "tradSesOpenTime", "tradSesPreCloseTime", "tradSesCloseTime", "tradSesEndTime"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			// set trading session list
			suite.fixKeeper.SetTradingSessionList(suite.ctx, tc.args.tradingSessionList.TradingSessionListRequest.TradSesReqID, tc.args.tradingSessionList)
			// call trading session list response
			res, err := suite.msgServer.TradingSessionListResponse(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionListResponse)

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

func (suite *KeeperTestSuite) TestTradingSessionListRequestReject() {
	type args struct {
		session                            types.Sessions
		msgTradingSessionListRequestReject types.MsgTradingSessionListRequestReject
		tradingSessionList                 types.TradingSessionList
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
			"Valid Trading Session List Request Request",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[1].String(), "sessionID", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"Same account can not used for creating Trading Session List Request and Trading Session List Request Reject with the same TradSesReqID",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[0].String(), "sessionID", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"The user responding is not the recipient of the Trading Session List Request",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[2].String(), "sessionID", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[1].String(), "sessionIDefehfieh", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
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
			"Trading Session List Request is rejected already",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[1].String(), "sessionID", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionListRequestReject: &types.TradingSessionListRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:           "TradSesReqID",
						TradSesStatus:          "TradSesStatus",
						TradSesStatusRejReason: "TradSesStatusRejReason",
						Text:                   "Text",
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
			"The Trading Session List Request is acknowledged already",
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
				msgTradingSessionListRequestReject: *types.NewMsgTradingSessionListRequestReject(suite.address[1].String(), "sessionID", "TradSesReqID", "tradSesStatus", "tradSesStatusRejReason", "text"),
				tradingSessionList: types.TradingSessionList{
					SessionID: "sessionID",
					TradingSessionListRequest: &types.TradingSessionListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:            "TradSesReqID",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SecurityExchange:        "SecurityExchange",
						TradSesMethod:           "TradSesMethod",
						TradSesMode:             "TradSesMode",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradingSessionListResponse: &types.TradingSessionListResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						TradSesReqID:           "TradSesReqID",
						NoTradingSessions:      "NoTradingSessions",
						TradingSessionID:       "TradingSessionID",
						TradingSessionSubID:    "TradingSessionSubID",
						SecurityExchange:       "SecurityExchange",
						TradSesMethod:          "TradSesMethod",
						TradSesMode:            "TradSesMode",
						UnsolicitedIndicator:   "UnsolicitedIndicator",
						TradSesStatus:          "TradSesStatus",
						TradSesStatusRejReason: "TradSesStatusRejReason",
						TradSesStartTime:       "TradSesStartTime",
						TradSesOpenTime:        "TradSesOpenTime",
						TradSesPreCloseTime:    "TradSesPreCloseTime",
						TradSesCloseTime:       "TradSesCloseTime",
						TradSesEndTime:         "TradSesEndTime",
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
			// set trading session list
			suite.fixKeeper.SetTradingSessionList(suite.ctx, tc.args.tradingSessionList.TradingSessionListRequest.TradSesReqID, tc.args.tradingSessionList)
			// call trading session list request reject method
			res, err := suite.msgServer.TradingSessionListRequestReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradingSessionListRequestReject)

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
