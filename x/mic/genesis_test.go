package mic_test

import (
	"testing"

	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/mic"
	"github.com/jim380/Re/x/mic/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MarketIdentificationCodeList: []types.MarketIdentificationCode{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MarketIdentificationCodeCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MicKeeper(t)
	mic.InitGenesis(ctx, *k, genesisState)
	got := mic.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MarketIdentificationCodeList, got.MarketIdentificationCodeList)
	require.Equal(t, genesisState.MarketIdentificationCodeCount, got.MarketIdentificationCodeCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
