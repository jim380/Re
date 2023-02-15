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

func createNAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Id = keeper.AppendAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAccount(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx, item.Id)
		_, found := keeper.GetAccount(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccount(ctx)),
	)
}

func TestAccountCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAccountCount(ctx))
}
