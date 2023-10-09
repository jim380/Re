package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestSecurityListRequest() {
	type args struct {
		session                types.Sessions
		msgSecurityListRequest types.MsgSecurityListRequest
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
			"Valid Security List Request",
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
				msgSecurityListRequest: *types.NewMsgSecurityListRequest(suite.address[0].String(), "sessionID", "securityReqID", "securityListRequestType", "noUnderlyings", "noLegs", "currency", "text", "encodedTextLen", "encodedText", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
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
				msgSecurityListRequest: *types.NewMsgSecurityListRequest(suite.address[0].String(), "sessionID", "securityReqID", "securityListRequestType", "noUnderlyings", "noLegs", "currency", "text", "encodedTextLen", "encodedText", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Parties involved in a session are not the ones using the sessionID and are not able to create Security List Request",
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
				msgSecurityListRequest: *types.NewMsgSecurityListRequest(suite.address[2].String(), "sessionID", "securityReqID", "securityListRequestType", "noUnderlyings", "noLegs", "currency", "text", "encodedTextLen", "encodedText", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
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
				msgSecurityListRequest: *types.NewMsgSecurityListRequest(suite.address[0].String(), "sessionIDsnfshfh", "securityReqID", "securityListRequestType", "noUnderlyings", "noLegs", "currency", "text", "encodedTextLen", "encodedText", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
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
			// call  Security List Request method
			res, err := suite.msgServer.SecurityListRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityListRequest)

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

func (suite *KeeperTestSuite) TestSecurityListResponse() {
	type args struct {
		session                 types.Sessions
		msgSecurityListResponse types.MsgSecurityListResponse
		securityList            types.SecurityList
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"Security List Request creator address is same address acknowledging the Security List Response with the same SecurityReqID",
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[0].String(), "sessionID", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The user responding is not the recipient of the Security List Response",
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[2].String(), "sessionID", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The Security List Request is rejected already",
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityListRequestReject: &types.SecurityListRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						SecurityRequestResult:   "SecurityRequestResult",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
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
			"The Security List Response is acknowledged already",
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityListResponse: &types.SecurityListResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:         "SecurityReqID",
						SecurityResponseID:    "SecurityResponseID",
						SecurityRequestResult: "SecurityRequestResult",
						TotNoRelatedSym:       "TotNoRelatedSym",
						LastFragment:          "LastFragment",
						NoRelatedSym:          "NoRelatedSym",
						NoUnderlyings:         "NoUnderlyings",
						Currency:              "Currency",
						NoLegs:                "NoLegs",
						RoundLot:              "RoundLot",
						MinTradeVol:           "MinTradeVol",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						ExpirationCycle:       "ExpirationCycle",
						Text:                  "Text",
						EncodedTextLen:        "EncodedTextLen",
						EncodedText:           "EncodedText",
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
				msgSecurityListResponse: *types.NewMsgSecurityListResponse(suite.address[1].String(), "sessionIDjfbejfbe", "SecurityReqID", "securityResponseID", "securityRequestResult", "totNoRelatedSym", "lastFragment", "noRelatedSym", "noUnderlyings", "currency", "noLegs", "roundLot", "minTradeVol", "tradingSessionID", "tradingSessionSubID", "expirationCycle", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			// set security list
			suite.fixKeeper.SetSecurityList(suite.ctx, tc.args.securityList.SecurityListRequest.SecurityReqID, tc.args.securityList)
			// call Security List Response method
			res, err := suite.msgServer.SecurityListResponse(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityListResponse)

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

func (suite *KeeperTestSuite) TestSecurityListRequestReject() {
	type args struct {
		session                      types.Sessions
		msgSecurityListRequestReject types.MsgSecurityListRequestReject
		securityList                 types.SecurityList
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
			"Valid Security List Request Reject",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The user to acknowledge the Security List Request is not the recipient of the Security List Request",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[2].String(), "sessionID", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"Security List Request Type reason is empty",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"Security List Request creator address is the same address acknowledging the Security List Request Reject with the same SecurityReqID",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[0].String(), "sessionID", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			"The Security List Request is rejected already",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityListRequestReject: &types.SecurityListRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						SecurityRequestResult:   "SecurityRequestResult",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
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
			"The Security List is acknowledged already",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
						SubscriptionRequestType: "SubscriptionRequestType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityListResponse: &types.SecurityListResponse{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:         "SecurityReqID",
						SecurityResponseID:    "SecurityResponseID",
						SecurityRequestResult: "SecurityRequestResult",
						TotNoRelatedSym:       "TotNoRelatedSym",
						LastFragment:          "LastFragment",
						NoRelatedSym:          "NoRelatedSym",
						NoUnderlyings:         "NoUnderlyings",
						Currency:              "Currency",
						NoLegs:                "NoLegs",
						RoundLot:              "RoundLot",
						MinTradeVol:           "MinTradeVol",
						TradingSessionID:      "TradingSessionID",
						TradingSessionSubID:   "TradingSessionSubID",
						ExpirationCycle:       "ExpirationCycle",
						Text:                  "Text",
						EncodedTextLen:        "EncodedTextLen",
						EncodedText:           "EncodedText",
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
				msgSecurityListRequestReject: *types.NewMsgSecurityListRequestReject(suite.address[1].String(), "sessionIDnfjefhj", "SecurityReqID", "securityListRequestType", "securityRequestResult", "text", "encodedTextLen", "encodedText"),
				securityList: types.SecurityList{
					SessionID: "sessionID",
					SecurityListRequest: &types.SecurityListRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:           "SecurityReqID",
						SecurityListRequestType: "SecurityListRequestType",
						NoUnderlyings:           "NoUnderlyings",
						NoLegs:                  "NoLegs",
						Currency:                "Currency",
						Text:                    "Text",
						EncodedTextLen:          "EncodedTextLen",
						EncodedText:             "EncodedText",
						TradingSessionID:        "TradingSessionID",
						TradingSessionSubID:     "TradingSessionSubID",
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
			// set security list
			suite.fixKeeper.SetSecurityList(suite.ctx, tc.args.securityList.SecurityListRequest.SecurityReqID, tc.args.securityList)
			// call security list request reject method
			res, err := suite.msgServer.SecurityListRequestReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityListRequestReject)

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
