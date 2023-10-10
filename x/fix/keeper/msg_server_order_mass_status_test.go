package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestOrderMassStatusRequest() {
	type args struct {
		session                   types.Sessions
		msgOrderMassStatusRequest types.MsgOrderMassStatusRequest
		order                     types.Orders
		ordersExecutionReport     types.OrdersExecutionReport
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
			"Valid Order Mass Status Request",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionID", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionID", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			"The parties involved in a session are the ones using the sessionID and are able to send Order Mass Status Request",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[2].String(), "sessionID", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			"The owner of the order does not matches the creator of order mass status request",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[1].String(), "sessionID", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			"The order mass status request symbol is not the same symbol as the order",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionID", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "symbol", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			"The mass status request id is empty",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionID", "", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			"The mass status request type field is empty",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionID", "massStatusReqID", "", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
				msgOrderMassStatusRequest: *types.NewMsgOrderMassStatusRequest(suite.address[0].String(), "sessionIDgugugg", "massStatusReqID", "massStatusReqType", "ClOrdID", "account", "TEXT_SYMBOL", "securityID", "tradingSessionID"),
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
					ClOrdID:      "ClOrdID",
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
					},
					ClOrdID:      "massStatusReqType",
					OrderID:      "OrderID",
					ExecID:       "ExecID",
					OrdStatus:    "OrdStatus",
					ExecType:     "ExecType",
					Symbol:       "Symbol",
					Side:         3,
					OrderQty:     "OrderQty",
					Price:        "Price",
					TimeInForce:  3,
					LastPx:       2,
					LastQty:      2,
					LeavesQty:    1,
					CumQty:       5,
					AvgPx:        6,
					Text:         "Text",
					TransactTime: "TransactTime",
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
			// set order execution report
			suite.fixKeeper.SetOrdersExecutionReport(suite.ctx, tc.args.ordersExecutionReport.ClOrdID, tc.args.ordersExecutionReport)
			// call  Order Mass Status Request method
			res, err := suite.msgServer.OrderMassStatusRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgOrderMassStatusRequest)

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

func (suite *KeeperTestSuite) TestOrderMassStatusReport() {
	type args struct {
		session                  types.Sessions
		msgOrderMassStatusReport types.MsgOrderMassStatusReport
		orderMassStatus          types.OrderMassStatus
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
			"Valid Order Mass Status Report",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[1].String(), "sessionID", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The account that created order mass status request is not allowed to respond to order mass status report",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[0].String(), "sessionID", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The parties involved in a session are not the ones using the sessionID and are not able to send Order Mass Status Report",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[2].String(), "sessionID", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The Order Mass Status Request is rejected already",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[1].String(), "sessionID", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					OrderMassStatusRequestReject: &types.OrderMassStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						RefSeqID: "RefSeqID",
						RejCode:  "RejCode",
						Text:     "Text",
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
			"The Order Mass Status Request is acknowledged already",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[1].String(), "sessionID", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					OrderMassStatusReport: &types.OrderMassStatusReport{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						ClOrdID:          "ClOrdID",
						MassStatusReqID:  "MassStatusReqID",
						Account:          "Account",
						Symbol:           "Symbol",
						SecurityID:       "SecurityID",
						TradingSessionID: "TradingSessionID",
						OrdStatus:        "OrdStatus",
						ExecType:         "ExecType",
						OrdQty:           "OrdQty",
						LastPx:           "LastPx",
						CumQty:           "CumQty",
						AvgPx:            "AvgPx",
						LeavesQty:        "LeavesQty",
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
				msgOrderMassStatusReport: *types.NewMsgOrderMassStatusReport(suite.address[1].String(), "sessionIDjvjjgjgcgc", "ClOrdID", "MassStatusReqID", "account", "Symbol", "securityID", "tradingSessionID", "ordStatus", "execType", "ordQty", "lastPx", "cumQty", "avgPx", "leavesQty"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "MassStatusReqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			// set order mass status
			suite.fixKeeper.SetOrderMassStatus(suite.ctx, tc.args.orderMassStatus.OrderMassStatusRequest.MassStatusReqID, tc.args.orderMassStatus)
			// call order mass status report method
			res, err := suite.msgServer.OrderMassStatusReport(sdk.WrapSDKContext(suite.ctx), &tc.args.msgOrderMassStatusReport)

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

func (suite *KeeperTestSuite) TestOrderMassStatusRequestReject() {
	type args struct {
		session                         types.Sessions
		msgOrderMassStatusRequestReject types.MsgOrderMassStatusRequestReject
		orderMassStatus                 types.OrderMassStatus
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
			"Valid Order Mass Status Request Reject",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[1].String(), "sessionID", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The account that created order mass status request is not allowed to respond to order mass status request reject",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[0].String(), "sessionID", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The user responding is not the recipient of the Order mass Status Request",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[2].String(), "sessionID", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			"The Order Mass Status Request is rejected already",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[1].String(), "sessionID", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					OrderMassStatusRequestReject: &types.OrderMassStatusRequestReject{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						RefSeqID: "RefSeqID",
						RejCode:  "RejCode",
						Text:     "Text",
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
			"The Order Mass Status Request is acknowledged already",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[1].String(), "sessionID", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
						Trailer: &types.Trailer{
							CheckSum: 10,
						},
					},
					OrderMassStatusReport: &types.OrderMassStatusReport{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						ClOrdID:          "ClOrdID",
						MassStatusReqID:  "MassStatusReqID",
						Account:          "Account",
						Symbol:           "Symbol",
						SecurityID:       "SecurityID",
						TradingSessionID: "TradingSessionID",
						OrdStatus:        "OrdStatus",
						ExecType:         "ExecType",
						OrdQty:           "OrdQty",
						LastPx:           "LastPx",
						CumQty:           "CumQty",
						AvgPx:            "AvgPx",
						LeavesQty:        "LeavesQty",
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
				msgOrderMassStatusRequestReject: *types.NewMsgOrderMassStatusRequestReject(suite.address[1].String(), "sessionIDhvhgygg", "RefSeqID", "rejCode", "text"),
				orderMassStatus: types.OrderMassStatus{
					SessionID: "sessionID",
					OrderMassStatusRequest: &types.OrderMassStatusRequest{
						Header: &types.Header{
							BeginString:  "FIX4.4",
							BodyLength:   10,
							MsgType:      "A",
							SenderCompID: suite.address[0].String(),
							TargetCompID: suite.address[1].String(),
							MsgSeqNum:    1,
							SendingTime:  time.Now().Format("2006-01-02 15:04:05"),
						},
						MassStatusReqID:   "RefSeqID",
						MassStatusReqType: "MassStatusReqType",
						ClOrdID:           "ClOrdID",
						Account:           "Account",
						Symbol:            "Symbol",
						SecurityID:        "SecurityID",
						TradingSessionID:  "TradingSessionID",
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
			// set order mass status
			suite.fixKeeper.SetOrderMassStatus(suite.ctx, tc.args.orderMassStatus.OrderMassStatusRequest.MassStatusReqID, tc.args.orderMassStatus)
			// call order mass status request reject method
			res, err := suite.msgServer.OrderMassStatusRequestReject(sdk.WrapSDKContext(suite.ctx), &tc.args.msgOrderMassStatusRequestReject)

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
