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

func (k Keeper) OrderMassStatusAll(goCtx context.Context, req *types.QueryAllOrderMassStatusRequest) (*types.QueryAllOrderMassStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orderMassStatuss []types.OrderMassStatus
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	orderMassStatusStore := prefix.NewStore(store, types.GetOrderMassStatusKey())

	pageRes, err := query.Paginate(orderMassStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var orderMassStatus types.OrderMassStatus
		if err := k.cdc.Unmarshal(value, &orderMassStatus); err != nil {
			return err
		}

		orderMassStatuss = append(orderMassStatuss, orderMassStatus)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrderMassStatusResponse{OrderMassStatus: orderMassStatuss, Pagination: pageRes}, nil
}

func (k Keeper) OrderMassStatus(goCtx context.Context, req *types.QueryGetOrderMassStatusRequest) (*types.QueryGetOrderMassStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	orderMassStatus, found := k.GetOrderMassStatus(ctx, req.MassStatusReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrderMassStatusResponse{OrderMassStatus: orderMassStatus}, nil
}
