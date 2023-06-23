package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgSecurityListRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityListRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityListRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityListRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityListResponse(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityListResponse{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityListResponse simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityListResponse simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityListRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityListRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityListRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityListRequestReject simulation not implemented"), nil, nil
	}
}
