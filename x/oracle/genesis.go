package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/oracle/keeper"
	"github.com/jim380/Re/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the multiChainTxOracle
	for _, elem := range genState.MultiChainTxOracleList {
		k.SetMultiChainTxOracle(ctx, elem)
	}

	// Set multiChainTxOracle count
	k.SetMultiChainTxOracleCount(ctx, genState.MultiChainTxOracleCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MultiChainTxOracleList = k.GetAllMultiChainTxOracle(ctx)
	genesis.MultiChainTxOracleCount = k.GetMultiChainTxOracleCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
