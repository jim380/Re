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

func createNTradeCapture(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TradeCapture {
	items := make([]types.TradeCapture, n)
	for i := range items {
		items[i].Id = keeper.AppendTradeCapture(ctx, items[i])
	}
	return items
}

func TestTradeCaptureGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradeCapture(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTradeCapture(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTradeCaptureRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradeCapture(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTradeCapture(ctx, item.Id)
		_, found := keeper.GetTradeCapture(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTradeCaptureGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradeCapture(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTradeCapture(ctx)),
	)
}

func TestTradeCaptureCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNTradeCapture(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTradeCaptureCount(ctx))
}
