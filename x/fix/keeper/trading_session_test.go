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

func createNTradingSession(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TradingSession {
	items := make([]types.TradingSession, n)
	for i := range items {
		items[i].Id = keeper.AppendTradingSession(ctx, items[i])
	}
	return items
}

func TestTradingSessionGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradingSession(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTradingSession(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTradingSessionRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradingSession(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTradingSession(ctx, item.Id)
		_, found := keeper.GetTradingSession(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTradingSessionGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradingSession(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTradingSession(ctx)),
	)
}

func TestTradingSessionCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradingSession(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTradingSessionCount(ctx))
}
