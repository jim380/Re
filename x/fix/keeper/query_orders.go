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

func (k Keeper) OrdersAll(goCtx context.Context, req *types.QueryAllOrdersRequest) (*types.QueryAllOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orderss []types.Orders
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ordersStore := prefix.NewStore(store, types.GetOrdersKey())

	pageRes, err := query.Paginate(ordersStore, req.Pagination, func(key []byte, value []byte) error {
		var orders types.Orders
		if err := k.cdc.Unmarshal(value, &orders); err != nil {
			return err
		}

		orderss = append(orderss, orders)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersResponse{Orders: orderss, Pagination: pageRes}, nil
}

func (k Keeper) Orders(goCtx context.Context, req *types.QueryGetOrdersRequest) (*types.QueryGetOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	orders, found := k.GetOrders(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrdersResponse{Orders: orders}, nil
}

func (k Keeper) OrdersCancelRequestAll(goCtx context.Context, req *types.QueryAllOrdersCancelRequestRequest) (*types.QueryAllOrdersCancelRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersCancelRequests []types.OrdersCancelRequest
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ordersCancelRequestStore := prefix.NewStore(store, types.GetOrdersCancelRequestKey())

	pageRes, err := query.Paginate(ordersCancelRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var ordersCancelRequest types.OrdersCancelRequest
		if err := k.cdc.Unmarshal(value, &ordersCancelRequest); err != nil {
			return err
		}

		ordersCancelRequests = append(ordersCancelRequests, ordersCancelRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersCancelRequestResponse{OrdersCancelRequest: ordersCancelRequests, Pagination: pageRes}, nil
}

func (k Keeper) OrdersCancelRequest(goCtx context.Context, req *types.QueryGetOrdersCancelRequestRequest) (*types.QueryGetOrdersCancelRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ordersCancelRequest, found := k.GetOrdersCancelRequest(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrdersCancelRequestResponse{OrdersCancelRequest: ordersCancelRequest}, nil
}


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

func (k Keeper) OrdersCancelRejectAll(goCtx context.Context, req *types.QueryAllOrdersCancelRejectRequest) (*types.QueryAllOrdersCancelRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersCancelRejects []types.OrdersCancelReject
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ordersCancelRejectStore := prefix.NewStore(store, types.GetOrdersCancelRejectKey())

	pageRes, err := query.Paginate(ordersCancelRejectStore, req.Pagination, func(key []byte, value []byte) error {
		var ordersCancelReject types.OrdersCancelReject
		if err := k.cdc.Unmarshal(value, &ordersCancelReject); err != nil {
			return err
		}

		ordersCancelRejects = append(ordersCancelRejects, ordersCancelReject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersCancelRejectResponse{OrdersCancelReject: ordersCancelRejects, Pagination: pageRes}, nil
}

func (k Keeper) OrdersCancelReject(goCtx context.Context, req *types.QueryGetOrdersCancelRejectRequest) (*types.QueryGetOrdersCancelRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ordersCancelReject, found := k.GetOrdersCancelReject(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrdersCancelRejectResponse{OrdersCancelReject: ordersCancelReject}, nil
}

