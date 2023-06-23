package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgQuoteRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgQuoteRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the QuoteRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "QuoteRequest simulation not implemented"), nil, nil
	}
}

func SimulateMsgQuoteAcknowledgement(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgQuoteAcknowledgement{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the QuoteAcknowledgement simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "QuoteAcknowledgement simulation not implemented"), nil, nil
	}
}

func SimulateMsgQuoteRequestReject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgQuoteRequestReject{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the QuoteRequestReject simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "QuoteRequestReject simulation not implemented"), nil, nil
	}
}
