package mic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/mic/keeper"
	"github.com/jim380/Re/x/mic/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the marketIdentificationCode
	for _, elem := range genState.MarketIdentificationCodeList {
		k.SetMarketIdentificationCode(ctx, elem.MIC, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MarketIdentificationCodeList = k.GetAllMarketIdentificationCode(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
