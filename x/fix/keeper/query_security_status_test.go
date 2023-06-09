package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jim380/Re/testutil/keeper"
	"github.com/jim380/Re/testutil/nullify"
	"github.com/jim380/Re/x/fix/types"
)

func TestSecurityStatusQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSecurityStatus(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSecurityStatusRequest
		response *types.QueryGetSecurityStatusResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSecurityStatusRequest{Id: msgs[0].Id},
			response: &types.QueryGetSecurityStatusResponse{SecurityStatus: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetSecurityStatusRequest{Id: msgs[1].Id},
			response: &types.QueryGetSecurityStatusResponse{SecurityStatus: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetSecurityStatusRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SecurityStatus(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestSecurityStatusQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSecurityStatus(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSecurityStatusRequest {
		return &types.QueryAllSecurityStatusRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SecurityStatusAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SecurityStatus), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SecurityStatus),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SecurityStatusAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SecurityStatus), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SecurityStatus),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SecurityStatusAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SecurityStatus),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SecurityStatusAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
