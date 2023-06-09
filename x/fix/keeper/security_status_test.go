package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
	"github.com/stretchr/testify/require"
)

func createNSecurityStatus(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecurityStatus {
	items := make([]types.SecurityStatus, n)
	for i := range items {
		items[i].Id = keeper.AppendSecurityStatus(ctx, items[i])
	}
	return items
}

func TestSecurityStatusGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityStatus(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSecurityStatus(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSecurityStatusRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityStatus(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecurityStatus(ctx, item.Id)
		_, found := keeper.GetSecurityStatus(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSecurityStatusGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityStatus(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecurityStatus(ctx)),
	)
}

func TestSecurityStatusCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityStatus(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSecurityStatusCount(ctx))
}
