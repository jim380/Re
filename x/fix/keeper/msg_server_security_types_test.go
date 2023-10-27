package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestSecurityTypesRequest() {
	type args struct {
		session                 types.Sessions
		msgSecurityTypesRequest types.MsgSecurityTypesRequest
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
			"Valid Security Type Request",
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
				msgSecurityTypesRequest: *types.NewMsgSecurityTypesRequest(suite.address[0].String(), "sessionID", "securityReqID", "text", "tradingSessionID", "tradingSessionSubID", "product", "securityType", "securitySubType"),
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			"logon status is not equals to loggedIn",
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
					Status:     "logged_out",
					IsAccepted: true,
				},
				msgSecurityTypesRequest: *types.NewMsgSecurityTypesRequest(suite.address[0].String(), "sessionID", "securityReqID", "text", "tradingSessionID", "tradingSessionSubID", "product", "securityType", "securitySubType"),
			},
			errArgs{
				shouldPass: false,
				contains:   "",
			},
		},
		{
			"Parties involved in a session are not the ones using the sessionID and are not able to create Security Types Request",
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
				msgSecurityTypesRequest: *types.NewMsgSecurityTypesRequest(suite.address[2].String(), "sessionID", "securityReqID", "text", "tradingSessionID", "tradingSessionSubID", "product", "securityType", "securitySubType"),
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
				msgSecurityTypesRequest: *types.NewMsgSecurityTypesRequest(suite.address[0].String(), "sessionIDjjfefje", "securityReqID", "text", "tradingSessionID", "tradingSessionSubID", "product", "securityType", "securitySubType"),
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

			// call  Security Types Request
			res, err := suite.msgServer.SecurityTypesRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityTypesRequest)

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

func (suite *KeeperTestSuite) TestSecurityTypesResponse() {
	type args struct {
		session                  types.Sessions
		msgSecurityTypesResponse types.MsgSecurityTypesResponse
		securityTypes            types.SecurityTypes
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
			"Valid Security Types Response",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"User responding is not the recipient of the Security Types Request",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[2].String(), "sessionID", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							ChainID:      "8",
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"Same account can not used for creating Security Types Request and Security Types Response with the same SecurityReqID",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[0].String(), "sessionID", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"The Security Types Request is rejected already",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityTypesRequestReject: &types.SecurityTypesRequestReject{
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
						SecurityReqID: "SecurityReqID",
						RejectReason:  "RejectReason",
						Text:          "Text",
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
			"The Security Types Request is acknowledged already",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[1].String(), "sessionID", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityTypesResponse: &types.SecurityTypesResponse{
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
						SecurityReqID:           "SecurityReqID",
						SecurityResponseID:      "SecurityResponseID",
						SecurityResponseType:    "SecurityResponseType",
						TotNoSecurityTypes:      "TotNoSecurityTypes",
						LastFragment:            "LastFragment",
						NoSecurityTypes:         "NoSecurityTypes",
						SecurityType:            "SecurityType",
						SecuritySubType:         "SecuritySubType",
						Product:                 "Product",
						CFICode:                 "CFICode",
						Text:                    "Text",
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
				msgSecurityTypesResponse: *types.NewMsgSecurityTypesResponse(suite.address[1].String(), "sessionIDkhhuyfftd", "SecurityReqID", "SecurityResponseID", "SecurityResponseType", "totNoSecurityTypes", "lastFragment", "noSecurityTypes", "securityType", "securitySubType", "product", "cFICode", "text", "tradingSessionID", "tradingSessionSubID", "subscriptionRequestType"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			// set security types
			suite.fixKeeper.SetSecurityTypes(suite.ctx, tc.args.securityTypes.SecurityTypesRequest.SecurityReqID, tc.args.securityTypes)

			// call security types response method
			res, err := suite.msgServer.SecurityTypesResponse(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityTypesResponse)

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

func (suite *KeeperTestSuite) TestSecurityTypesRequestReject() {
	type args struct {
		session                       types.Sessions
		msgSecurityTypesRequestReject types.MsgSecurityTypesRequestReject
		securityTypes                 types.SecurityTypes
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
			"Valid Security Types Request Reject",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"User responding is not the recipient of the Security Types Request",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[2].String(), "sessionID", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"Same account can not used for creating Security Types Request and Security Types Response with the same SecurityReqID",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[0].String(), "sessionID", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			"The Security Types Request is rejected already",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityTypesRequestReject: &types.SecurityTypesRequestReject{
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
						SecurityReqID: "SecurityReqID",
						RejectReason:  "RejectReason",
						Text:          "Text",
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
			"The Security Types Request is acknowledged already",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[1].String(), "sessionID", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					SecurityTypesResponse: &types.SecurityTypesResponse{
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
						SecurityReqID:           "SecurityReqID",
						SecurityResponseID:      "SecurityResponseID",
						SecurityResponseType:    "SecurityResponseType",
						TotNoSecurityTypes:      "TotNoSecurityTypes",
						LastFragment:            "LastFragment",
						NoSecurityTypes:         "NoSecurityTypes",
						SecurityType:            "SecurityType",
						SecuritySubType:         "SecuritySubType",
						Product:                 "Product",
						CFICode:                 "CFICode",
						Text:                    "Text",
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
				msgSecurityTypesRequestReject: *types.NewMsgSecurityTypesRequestReject(suite.address[1].String(), "sessionIDgggujj", "SecurityReqID", "rejectReason", "text"),
				securityTypes: types.SecurityTypes{
					SessionID: "sessionID",
					SecurityTypesRequest: &types.SecurityTypesRequest{
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
						SecurityReqID:       "SecurityReqID",
						Text:                "Text",
						TradingSessionID:    "TradingSessionID",
						TradingSessionSubID: "TradingSessionSubID",
						Product:             "Product",
						SecurityType:        "SecurityType",
						SecuritySubType:     "SecuritySubType",
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
			// set security types
			suite.fixKeeper.SetSecurityTypes(suite.ctx, tc.args.securityTypes.SecurityTypesRequest.SecurityReqID, tc.args.securityTypes)

			// call security types request reject method
			res, err := suite.msgServer.SecurityTypesRequestReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgSecurityTypesRequestReject)

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
