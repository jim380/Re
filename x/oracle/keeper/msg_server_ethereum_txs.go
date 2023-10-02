package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	"github.com/jim380/Re/utils/ethereum"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

func (k msgServer) EthereumTxs(goCtx context.Context, msg *types.MsgEthereumTxs) (*types.MsgEthereumTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}
	if msg.Creator != msg.Address {
		return nil, sdkerrors.Wrapf(types.ErrWrongAddress, "Address: %s", msg.Address)
	}

	// set ethereum txs, map to the fix module
	// update fix module from oracle
	// start refetching data every 5 seconds and return latest ethereum Txs
	go ethereum.RefetchEthereumTxsDataPeriodically(constants.EthereumCacheKey)
	transactions, err := ethereum.GetCachedEthereumTransactions(constants.BitcoinCacheKey)
	if err != nil {
		return nil, fmt.Errorf("failed to return queried chain transactions data: %w", err)
	}
	for _, tx := range transactions.Transactions {

		// set sessions for parties existing in ethereum transactions
		session := fixTypes.Sessions{
			SessionID: tx.Hash,
			LogonInitiator: &fixTypes.LogonInitiator{
				Header: &fixTypes.Header{
					BeginString:  "fix4.4",
					MsgType:      "A",
					SenderCompID: tx.From,
					TargetCompID: tx.To,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			},
			LogonAcceptor: &fixTypes.LogonAcceptor{
				Header: &fixTypes.Header{
					BeginString:  "fix4.4",
					MsgType:      "A",
					SenderCompID: tx.To,
					TargetCompID: tx.From,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			},
			Status:     "loggedIn",
			IsAccepted: true,
		}
		// set new session to store
		k.fixKeeper.SetSessions(ctx, tx.Hash, session)

		switch tx.Status {
		// set New Single Order
		// set Orders Execution Report
		// set Orders Rejection
		// set Trade Capture
		// check that Txs was successful
		case 0:
			// when status = 0, it means the transaction failed
			order := &fixTypes.Orders{
				SessionID: tx.Hash,
				Header: &fixTypes.Header{
					BeginString:  "FIX4.4",
					MsgType:      "D",
					SenderCompID: tx.From,
					TargetCompID: tx.To,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				ClOrdID:      tx.Hash,
				Symbol:       "ether",
				Side:         2,
				OrderQty:     "",
				OrdType:      1,
				Price:        tx.Value,
				TimeInForce:  1,
				Text:         "",
				TransactTime: strconv.FormatInt(tx.Timestamp, 10),
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			}
			k.fixKeeper.SetOrders(ctx, tx.Hash, *order)

			// when a transaction fails, order cancel reject is generated
			ordersCancelReject := fixTypes.OrdersCancelReject{
				SessionID: tx.Hash,
				Header: &fixTypes.Header{
					BeginString:  "FIX4.4",
					MsgType:      "9",
					SenderCompID: tx.To,
					TargetCompID: tx.From,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				OrderID:      tx.Hash,
				OrigClOrdID:  tx.Hash,
				ClOrdID:      tx.Hash,
				CxlRejReason: 4,
				TransactTime: strconv.FormatInt(tx.Timestamp, 10),
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			}
			k.fixKeeper.SetOrdersCancelReject(ctx, tx.Hash, ordersCancelReject)

			// set Trade Capture Report Reject when Txs fails
			tradeCaptureReportReject := &fixTypes.TradeCapture{
				SessionID: tx.Hash,
				TradeCaptureReport: &fixTypes.TradeCaptureReport{
					Header: &fixTypes.Header{
						BeginString:  "FIX4.4",
						MsgType:      "AE",
						SenderCompID: tx.From,
						TargetCompID: tx.To,
						MsgSeqNum:    tx.BlockNumber,
						SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
					},
					TradeReportID:        tx.Hash,
					TradeReportTransType: "New",
					TradeReportType:      "real-time",
					TrdType:              "Block Trade",
					TrdSubType:           "",
					Side:                 "2",
					OrderQty:             "",
					LastQty:              "",
					LastPx:               "",
					GrossTradeAmt:        tx.Value,
					ExecID:               tx.Hash,
					OrderID:              tx.Hash,
					TradeID:              tx.Hash,
					Symbol:               "ether",
					TransactTime:         strconv.FormatInt(tx.Timestamp, 10),
					Trailer: &fixTypes.Trailer{
						CheckSum: 0,
					},
				},
				TradeCaptureReportRejection: &fixTypes.TradeCaptureReportRejection{
					Header: &fixTypes.Header{
						BeginString:  "FIX4.4",
						MsgType:      "j",
						SenderCompID: tx.To,
						TargetCompID: tx.From,
						MsgSeqNum:    tx.BlockNumber,
						SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
					},
					TradeReportID:           tx.Hash,
					TradeReportRejectReason: 0,
					TradeReportRejectRefID:  tx.Hash,
					Text:                    "",
					Trailer: &fixTypes.Trailer{
						CheckSum: 0,
					},
				},
			}
			k.fixKeeper.SetTradeCapture(ctx, tx.Hash, *tradeCaptureReportReject)

		case 1:
			// set successful orders
			order := &fixTypes.Orders{
				SessionID: tx.Hash,
				Header: &fixTypes.Header{
					BeginString:  "FIX4.4",
					MsgType:      "D",
					SenderCompID: tx.From,
					TargetCompID: tx.To,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				ClOrdID:      tx.Hash,
				Symbol:       "ether",
				Side:         2,
				OrderQty:     "",
				OrdType:      1,
				Price:        tx.Value,
				TimeInForce:  1,
				Text:         "",
				TransactTime: strconv.FormatInt(tx.Timestamp, 10),
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			}
			k.fixKeeper.SetOrders(ctx, tx.Hash, *order)

			// set execution report
			orderExecutionReport := &fixTypes.OrdersExecutionReport{
				SessionID: tx.Hash,
				Header: &fixTypes.Header{
					BeginString:  "FIX4.4",
					MsgType:      "8",
					SenderCompID: tx.To,
					TargetCompID: tx.From,
					MsgSeqNum:    tx.BlockNumber,
					SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
				},
				ClOrdID:      tx.Hash,
				OrderID:      tx.Hash,
				ExecID:       tx.Hash,
				OrdStatus:    "New",
				ExecType:     "New",
				Symbol:       "ether",
				Side:         2,
				OrderQty:     "",
				Price:        tx.Value,
				TimeInForce:  1,
				LastPx:       0,
				LastQty:      0,
				LeavesQty:    0,
				CumQty:       0,
				AvgPx:        0,
				Text:         "",
				TransactTime: strconv.FormatInt(tx.Timestamp, 10),
				Trailer: &fixTypes.Trailer{
					CheckSum: 0,
				},
			}
			k.fixKeeper.SetOrdersExecutionReport(ctx, tx.Hash, *orderExecutionReport)

			// set Trade Capture Report if Txs was successful
			tradeCapture := &fixTypes.TradeCapture{
				SessionID: tx.Hash,
				TradeCaptureReport: &fixTypes.TradeCaptureReport{
					Header: &fixTypes.Header{
						BeginString:  "FIX4.4",
						MsgType:      "AE",
						SenderCompID: tx.From,
						TargetCompID: tx.To,
						MsgSeqNum:    tx.BlockNumber,
						SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
					},
					TradeReportID:        tx.Hash,
					TradeReportTransType: "New",
					TradeReportType:      "real-time",
					TrdType:              "Block Trade",
					TrdSubType:           "",
					Side:                 "2",
					OrderQty:             "",
					LastQty:              "",
					LastPx:               "",
					GrossTradeAmt:        tx.Value,
					ExecID:               tx.Hash,
					OrderID:              tx.Hash,
					TradeID:              tx.Hash,
					Symbol:               "ether",
					TransactTime:         strconv.FormatInt(tx.Timestamp, 10),
					Trailer: &fixTypes.Trailer{
						CheckSum: 0,
					},
				},
				TradeCaptureReportAcknowledgement: &fixTypes.TradeCaptureReportAcknowledgement{
					Header: &fixTypes.Header{
						BeginString:  "FIX4.4",
						MsgType:      "AR",
						SenderCompID: tx.To,
						TargetCompID: tx.From,
						MsgSeqNum:    tx.BlockNumber,
						SendingTime:  strconv.FormatInt(tx.Timestamp, 10),
					},
					TradeReportID:          tx.Hash,
					TradeID:                tx.Hash,
					SecondaryTradeID:       "",
					TradeReportType:        "real-time",
					TrdType:                "Block Trade",
					ExecType:               "New",
					TradeReportRefID:       tx.Hash,
					SecondaryTradeReportID: "",
					TradeReportStatus:      "Accepted",
					TradeTransType:         "",
					Text:                   "",
					Trailer: &fixTypes.Trailer{
						CheckSum: 0,
					},
				},
			}
			k.fixKeeper.SetTradeCapture(ctx, tx.Hash, *tradeCapture)
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

	return &types.MsgEthereumTxsResponse{}, err
}
