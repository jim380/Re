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

func createNSecurityTypes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecurityTypes {
	items := make([]types.SecurityTypes, n)
	for i := range items {
		items[i].Id = keeper.AppendSecurityTypes(ctx, items[i])
	}
	return items
}

func TestSecurityTypesGet(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityTypes(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSecurityTypes(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSecurityTypesRemove(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityTypes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecurityTypes(ctx, item.Id)
		_, found := keeper.GetSecurityTypes(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSecurityTypesGetAll(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityTypes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecurityTypes(ctx)),
	)
}

func TestSecurityTypesCount(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	items := createNSecurityTypes(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSecurityTypesCount(ctx))
}
