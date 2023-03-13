package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/mic/keeper"
	"github.com/jim380/Re/x/mic/types"
	"github.com/stretchr/testify/require"
)

func createNMarketIdentificationCode(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MarketIdentificationCode {
	items := make([]types.MarketIdentificationCode, n)
	for i := range items {
		items[i].Id = keeper.AppendMarketIdentificationCode(ctx, items[i])
	}
	return items
}

func TestMarketIdentificationCodeGet(t *testing.T) {
	keeper, ctx := keepertest.MicKeeper(t)
	items := createNMarketIdentificationCode(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMarketIdentificationCode(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMarketIdentificationCodeRemove(t *testing.T) {
	keeper, ctx := keepertest.MicKeeper(t)
	items := createNMarketIdentificationCode(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMarketIdentificationCode(ctx, item.Id)
		_, found := keeper.GetMarketIdentificationCode(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMarketIdentificationCodeGetAll(t *testing.T) {
	keeper, ctx := keepertest.MicKeeper(t)
	items := createNMarketIdentificationCode(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMarketIdentificationCode(ctx)),
	)
}

func TestMarketIdentificationCodeCount(t *testing.T) {
	keeper, ctx := keepertest.MicKeeper(t)
	items := createNMarketIdentificationCode(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMarketIdentificationCodeCount(ctx))
}
