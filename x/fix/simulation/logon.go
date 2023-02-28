package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgLogonInitiator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLogonInitiator{
			InitiatorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the LogonInitiator simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LogonInitiator simulation not implemented"), nil, nil
	}
}

func SimulateMsgLogonAcceptor(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLogonAcceptor{
			AcceptorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the LogonAcceptor simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LogonAcceptor simulation not implemented"), nil, nil
	}
}

func SimulateMsgLogonReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLogonReject{
			AcceptorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the LogonReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LogonReject simulation not implemented"), nil, nil
	}
}

func SimulateMsgTerminateLogon(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTerminateLogon{
			InitiatorAddress: simAccount.Address.String(),
		}

		// TODO: Handling the TerminateLogon simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TerminateLogon simulation not implemented"), nil, nil
	}
}
