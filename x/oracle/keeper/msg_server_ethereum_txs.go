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
