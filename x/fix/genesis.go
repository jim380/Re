package fix

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, elem.Did, elem)
	}

	// Set all the sessions
	for _, elem := range genState.SessionsList {
		k.SetSessions(ctx, elem)
	}

	// Set sessions count
	k.SetSessionsCount(ctx, genState.SessionsCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AccountList = k.GetAllAccount(ctx)
	genesis.SessionsList = k.GetAllSessions(ctx)
	genesis.SessionsCount = k.GetSessionsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
