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

func createNSessions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Sessions {
	items := make([]types.Sessions, n)
	for i := range items {
		items[i].Id = keeper.AppendSessions(ctx, items[i])
	}
	return items
}

func TestSessionsGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSessions(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSessions(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSessionsRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSessions(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSessions(ctx, item.Id)
		_, found := keeper.GetSessions(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSessionsGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSessions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSessions(ctx)),
	)
}

func TestSessionsCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSessions(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSessionsCount(ctx))
}
