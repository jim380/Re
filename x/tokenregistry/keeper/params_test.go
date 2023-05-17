package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/x/tokenregistry/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenregistryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
