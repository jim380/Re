package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgSecurityDefinitionRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityDefinitionRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityDefinitionRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityDefinitionRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityDefinition(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityDefinition{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityDefinition simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityDefinition simulation not implemented"), nil, nil
	}
}

func SimulateMsgSecurityDefinitionRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSecurityDefinitionRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SecurityDefinitionRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SecurityDefinitionRequestReject simulation not implemented"), nil, nil
	}
}
