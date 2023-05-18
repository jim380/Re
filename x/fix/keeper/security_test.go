package keeper_test

import (
	"testing"

    "github.com/jim380/Re/x/fix/keeper"
    "github.com/jim380/Re/x/fix/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNSecurity(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Security {
	items := make([]types.Security, n)
	for i := range items {
		items[i].Id = keeper.AppendSecurity(ctx, items[i])
	}
	return items
}

func TestSecurityGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurity(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSecurity(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSecurityRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecurity(ctx, item.Id)
		_, found := keeper.GetSecurity(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSecurityGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecurity(ctx)),
	)
}

func TestSecurityCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurity(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSecurityCount(ctx))
}
