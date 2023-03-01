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

func TestOrdersExecutionReportQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrdersExecutionReport(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOrdersExecutionReportRequest
		response *types.QueryGetOrdersExecutionReportResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetOrdersExecutionReportRequest{Id: msgs[0].Id},
			response: &types.QueryGetOrdersExecutionReportResponse{OrdersExecutionReport: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetOrdersExecutionReportRequest{Id: msgs[1].Id},
			response: &types.QueryGetOrdersExecutionReportResponse{OrdersExecutionReport: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetOrdersExecutionReportRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.OrdersExecutionReport(wctx, tc.request)
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

func TestOrdersExecutionReportQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FixKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrdersExecutionReport(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOrdersExecutionReportRequest {
		return &types.QueryAllOrdersExecutionReportRequest{
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
			resp, err := keeper.OrdersExecutionReportAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrdersExecutionReport), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrdersExecutionReport),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OrdersExecutionReportAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrdersExecutionReport), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrdersExecutionReport),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OrdersExecutionReportAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.OrdersExecutionReport),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OrdersExecutionReportAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
