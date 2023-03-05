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
