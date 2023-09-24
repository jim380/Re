package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/x/fix/types"
)

func (suite *KeeperTestSuite) TestNewOrderSingle() {
	type args struct {
		session           types.Sessions
		msgNewOrderSingle types.MsgNewOrderSingle
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
			"Valid New Single Order",
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[0].String(), "sessionID", "clOrdID", "TEXT_SYMBOL", 1, "100", 1, "500", 4, "New Order Single"),
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[2].String(), "sessionID", "clOrdID", "TEXT_SYMBOL", 1, "100", 1, "500", 4, "New Order Single"),
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
				msgNewOrderSingle: *types.NewMsgNewOrderSingle(suite.address[2].String(), "sessionIDxyzs", "clOrdID", "TEXT_SYMBOL", 1, "100", 1, "500", 4, "New Order Single"),
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

			// call  New Single Order method
			res, err := suite.msgServer.NewOrderSingle(sdk.WrapSDKContext(suite.ctx), &tc.args.msgNewOrderSingle)

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

func (suite *KeeperTestSuite) TestOrderCancelRequest() {
	type args struct {
		session               types.Sessions
		msgOrderCancelRequest types.MsgOrderCancelRequest
		order                 types.Orders
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
			"Valid Order Cancel Request",
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
				msgOrderCancelRequest: *types.NewMsgOrderCancelRequest(suite.address[0].String(), "sessionID", "OrgclOrdID", "clOrdID"),
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
			"Account address is not the creator of the order",
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
				msgOrderCancelRequest: *types.NewMsgOrderCancelRequest(suite.address[2].String(), "sessionID", "OrgclOrdID", "clOrdID"),
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
				msgOrderCancelRequest: *types.NewMsgOrderCancelRequest(suite.address[0].String(), "sessionIDqwerrty", "OrgclOrdID", "clOrdID"),
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
			// set order cancel request
			suite.fixKeeper.SetOrders(suite.ctx, tc.args.order.ClOrdID, tc.args.order)

			// call Order cancel request method
			res, err := suite.msgServer.OrderCancelRequest(sdk.WrapSDKContext(suite.ctx), &tc.args.msgOrderCancelRequest)

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
