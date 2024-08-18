package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/bitcoin"
	"github.com/jim380/Re/utils/constants"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

// BitcoinTxs generate the logon, new single order, execution report and
// trade capture report for transactions on Bitcoin
func (k msgServer) BitcoinTxs(goCtx context.Context, msg *types.MsgBitcoinTxs) (*types.MsgBitcoinTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrorstypes.ErrInvalidAddress, msg.Creator)
	}
	if msg.Creator != msg.Address {
		return nil, sdkerrors.Wrapf(types.ErrWrongAddress, "Address: %s", msg.Address)
	}

	// set Bitcoin txs, map to the fix module
	// update fix module from oracle
	// start refetching data every 5 seconds and return latest Bitcoin Txs
	go bitcoin.RefetchBitcoinTxsDataPeriodically(constants.BitcoinCacheKey)
	transactions, err := bitcoin.GetCachedBitcoinTransactions(constants.BitcoinCacheKey)
	if err != nil {
		return nil, fmt.Errorf("failed to return queried chain transactions data: %w", err)
	}
	for _, tx := range transactions.Transactions {
		switch tx.IsCoinbase {
		case true:
			for _, output := range tx.Outputs {
				// check that address is not empty
				if output.Addresses[0] != "" {
					// bitcoin coinbase txs type
					// set sessions for parties existing in the transactions
					// A Coinbase transaction, often referred to as a "coinbase tx" or "generation transaction,"
					// is a special type of transaction in the Bitcoin blockchain.
					// It plays a crucial role in the functioning of the Bitcoin network and
					// is the first transaction in a new block, in the context of Bitcoin,
					// "Coinbase" refers to the transaction that creates new cryptocurrency coins as a
					// block reward for miners. This is sometimes referred to as the "coinbase reward"
					// or "block reward.
					// Assumptions are made here, the values can be changed here
					session := fixTypes.Sessions{
						SessionID: tx.Hash,
						LogonInitiator: &fixTypes.LogonInitiator{
							Header: &fixTypes.Header{
								BeginString:  "fix4.4",
								MsgType:      "A",
								SenderCompID: output.Addresses[0],
								TargetCompID: tx.WitnessHash,
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
							},
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						LogonAcceptor: &fixTypes.LogonAcceptor{
							Header: &fixTypes.Header{
								BeginString:  "fix4.4",
								MsgType:      "A",
								SenderCompID: tx.WitnessHash,
								TargetCompID: output.Addresses[0],
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
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

					// set New Single Order
					// set Orders Execution Report
					// set Orders Rejection
					// set Trade Capture
					// check that Txs was successful
					if tx.IsDoubleSpend {
						// orders here were rejected
						order := &fixTypes.Orders{
							SessionID: tx.Hash,
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "D",
								SenderCompID: output.Addresses[0],
								TargetCompID: tx.WitnessHash,
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
							},
							ClOrdID:      tx.Hash,
							Symbol:       "btc",
							Side:         2,
							OrderQty:     "",
							OrdType:      1,
							Price:        strconv.Itoa(output.Value),
							TimeInForce:  1,
							Text:         "",
							TransactTime: tx.BlockTime,
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						}
						k.fixKeeper.SetOrders(ctx, tx.Hash, *order)

						// when a send transaction fails, order cancel reject is generated
						ordersCancelReject := fixTypes.OrdersCancelReject{
							SessionID: tx.Hash,
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "9",
								SenderCompID: tx.WitnessHash,
								TargetCompID: output.Addresses[0],
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
							},
							OrderID:      tx.Hash,
							OrigClOrdID:  tx.Hash,
							ClOrdID:      tx.Hash,
							CxlRejReason: 4,
							TransactTime: tx.BlockTime,
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
									SenderCompID: output.Addresses[0],
									TargetCompID: tx.WitnessHash,
									MsgSeqNum:    int64(tx.BlockHeight),
									SendingTime:  tx.BlockTime,
									ChainID:      constants.BitcoinChainID,
								},
								TradeReportID:        tx.Hash,
								TradeReportTransType: "New",
								TradeReportType:      "real-time",
								TrdType:              "Block Trade",
								TrdSubType:           "Coinbase",
								Side:                 "2",
								OrderQty:             "",
								LastQty:              "",
								LastPx:               "",
								GrossTradeAmt:        strconv.Itoa(output.Value),
								ExecID:               tx.Hash,
								OrderID:              tx.Hash,
								TradeID:              tx.Hash,
								Symbol:               "btc",
								TransactTime:         tx.BlockTime,
								Trailer: &fixTypes.Trailer{
									CheckSum: 0,
								},
							},
							TradeCaptureReportRejection: &fixTypes.TradeCaptureReportRejection{
								Header: &fixTypes.Header{
									BeginString:  "FIX4.4",
									MsgType:      "j",
									SenderCompID: tx.WitnessHash,
									TargetCompID: output.Addresses[0],
									MsgSeqNum:    int64(tx.BlockHeight),
									SendingTime:  tx.BlockTime,
									ChainID:      constants.BitcoinChainID,
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

					} else {
						// set successful orders
						order := &fixTypes.Orders{
							SessionID: tx.Hash,
							Header: &fixTypes.Header{
								BeginString:  "FIX4.4",
								MsgType:      "D",
								SenderCompID: output.Addresses[0],
								TargetCompID: tx.WitnessHash,
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
							},
							ClOrdID:      tx.Hash,
							Symbol:       "btc",
							Side:         2,
							OrderQty:     "",
							OrdType:      1,
							Price:        strconv.Itoa(output.Value),
							TimeInForce:  1,
							Text:         "",
							TransactTime: tx.BlockTime,
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
								SenderCompID: tx.WitnessHash,
								TargetCompID: output.Addresses[0],
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
								ChainID:      constants.BitcoinChainID,
							},
							ClOrdID:      tx.Hash,
							OrderID:      tx.Hash,
							ExecID:       tx.Hash,
							OrdStatus:    "New",
							ExecType:     "New",
							Symbol:       "btc",
							Side:         2,
							OrderQty:     "",
							Price:        strconv.Itoa(output.Value),
							TimeInForce:  1,
							LastPx:       0,
							LastQty:      0,
							LeavesQty:    0,
							CumQty:       0,
							AvgPx:        0,
							Text:         "",
							TransactTime: tx.BlockTime,
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
									SenderCompID: output.Addresses[0],
									TargetCompID: tx.WitnessHash,
									MsgSeqNum:    int64(tx.BlockHeight),
									SendingTime:  tx.BlockTime,
									ChainID:      constants.BitcoinChainID,
								},
								TradeReportID:        tx.Hash,
								TradeReportTransType: "New",
								TradeReportType:      "real-time",
								TrdType:              "Block Trade",
								TrdSubType:           "Coinbase",
								Side:                 "2",
								OrderQty:             "",
								LastQty:              "",
								LastPx:               "",
								GrossTradeAmt:        strconv.Itoa(output.Value),
								ExecID:               tx.Hash,
								OrderID:              tx.Hash,
								TradeID:              tx.Hash,
								Symbol:               "btc",
								TransactTime:         tx.BlockTime,
								Trailer: &fixTypes.Trailer{
									CheckSum: 0,
								},
							},
							TradeCaptureReportAcknowledgement: &fixTypes.TradeCaptureReportAcknowledgement{
								Header: &fixTypes.Header{
									BeginString:  "FIX4.4",
									MsgType:      "AR",
									SenderCompID: tx.WitnessHash,
									TargetCompID: output.Addresses[0],
									MsgSeqNum:    int64(tx.BlockHeight),
									SendingTime:  tx.BlockTime,
									ChainID:      constants.BitcoinChainID,
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
								TradeTransType:         "Coinbase",
								Text:                   "",
								Trailer: &fixTypes.Trailer{
									CheckSum: 0,
								},
							},
						}
						k.fixKeeper.SetTradeCapture(ctx, tx.Hash, *tradeCapture)
					}
				}
			}

		case false:
			// TODO
			// add other types of bitcoin transactions

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

	return &types.MsgBitcoinTxsResponse{}, err
}
