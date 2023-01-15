package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/fix/keeper"
	"github.com/jim380/Re/x/fix/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx,
			item.Index,
		)
		_, found := keeper.GetAccount(ctx,
			item.Index,
		)
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
