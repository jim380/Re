package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jim380/Re/utils/bitcoin"
	"github.com/jim380/Re/utils/constants"
	fixTypes "github.com/jim380/Re/x/fix/types"
	"github.com/jim380/Re/x/oracle/types"
)

func (k msgServer) BitcoinTxs(goCtx context.Context, msg *types.MsgBitcoinTxs) (*types.MsgBitcoinTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate the message creator
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
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
								SenderCompID: tx.WitnessHash,
								TargetCompID: output.Addresses[0],
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
							},
							Trailer: &fixTypes.Trailer{
								CheckSum: 0,
							},
						},
						LogonAcceptor: &fixTypes.LogonAcceptor{
							Header: &fixTypes.Header{
								BeginString:  "fix4.4",
								MsgType:      "A",
								SenderCompID: output.Addresses[0],
								TargetCompID: tx.WitnessHash,
								MsgSeqNum:    int64(tx.BlockHeight),
								SendingTime:  tx.BlockTime,
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
					if tx.IsDoubleSpend == true {
						// orders are rejected
						
					}

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

	return &types.MsgBitcoinTxsResponse{}, err
}
