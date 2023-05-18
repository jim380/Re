package fix_test

import (
	"testing"

	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/fix"
	"github.com/jim380/Re/x/fix/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		SecurityList: []types.Security{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	SecurityCount: 2,
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FixKeeper(t)
	fix.InitGenesis(ctx, *k, genesisState)
	got := fix.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	require.ElementsMatch(t, genesisState.SecurityList, got.SecurityList)
require.Equal(t, genesisState.SecurityCount, got.SecurityCount)
// this line is used by starport scaffolding # genesis/test/assert
}
