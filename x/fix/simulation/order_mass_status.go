package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgOrderMassStatusRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderMassStatusRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderMassStatusRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderMassStatusRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgOrderMassStatusReport(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderMassStatusReport{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderMassStatusReport simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderMassStatusReport simulation not implemented"), nil, nil
	}
}

func SimulateMsgOrderMassStatusRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOrderMassStatusRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OrderMassStatusRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OrderMassStatusRequestReject simulation not implemented"), nil, nil
	}
}
