package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

func SimulateMsgTradeCaptureReport(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradeCaptureReport{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradeCaptureReport simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradeCaptureReport simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradeCaptureReportAcknowledgement(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradeCaptureReportAcknowledgement{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradeCaptureReportAcknowledgement simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradeCaptureReportAcknowledgement simulation not implemented"), nil, nil
	}
}

func SimulateMsgTradeCaptureReportRejection(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTradeCaptureReportRejection{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TradeCaptureReportRejection simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TradeCaptureReportRejection simulation not implemented"), nil, nil
	}
}
