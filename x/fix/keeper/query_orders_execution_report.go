package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/fix/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrdersExecutionReportAll(goCtx context.Context, req *types.QueryAllOrdersExecutionReportRequest) (*types.QueryAllOrdersExecutionReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersExecutionReports []types.OrdersExecutionReport
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ordersExecutionReportStore := prefix.NewStore(store, types.GetOrdersExecutionReportKey())

	pageRes, err := query.Paginate(ordersExecutionReportStore, req.Pagination, func(key []byte, value []byte) error {
		var ordersExecutionReport types.OrdersExecutionReport
		if err := k.cdc.Unmarshal(value, &ordersExecutionReport); err != nil {
			return err
		}

		ordersExecutionReports = append(ordersExecutionReports, ordersExecutionReport)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersExecutionReportResponse{OrdersExecutionReport: ordersExecutionReports, Pagination: pageRes}, nil
}

func (k Keeper) OrdersExecutionReport(goCtx context.Context, req *types.QueryGetOrdersExecutionReportRequest) (*types.QueryGetOrdersExecutionReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ordersExecutionReport, found := k.GetOrdersExecutionReport(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrdersExecutionReportResponse{OrdersExecutionReport: ordersExecutionReport}, nil
}
