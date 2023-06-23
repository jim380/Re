package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgTradingSessionStatusRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionStatusRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionStatusRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionStatusRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradingSessionStatus(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionStatus{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionStatus simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionStatus simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradingSessionStatusRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradingSessionStatusRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradingSessionStatusRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradingSessionStatusRequestReject simulation not implemented"), nil, nil
	}
}
