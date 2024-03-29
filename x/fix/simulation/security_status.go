package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgSecurityStatusRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityStatusRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityStatusRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityStatusRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityStatusResponse(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityStatusResponse{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityStatusResponse simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityStatusResponse simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityStatusRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityStatusRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityStatusRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityStatusRequestReject simulation not implemented"), nil, nil
	}
}
