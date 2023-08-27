package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/constants"
	cosmostxs "github.com/jim380/Re/utils/cosmos_txs"
	"github.com/jim380/Re/utils/helpers"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

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
	// start refetching data every 5 seconds
	go cosmostxs.RefetchTxsDataPeriodically(cosmostxs.CacheKey)
	transactions, err := cosmostxs.GetTxsDataWithCache(cosmostxs.CacheKey)
	for _, tx := range transactions {
		for _, message := range tx.Messages {

			// type of Txs message = Send
			if helpers.AbbrTxMessage(message.Type) == "Send" {
				// set sessions for parties existing in the transactions
				session := fixTypes.Sessions{
					SessionID: tx.TxHash,
					LogonInitiator: &fixTypes.LogonInitiator{
						Header: &fixTypes.Header{
							BeginString:  "fix4.4",
							MsgType:      "A",
							SenderCompID: message.FromAddress,
							TargetCompID: message.ToAddress,
							SendingTime:  tx.Time,
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
							SendingTime:  tx.Time,
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
				log.Println("When it is in range", session)

				// set New Single Order
				// check that Txs was success
				if tx.Result != 0 {
					// orders here were rejected
					ordersCancelReject := fixTypes.OrdersCancelReject{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString: "FIX4.2",
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

				} else {
					// set successful orders
					order := &fixTypes.Orders{
						SessionID: tx.TxHash,
						Header: &fixTypes.Header{
							BeginString: "FIX4.2",
							MsgType:     "D",
						},
						ClOrdID:      tx.TxHash,
						Symbol:       "",
						Side:         2,
						OrderQty:     "",
						OrdType:      1,
						Price:        "",
						TimeInForce:  1,
						Text:         "",
						TransactTime: "",
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrders(ctx, tx.TxHash, *order)

					// set execution report
					orderExecutionReport := &fixTypes.OrdersExecutionReport{
						SessionID:    tx.TxHash,
						Header:       &fixTypes.Header{},
						ClOrdID:      tx.TxHash,
						OrderID:      "",
						ExecID:       "",
						OrdStatus:    "",
						ExecType:     "",
						Symbol:       "",
						Side:         2,
						OrderQty:     "",
						Price:        "",
						TimeInForce:  1,
						LastPx:       2,
						LastQty:      2,
						LeavesQty:    2,
						CumQty:       2,
						AvgPx:        2,
						Text:         "",
						TransactTime: tx.Time,
						Trailer: &fixTypes.Trailer{
							CheckSum: 0,
						},
					}
					k.fixKeeper.SetOrdersExecutionReport(ctx, tx.TxHash, *orderExecutionReport)

				}

			}
		}
	}

	// set
	//log.Println("When it is in range", account)
	//k.fixKeeper.SetAccountRegistration(ctx, account.Address, account)

	oracle := types.MultiChainTxOracle{
		OracleId:  msg.OracleId,
		Address:   msg.Address,
		Timestamp: constants.CreatedAt,
	}

	// set oracle to store
	k.SetMultiChainTxOracle(ctx, msg.OracleId, oracle)

	// emit event
	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgCosmoshubTxsResponse{}, err
}
