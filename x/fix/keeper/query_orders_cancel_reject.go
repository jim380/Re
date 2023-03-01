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

func (k Keeper) OrdersCancelRejectAll(goCtx context.Context, req *types.QueryAllOrdersCancelRejectRequest) (*types.QueryAllOrdersCancelRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersCancelRejects []types.OrdersCancelReject
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ordersCancelRejectStore := prefix.NewStore(store, types.KeyPrefix(types.OrdersCancelRejectKey))

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
