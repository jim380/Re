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
		// set sessions for parties existing in the transactions
		session := fixTypes.Sessions{
			SessionID: tx.Hash,
			LogonInitiator: &fixTypes.LogonInitiator{
				Header: &fixTypes.Header{
					BeginString:  "fix4.4",
					MsgType:      "A",
					SenderCompID: "",
					TargetCompID: "",
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
					SenderCompID: "",
					TargetCompID: "",
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
