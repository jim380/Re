package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgNewOrderSingle(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgNewOrderSingle{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the NewOrderSingle simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "NewOrderSingle simulation not implemented"), nil, nil
	}
}

func SimulateMsgOrderCancelRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderCancelRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderCancelRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderCancelRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgOrderExecutionReport(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderExecutionReport{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderExecutionReport simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderExecutionReport simulation not implemented"), nil, nil
	}
}

func SimulateMsgOrderCancelReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderCancelReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderCancelReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderCancelReject simulation not implemented"), nil, nil
	}
}
