package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgLogoutInitiator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLogoutInitiator{
			InitiatorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the LogoutInitiator simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LogoutInitiator simulation not implemented"), nil, nil
	}
}

func SimulateMsgLogoutAcceptor(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLogoutAcceptor{
			AcceptorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the LogoutAcceptor simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LogoutAcceptor simulation not implemented"), nil, nil
	}
}
