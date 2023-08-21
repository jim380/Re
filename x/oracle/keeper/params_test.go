package keeper_test

import (
	"testing"

	testkeeper "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
