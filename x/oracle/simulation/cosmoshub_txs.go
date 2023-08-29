package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jim380/Re/x/oracle/keeper"
	"github.com/jim380/Re/x/oracle/types"
)

func SimulateMsgCosmoshubTxs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCosmoshubTxs{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CosmoshubTxs simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CosmoshubTxs simulation not implemented"), nil, nil
	}
}
