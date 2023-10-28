package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	cosmostxs "github.com/jim380/Re/utils/cosmos_txs"
	cosmosTxTypes "github.com/jim380/Re/utils/cosmos_txs/types"
	"github.com/jim380/Re/utils/helpers"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

// CosmoshubTxs generate the logon, new single order, execution report and
// trade capture report for transactions on the Cosmos Hub
func (k msgServer) CosmoshubTxs(goCtx context.Context, msg *types.MsgCosmoshubTxs) (*types.MsgCosmoshubTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}
	if msg.Creator != msg.Address {
		return nil, sdkerrors.Wrapf(types.ErrWrongAddress, "Address: %s", msg.Address)
	}

	// set cosmoshub txs, map to the fix module
	// update fix module from oracle
	// start refetching data every 5 seconds and return 20 latest Txs
	go cosmostxs.RefetchCosmosTxsDataPeriodically(constants.CosmosHubCacheKey)
	transactions, err := cosmostxs.GetCachedCosmosTransactions(constants.CosmosHubCacheKey)
	if err != nil {
		return nil, fmt.Errorf("failed to return queried chain transactions data: %w", err)
	}
	for _, tx := range transactions.Transactions {
		for _, message := range tx.Messages {
			// type of Txs message = Send
			switch helpers.AbbrTxMessage(message.Type) {
			case "Send":
				// set sessions for parties existing in the SEND transactions
				session := fixTypes.Sessions{
					SessionID: tx.TxHash,
					LogonInitiator: &fixTypes.LogonInitiator{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.FromAddress,
							TargetCompID: message.ToAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					LogonAcceptor: &fixTypes.LogonAcceptor{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.ToAddress,
							TargetCompID: message.FromAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					Status:     "loggedIn",
					IsAccepted: true,
				}
				// set new session to store
				k.fixKeeper.SetSessions(ctx, tx.TxHash, session)

				// unmarshal Amount manually before using it
				var coins []cosmosTxTypes.Coin
				err := json.Unmarshal(message.Amount, &coins)
				if err != nil {
					return nil, fmt.Errorf("failed to unmarshal Amount as an array of Coin objects. Details: %w", err)
				}

				// set New Single Order
				// set Orders Execution Report
				// set Orders Rejection
				// set Trade Capture
				// check that Txs was successful
				if tx.Result != 0 {
					// orders here were rejected
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.FromAddress,
							TargetCompID: message.ToAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       coins[0].Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        coins[0].Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// when a send transaction fails, order cancel reject is generated
					ordersCancelReject := fixTypes.OrdersCancelReject{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "9",
							SenderCompID: message.ToAddress,
							TargetCompID: message.FromAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						OrderID:      tx.TxHash,
						OrigClOrdID:  tx.TxHash,
						ClOrdID:      tx.TxHash,
						CxlRejReason: 4,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersCancelReject(ctx, tx.TxHash, ordersCancelReject)

					// set Trade Capture Report Reject when Txs fails
					tradeCaptureReportReject := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.FromAddress,
								TargetCompID: message.ToAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Send",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        coins[0].Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               coins[0].Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportRejection: &fixTypes.TradeCaptureReportRejection{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "j",
								SenderCompID: message.ToAddress,
								TargetCompID: message.FromAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:           tx.TxHash,
							TradeReportRejectReason: 0,
							TradeReportRejectRefID:  tx.TxHash,
							Text:                    tx.Memo,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCaptureReportReject)

				} else {

					// set successful orders
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.FromAddress,
							TargetCompID: message.ToAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       coins[0].Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        coins[0].Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// set execution report
					orderExecutionReport := &fixTypes.OrdersExecutionReport{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "8",
							SenderCompID: message.ToAddress,
							TargetCompID: message.FromAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						OrderID:      tx.TxHash,
						ExecID:       tx.TxHash,
						OrdStatus:    "New",
						ExecType:     "New",
						Symbol:       coins[0].Denom,
						Side:         2,
						OrderQty:     "",
						Price:        coins[0].Amount,
						TimeInForce:  1,
						LastPx:       0,
						LastQty:      0,
						LeavesQty:    0,
						CumQty:       0,
						AvgPx:        0,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersExecutionReport(ctx, tx.TxHash, *orderExecutionReport)

					// set Trade Capture Report if Txs was successful
					tradeCapture := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.FromAddress,
								TargetCompID: message.ToAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Send",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        coins[0].Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               coins[0].Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportAcknowledgement: &fixTypes.TradeCaptureReportAcknowledgement{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AR",
								SenderCompID: message.ToAddress,
								TargetCompID: message.FromAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:          tx.TxHash,
							TradeID:                tx.TxHash,
							SecondaryTradeID:       "",
							TradeReportType:        "real-time",
							TrdType:                "Block Trade",
							ExecType:               "New",
							TradeReportRefID:       tx.TxHash,
							SecondaryTradeReportID: "",
							TradeReportStatus:      "Accepted",
							TradeTransType:         "Send",
							Text:                   tx.Memo,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCapture)
				}

			case "Transfer":
				// set sessions for parties existing in the Transfer transactions
				session := fixTypes.Sessions{
					SessionID: tx.TxHash,
					LogonInitiator: &fixTypes.LogonInitiator{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.Sender,
							TargetCompID: message.Receiver,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					LogonAcceptor: &fixTypes.LogonAcceptor{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.Receiver,
							TargetCompID: message.Sender,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					Status:     "loggedIn",
					IsAccepted: true,
				}
				// set new session to store
				k.fixKeeper.SetSessions(ctx, tx.TxHash, session)

				// set New Single Order
				// set Orders Execution Report
				// set Orders Rejection
				// set Trade Capture
				// check that Txs was success
				if tx.Result != 0 {
					// orders here were rejected
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.Sender,
							TargetCompID: message.Receiver,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       message.Token.Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        message.Token.Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// when a Transfer transaction fails, order cancel reject is generated
					ordersCancelReject := fixTypes.OrdersCancelReject{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "9",
							SenderCompID: message.Receiver,
							TargetCompID: message.Sender,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						OrderID:      tx.TxHash,
						OrigClOrdID:  tx.TxHash,
						ClOrdID:      tx.TxHash,
						CxlRejReason: 4,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersCancelReject(ctx, tx.TxHash, ordersCancelReject)

					// set Trade Capture Report Reject when Transfer Txs fails
					tradeCaptureReportReject := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.Sender,
								TargetCompID: message.Receiver,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Transfer",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        message.Token.Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               message.Token.Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportRejection: &fixTypes.TradeCaptureReportRejection{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "j",
								SenderCompID: message.Receiver,
								TargetCompID: message.Sender,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:           tx.TxHash,
							TradeReportRejectReason: 0,
							TradeReportRejectRefID:  tx.TxHash,
							Text:                    tx.Memo + tx.RawLog,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCaptureReportReject)

				} else {

					// set successful orders
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.Sender,
							TargetCompID: message.Receiver,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       message.Token.Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        message.Token.Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// set execution report
					orderExecutionReport := &fixTypes.OrdersExecutionReport{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "8",
							SenderCompID: message.Receiver,
							TargetCompID: message.Sender,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						OrderID:      tx.TxHash,
						ExecID:       tx.TxHash,
						OrdStatus:    "New",
						ExecType:     "New",
						Symbol:       message.Token.Denom,
						Side:         2,
						OrderQty:     "",
						Price:        message.Token.Amount,
						TimeInForce:  1,
						LastPx:       0,
						LastQty:      0,
						LeavesQty:    0,
						CumQty:       0,
						AvgPx:        0,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersExecutionReport(ctx, tx.TxHash, *orderExecutionReport)

					// set Trade Capture Report if Txs was successful
					tradeCapture := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.Sender,
								TargetCompID: message.Receiver,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Transfer",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        message.Token.Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               message.Token.Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportAcknowledgement: &fixTypes.TradeCaptureReportAcknowledgement{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AR",
								SenderCompID: message.Receiver,
								TargetCompID: message.Sender,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:          tx.TxHash,
							TradeID:                tx.TxHash,
							SecondaryTradeID:       "",
							TradeReportType:        "real-time",
							TrdType:                "Block Trade",
							ExecType:               "New",
							TradeReportRefID:       tx.TxHash,
							SecondaryTradeReportID: "",
							TradeReportStatus:      "Accepted",
							TradeTransType:         "Transfer",
							Text:                   tx.Memo,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCapture)
				}
			case "Delegate":
				// set sessions for parties existing in the Delegate transactions
				session := fixTypes.Sessions{
					SessionID: tx.TxHash,
					LogonInitiator: &fixTypes.LogonInitiator{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.DelegatorAddress,
							TargetCompID: message.ValidatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					LogonAcceptor: &fixTypes.LogonAcceptor{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.ValidatorAddress,
							TargetCompID: message.DelegatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					},
					Status:     "loggedIn",
					IsAccepted: true,
				}
				// set new session to store
				k.fixKeeper.SetSessions(ctx, tx.TxHash, session)

				// unmarshal Amount manually before using it
				var coins cosmosTxTypes.Coin
				err := json.Unmarshal(message.Amount, &coins)
				if err != nil {
					return nil, fmt.Errorf("failed to unmarshal Amount as a Coin object. Details: %w", err)
				}

				// set New Single Order
				// set Orders Execution Report
				// set Orders Rejection
				// set Trade Capture
				// check that Txs was success
				if tx.Result != 0 {
					// orders here were rejected
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.DelegatorAddress,
							TargetCompID: message.ValidatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       coins.Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        coins.Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// when a send transaction fails, order cancel reject is generated
					ordersCancelReject := fixTypes.OrdersCancelReject{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "9",
							SenderCompID: message.ValidatorAddress,
							TargetCompID: message.DelegatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						OrderID:      tx.TxHash,
						OrigClOrdID:  tx.TxHash,
						ClOrdID:      tx.TxHash,
						CxlRejReason: 4,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersCancelReject(ctx, tx.TxHash, ordersCancelReject)

					// set Trade Capture Report Reject when Txs fails
					tradeCaptureReportReject := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.DelegatorAddress,
								TargetCompID: message.ValidatorAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Delegate",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        coins.Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               coins.Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportRejection: &fixTypes.TradeCaptureReportRejection{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "j",
								SenderCompID: message.ValidatorAddress,
								TargetCompID: message.DelegatorAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:           tx.TxHash,
							TradeReportRejectReason: 0,
							TradeReportRejectRefID:  tx.TxHash,
							Text:                    tx.Memo + tx.RawLog,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCaptureReportReject)

				} else {

					// set successful orders
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "D",
							SenderCompID: message.DelegatorAddress,
							TargetCompID: message.ValidatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						Symbol:       coins.Denom,
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        coins.Amount,
						TimeInForce:  1,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// set execution report
					orderExecutionReport := &fixTypes.OrdersExecutionReport{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString:  "FIX4.4",
							MsgType:      "8",
							SenderCompID: message.ValidatorAddress,
							TargetCompID: message.DelegatorAddress,
							MsgSeqNum:    tx.Height,
							SendingTime:  tx.Time,
							ChainID:      constants.CosmosChainID,
						},
						ClOrdID:      tx.TxHash,
						OrderID:      tx.TxHash,
						ExecID:       tx.TxHash,
						OrdStatus:    "New",
						ExecType:     "New",
						Symbol:       coins.Denom,
						Side:         2,
						OrderQty:     "",
						Price:        coins.Amount,
						TimeInForce:  1,
						LastPx:       0,
						LastQty:      0,
						LeavesQty:    0,
						CumQty:       0,
						AvgPx:        0,
						Text:         tx.Memo,
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersExecutionReport(ctx, tx.TxHash, *orderExecutionReport)

					// set Trade Capture Report if Txs was successful
					tradeCapture := &fixTypes.TradeCapture{
						SessionID: tx.TxHash,
						TradeCaptureReport: &fixTypes.TradeCaptureReport{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AE",
								SenderCompID: message.DelegatorAddress,
								TargetCompID: message.ValidatorAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:        tx.TxHash,
							TradeReportTransType: "New",
							TradeReportType:      "real-time",
							TrdType:              "Block Trade",
							TrdSubType:           "Delegate",
							Side:                 "2",
							OrderQty:             "",
							LastQty:              "",
							LastPx:               "",
							GrossTradeAmt:        coins.Amount,
							ExecID:               tx.TxHash,
							OrderID:              tx.TxHash,
							TradeID:              tx.TxHash,
							Symbol:               coins.Denom,
							TransactTime:         tx.Time,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						TradeCaptureReportAcknowledgement: &fixTypes.TradeCaptureReportAcknowledgement{
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "AR",
								SenderCompID: message.ValidatorAddress,
								TargetCompID: message.DelegatorAddress,
								MsgSeqNum:    tx.Height,
								SendingTime:  tx.Time,
								ChainID:      constants.CosmosChainID,
							},
							TradeReportID:          tx.TxHash,
							TradeID:                tx.TxHash,
							SecondaryTradeID:       "",
							TradeReportType:        "real-time",
							TrdType:                "Block Trade",
							ExecType:               "New",
							TradeReportRefID:       tx.TxHash,
							SecondaryTradeReportID: "",
							TradeReportStatus:      "Accepted",
							TradeTransType:         "Delegate",
							Text:                   tx.Memo,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
					}
					k.fixKeeper.SetTradeCapture(ctx, tx.TxHash, *tradeCapture)
				}

			}
		}
	}

	oracle := types.MultiChainTxOracle{
		OracleID:  msg.OracleID,
		Address:   msg.Address,
		Timestamp: constants.CreatedAt,
	}

	// set oracle to store
	k.SetMultiChainTxOracle(ctx, msg.OracleID, oracle)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgCosmoshubTxsResponse{}, err
}
