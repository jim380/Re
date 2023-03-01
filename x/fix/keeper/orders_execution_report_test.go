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

func createNOrdersExecutionReport(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OrdersExecutionReport {
	items := make([]types.OrdersExecutionReport, n)
	for i := range items {
		items[i].Id = keeper.AppendOrdersExecutionReport(ctx, items[i])
	}
	return items
}

func TestOrdersExecutionReportGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNOrdersExecutionReport(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetOrdersExecutionReport(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestOrdersExecutionReportRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNOrdersExecutionReport(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrdersExecutionReport(ctx, item.Id)
		_, found := keeper.GetOrdersExecutionReport(ctx, item.Id)
		require.False(t, found)
	}
}

func TestOrdersExecutionReportGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNOrdersExecutionReport(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrdersExecutionReport(ctx)),
	)
}

func TestOrdersExecutionReportCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNOrdersExecutionReport(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetOrdersExecutionReportCount(ctx))
}
