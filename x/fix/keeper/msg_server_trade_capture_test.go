package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestTradeCaptureReport() {
	type args struct {
		session               types.Sessions
		msgTradeCaptureReport types.MsgTradeCaptureReport
		ordersExecutionReport types.OrdersExecutionReport
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
			"Valid trade capture report",
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
				msgTradeCaptureReport: *types.NewMsgTradeCaptureReport(suite.address[0].String(), "sessionID", "tradeReportID", "tradeReportTransType", "tradeReportType", "trdType", "trdSubType", "side", "orderQty", "lastQty", "lastPx", "400", "ExecID", "OrderID", "tradeID", "origTradeID", "symbol", "securityID", "securityIDSource", "tradeDate", "transactTime", "settlType", "settlDate"),
				ordersExecutionReport: types.OrdersExecutionReport{
					SessionID: "sessionID",
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
					ClOrdID:      "OrderID",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         1,
					OrderQty:     "OrderQty",
					Price:        "5000",
					TimeInForce:  2,
					LastPx:       1,
					LastQty:      2,
					CumQty:       2,
					AvgPx:        2,
					Text:         "Text",
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
			"Parties involved are not in the session",
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
				msgTradeCaptureReport: *types.NewMsgTradeCaptureReport(suite.address[2].String(), "sessionID", "tradeReportID", "tradeReportTransType", "tradeReportType", "trdType", "trdSubType", "side", "orderQty", "lastQty", "lastPx", "400", "execID", "OrderID", "tradeID", "origTradeID", "symbol", "securityID", "securityIDSource", "tradeDate", "transactTime", "settlType", "settlDate"),
				ordersExecutionReport: types.OrdersExecutionReport{
					SessionID: "sessionID",
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
					ClOrdID:      "ClOrdID",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         1,
					OrderQty:     "OrderQty",
					Price:        "5000",
					TimeInForce:  2,
					LastPx:       1,
					LastQty:      2,
					CumQty:       2,
					AvgPx:        2,
					Text:         "Text",
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
			"Logon status is not logged In",
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
					Status:     "LoggedOut",
					IsAccepted: true,
				},
				msgTradeCaptureReport: *types.NewMsgTradeCaptureReport(suite.address[0].String(), "sessionID", "tradeReportID", "tradeReportTransType", "tradeReportType", "trdType", "trdSubType", "side", "orderQty", "lastQty", "lastPx", "400", "execID", "OrderID", "tradeID", "origTradeID", "symbol", "securityID", "securityIDSource", "tradeDate", "transactTime", "settlType", "settlDate"),
				ordersExecutionReport: types.OrdersExecutionReport{
					SessionID: "sessionID",
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
					ClOrdID:      "ClOrdID",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         1,
					OrderQty:     "OrderQty",
					Price:        "5000",
					TimeInForce:  2,
					LastPx:       1,
					LastQty:      2,
					CumQty:       2,
					AvgPx:        2,
					Text:         "Text",
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
				msgTradeCaptureReport: *types.NewMsgTradeCaptureReport(suite.address[0].String(), "sessionID", "tradeReportID", "tradeReportTransType", "tradeReportType", "trdType", "trdSubType", "side", "orderQty", "lastQty", "lastPx", "400", "execID", "OrderID", "tradeID", "origTradeID", "symbol", "securityID", "securityIDSource", "tradeDate", "transactTime", "settlType", "settlDate"),
				ordersExecutionReport: types.OrdersExecutionReport{
					SessionID: "sessionID",
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
					ClOrdID:      "ClOrdID",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         1,
					OrderQty:     "OrderQty",
					Price:        "5000",
					TimeInForce:  2,
					LastPx:       1,
					LastQty:      2,
					CumQty:       2,
					AvgPx:        2,
					Text:         "Text",
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
			// set order execution report
			suite.fixKeeper.SetOrdersExecutionReport(suite.ctx, tc.args.ordersExecutionReport.ClOrdID, tc.args.ordersExecutionReport)

			// call  trade capture report method
			res, err := suite.msgServer.TradeCaptureReport(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradeCaptureReport)

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

func (suite *KeeperTestSuite) TestTradeCaptureReportAcknowledgement() {
	type args struct {
		session                              types.Sessions
		msgTradeCaptureReportAcknowledgement types.MsgTradeCaptureReportAcknowledgement
		tradeCapture                         types.TradeCapture
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
			"Valid trade capture report acknowledgement",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[1].String(), "sessionID", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			"Same account cannot used for creating trade capture report and trade capture report acknowledgement with the same TradeReportID",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[0].String(), "sessionID", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			"Trade Capture Report is rejected already",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[1].String(), "sessionID", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradeCaptureReportRejection: &types.TradeCaptureReportRejection{
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
						TradeReportID:           "TradeReportID",
						TradeReportRejectReason: 2,
						TradeReportRejectRefID:  "TradeReportRejectRefID",
						Text:                    "Text",
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
			"Trade Capture Report is acknowledged already",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[0].String(), "sessionID", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradeCaptureReportAcknowledgement: &types.TradeCaptureReportAcknowledgement{
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
						TradeReportID:           "TradeReportID",
						TradeID:                 "TradeID",
						SecondaryTradeID:        "SecondaryTradeID",
						TradeReportType:         "TradeReportType",
						TrdType:                 "TrdType",
						TrdSubType:              "TrdSubType",
						ExecType:                "ExecType",
						TradeReportRefID:        "TradeReportRefID",
						SecondaryTradeReportID:  "SecondaryTradeReportID",
						TradeReportStatus:       "TradeReportStatus",
						TradeTransType:          "TradeTransType",
						TradeReportRejectReason: 2,
						Text:                    "Text",
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
			"Account address not involved in session",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[2].String(), "sessionID", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
				msgTradeCaptureReportAcknowledgement: *types.NewMsgTradeCaptureReportAcknowledgement(suite.address[1].String(), "sessionIDhgygyg", "TradeReportID", "TradeID", "secondaryTradeID", "TradeReportType", "TrdType", "TrdSubType", "execType", "TradeReportRefID", "secondaryTradeReportID", "TradeReportStatus", "TradeTransType", 2, "Text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			// set trade capture
			suite.fixKeeper.SetTradeCapture(suite.ctx, tc.args.tradeCapture.TradeCaptureReport.TradeReportID, tc.args.tradeCapture)

			// call trade capture report acknowledgement method
			res, err := suite.msgServer.TradeCaptureReportAcknowledgement(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradeCaptureReportAcknowledgement)

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

func (suite *KeeperTestSuite) TestTradeCaptureReportRejection() {
	type args struct {
		session                        types.Sessions
		msgTradeCaptureReportRejection types.MsgTradeCaptureReportRejection
		tradeCapture                   types.TradeCapture
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
			"Valid trade capture report rejection",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[1].String(), "sessionID", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			"Same account can not used for creating trade capture report and trade capture report rejection with the same TradeReportID",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[0].String(), "sessionID", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			"Trade capture report is rejected already",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[1].String(), "sessionID", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradeCaptureReportRejection: &types.TradeCaptureReportRejection{
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
						TradeReportID:           "TradeReportID",
						TradeReportRejectReason: 2,
						TradeReportRejectRefID:  "TradeReportRejectRefID",
						Text:                    "Text",
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
			"Trade capture report is acknowledged already",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[1].String(), "sessionID", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					TradeCaptureReportAcknowledgement: &types.TradeCaptureReportAcknowledgement{
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
						TradeReportID:           "TradeReportID",
						TradeID:                 "TradeID",
						SecondaryTradeID:        "SecondaryTradeID",
						TradeReportType:         "TradeReportType",
						TrdType:                 "TrdType",
						TrdSubType:              "TrdSubType",
						ExecType:                "ExecType",
						TradeReportRefID:        "TradeReportRefID",
						SecondaryTradeReportID:  "SecondaryTradeReportID",
						TradeReportStatus:       "TradeReportStatus",
						TradeTransType:          "TradeTransType",
						TradeReportRejectReason: 2,
						Text:                    "Text",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[2].String(), "sessionID", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
				msgTradeCaptureReportRejection: *types.NewMsgTradeCaptureReportRejection(suite.address[1].String(), "sessionIDgugug", "TradeReportID", 2, "tradeReportRejectRefID", "text"),
				tradeCapture: types.TradeCapture{
					SessionID: "sessionID",
					TradeCaptureReport: &types.TradeCaptureReport{
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
						TradeReportID:        "TradeReportID",
						TradeReportTransType: "TradeReportTransType",
						TradeReportType:      "TradeReportType",
						TrdType:              "TrdType",
						TrdSubType:           "TrdSubType",
						Side:                 "Side",
						OrderQty:             "OrderQty",
						LastQty:              "LastQty",
						LastPx:               "LastPx",
						GrossTradeAmt:        "GrossTradeAmt",
						ExecID:               "ExecID",
						OrderID:              "OrderID",
						TradeID:              "TradeID",
						OrigTradeID:          "OrigTradeID",
						Symbol:               "Symbol",
						SecurityID:           "SecurityID",
						SecurityIDSource:     "SecurityIDSource",
						TradeDate:            "TradeDate",
						TransactTime:         time.Now().Format("2006-01-02 15:04:05"),
						SettlType:            "SettlType",
						SettlDate:            "SettlDate",
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
			// set trade capture
			suite.fixKeeper.SetTradeCapture(suite.ctx, tc.args.tradeCapture.TradeCaptureReport.TradeReportID, tc.args.tradeCapture)
			// call trade capture report rejection method
			res, err := suite.msgServer.TradeCaptureReportRejection(sdk.WrapSDKContext(suite.ctx), &tc.args.msgTradeCaptureReportRejection)

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
