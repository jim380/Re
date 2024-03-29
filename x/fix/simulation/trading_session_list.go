package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgTradingSessionListRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionListRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionListRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionListRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradingSessionListResponse(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionListResponse{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionListResponse simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionListResponse simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradingSessionListRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionListRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionListRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionListRequestReject simulation not implemented"), nil, nil
	}
}
