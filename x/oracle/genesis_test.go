package oracle_test

import (
	"testing"

	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/oracle"
	"github.com/jim380/Re/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MultiChainTxOracleList: []types.MultiChainTxOracle{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MultiChainTxOracleCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MultiChainTxOracleList, got.MultiChainTxOracleList)
	require.Equal(t, genesisState.MultiChainTxOracleCount, got.MultiChainTxOracleCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
