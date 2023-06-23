package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgMarketDataRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMarketDataRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MarketDataRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MarketDataRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgMarketDataSnapshotFullRefresh(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMarketDataSnapshotFullRefresh{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MarketDataSnapshotFullRefresh simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MarketDataSnapshotFullRefresh simulation not implemented"), nil, nil
	}
}

func SimulateMsgMarketDataIncremental(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMarketDataIncremental{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MarketDataIncremental simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MarketDataIncremental simulation not implemented"), nil, nil
	}
}

func SimulateMsgMarketDataRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMarketDataRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MarketDataRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MarketDataRequestReject simulation not implemented"), nil, nil
	}
}
