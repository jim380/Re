package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestSecurityStatusRequest() {
	type args struct {
		session                  types.Sessions
		msgSecurityStatusRequest types.MsgSecurityStatusRequest
		securityStatus           types.SecurityStatus
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
			"Valid Security Status Request",
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
				msgSecurityStatusRequest: *types.NewMsgSecurityStatusRequest(suite.address[0].String(), "sessionID", "SecurityStatusReqID", "instrument", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "subscriptionRequestType", "tradingSessionID", "tradingSessionSubID"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusResponse: &types.SecurityStatusResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:   "SecurityStatusReqID",
						NoUnderlyings:         "NoUnderlyings",
						UnderlyingInstrument:  "UnderlyingInstrument",
						NoLegs:                "NoLegs",
						InstrumentLeg:         "InstrumentLeg",
						Currency:              "Currency",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						UnsolicitedIndicator:  "UnsolicitedIndicator",
						SecurityTradingStatus: "SecurityTradingStatus",
						FinancialStatus:       "FinancialStatus",
						CorporateAction:       "CorporateAction",
						HaltReason:            "HaltReason",
						InViewOfCommon:        "InViewOfCommon",
						DueToRelated:          "DueToRelated",
						BuyVolume:             "BuyVolume",
						SellVolume:            "SellVolume",
						HighPx:                "HighPx",
						LowPx:                 "LowPx",
						LastPx:                "LastPx",
						TransactTime:          "TransactTime",
						Adjustment:            "Adjustment",
						Text:                  "Text",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:  "SecurityStatusReqID",
						SecurityRejectReason: "SecurityRejectReason",
						Text:                 "Text",
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
				msgSecurityStatusRequest: *types.NewMsgSecurityStatusRequest(suite.address[0].String(), "sessionID", "SecurityStatusReqID", "instrument", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "subscriptionRequestType", "tradingSessionID", "tradingSessionSubID"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusResponse: &types.SecurityStatusResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:   "SecurityStatusReqID",
						NoUnderlyings:         "NoUnderlyings",
						UnderlyingInstrument:  "UnderlyingInstrument",
						NoLegs:                "NoLegs",
						InstrumentLeg:         "InstrumentLeg",
						Currency:              "Currency",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						UnsolicitedIndicator:  "UnsolicitedIndicator",
						SecurityTradingStatus: "SecurityTradingStatus",
						FinancialStatus:       "FinancialStatus",
						CorporateAction:       "CorporateAction",
						HaltReason:            "HaltReason",
						InViewOfCommon:        "InViewOfCommon",
						DueToRelated:          "DueToRelated",
						BuyVolume:             "BuyVolume",
						SellVolume:            "SellVolume",
						HighPx:                "HighPx",
						LowPx:                 "LowPx",
						LastPx:                "LastPx",
						TransactTime:          "TransactTime",
						Adjustment:            "Adjustment",
						Text:                  "Text",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:  "SecurityStatusReqID",
						SecurityRejectReason: "SecurityRejectReason",
						Text:                 "Text",
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
			"Parties involved in a session are not the ones using the sessionID and are not able to create Security Status Request",
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
				msgSecurityStatusRequest: *types.NewMsgSecurityStatusRequest(suite.address[2].String(), "sessionID", "SecurityStatusReqID", "instrument", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "subscriptionRequestType", "tradingSessionID", "tradingSessionSubID"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusResponse: &types.SecurityStatusResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:   "SecurityStatusReqID",
						NoUnderlyings:         "NoUnderlyings",
						UnderlyingInstrument:  "UnderlyingInstrument",
						NoLegs:                "NoLegs",
						InstrumentLeg:         "InstrumentLeg",
						Currency:              "Currency",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						UnsolicitedIndicator:  "UnsolicitedIndicator",
						SecurityTradingStatus: "SecurityTradingStatus",
						FinancialStatus:       "FinancialStatus",
						CorporateAction:       "CorporateAction",
						HaltReason:            "HaltReason",
						InViewOfCommon:        "InViewOfCommon",
						DueToRelated:          "DueToRelated",
						BuyVolume:             "BuyVolume",
						SellVolume:            "SellVolume",
						HighPx:                "HighPx",
						LowPx:                 "LowPx",
						LastPx:                "LastPx",
						TransactTime:          "TransactTime",
						Adjustment:            "Adjustment",
						Text:                  "Text",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:  "SecurityStatusReqID",
						SecurityRejectReason: "SecurityRejectReason",
						Text:                 "Text",
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
				msgSecurityStatusRequest: *types.NewMsgSecurityStatusRequest(suite.address[0].String(), "sessionIDdsdhgshd", "SecurityStatusReqID", "instrument", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "subscriptionRequestType", "tradingSessionID", "tradingSessionSubID"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusResponse: &types.SecurityStatusResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:   "SecurityStatusReqID",
						NoUnderlyings:         "NoUnderlyings",
						UnderlyingInstrument:  "UnderlyingInstrument",
						NoLegs:                "NoLegs",
						InstrumentLeg:         "InstrumentLeg",
						Currency:              "Currency",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						UnsolicitedIndicator:  "UnsolicitedIndicator",
						SecurityTradingStatus: "SecurityTradingStatus",
						FinancialStatus:       "FinancialStatus",
						CorporateAction:       "CorporateAction",
						HaltReason:            "HaltReason",
						InViewOfCommon:        "InViewOfCommon",
						DueToRelated:          "DueToRelated",
						BuyVolume:             "BuyVolume",
						SellVolume:            "SellVolume",
						HighPx:                "HighPx",
						LowPx:                 "LowPx",
						LastPx:                "LastPx",
						TransactTime:          "TransactTime",
						Adjustment:            "Adjustment",
						Text:                  "Text",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:  "SecurityStatusReqID",
						SecurityRejectReason: "SecurityRejectReason",
						Text:                 "Text",
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
			// set security status
			suite.fixKeeper.SetSecurityStatus(suite.ctx, tc.args.securityStatus.SecurityStatusRequest.SecurityStatusReqID, tc.args.securityStatus)
			// call  Security Status Request method
			res, err := suite.msgServer.SecurityStatusRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityStatusRequest)

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

func (suite *KeeperTestSuite) TestSecurityStatusResponse() {
	type args struct {
		session                   types.Sessions
		msgSecurityStatusResponse types.MsgSecurityStatusResponse
		securityStatus            types.SecurityStatus
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
			"Valid Security Status Response",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[1].String(), "sessionID", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"Same account can not used for creating Security Status Request and Security Status Response with the same SecurityStatusReqID",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[0].String(), "sessionID", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The user responding is not the recipient of the Security Status Request",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[2].String(), "sessionID", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The Security Status Request is rejected already",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[1].String(), "sessionID", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusRequestReject: &types.SecurityStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:  "SecurityStatusReqID",
						SecurityRejectReason: "SecurityRejectReason",
						Text:                 "Text",
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
			"The Security Status Request is acknowledged already",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[1].String(), "sessionID", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityStatusResponse: &types.SecurityStatusResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:   "SecurityStatusReqID",
						NoUnderlyings:         "NoUnderlyings",
						UnderlyingInstrument:  "UnderlyingInstrument",
						NoLegs:                "NoLegs",
						InstrumentLeg:         "InstrumentLeg",
						Currency:              "Currency",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						UnsolicitedIndicator:  "UnsolicitedIndicator",
						SecurityTradingStatus: "SecurityTradingStatus",
						FinancialStatus:       "FinancialStatus",
						CorporateAction:       "CorporateAction",
						HaltReason:            "HaltReason",
						InViewOfCommon:        "InViewOfCommon",
						DueToRelated:          "DueToRelated",
						BuyVolume:             "BuyVolume",
						SellVolume:            "SellVolume",
						HighPx:                "HighPx",
						LowPx:                 "LowPx",
						LastPx:                "LastPx",
						TransactTime:          "TransactTime",
						Adjustment:            "Adjustment",
						Text:                  "Text",
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
				msgSecurityStatusResponse: *types.NewMsgSecurityStatusResponse(suite.address[1].String(), "sessionIDbchvhs", "SecurityStatusReqID", "noUnderlyings", "underlyingInstrument", "noLegs", "instrumentLeg", "currency", "tradingSessionID", "tradingSessionSubID", "unsolicitedIndicator", "securityTradingStatus", "financialStatusg", "corporateAction", "haltReason", "inViewOfCommon", "dueToRelated", "buyVolume", "sellVolume", "highPx", "lowPx", "lastPx", "transactTime", "adjustment", "text"),
				securityStatus: types.SecurityStatus{
					SessionID: "sessionID",
					SecurityStatusRequest: &types.SecurityStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityStatusReqID:     "SecurityStatusReqID",
						Instrument:              "Instrument",
						NoUnderlyings:           "NoUnderlyings",
						UnderlyingInstrument:    "UnderlyingInstrument",
						NoLegs:                  "NoLegs",
						InstrumentLeg:           "InstrumentLeg",
						Currency:                "Currency",
						SubscriptionRequestType: "SubscriptionRequestType",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			// set security status
			suite.fixKeeper.SetSecurityStatus(suite.ctx, tc.args.securityStatus.SecurityStatusRequest.SecurityStatusReqID, tc.args.securityStatus)
			// call Security Status Response method
			res, err := suite.msgServer.SecurityStatusResponse(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityStatusResponse)

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
}*/
