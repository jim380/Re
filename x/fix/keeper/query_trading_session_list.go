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

func (k Keeper) TradingSessionListAll(goCtx context.Context, req *types.QueryAllTradingSessionListRequest) (*types.QueryAllTradingSessionListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tradingSessionLists []types.TradingSessionList
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tradingSessionListStore := prefix.NewStore(store, types.GetTradingSessionListKey())

	pageRes, err := query.Paginate(tradingSessionListStore, req.Pagination, func(key []byte, value []byte) error {
		var tradingSessionList types.TradingSessionList
		if err := k.cdc.Unmarshal(value, &tradingSessionList); err != nil {
			return err
		}

		tradingSessionLists = append(tradingSessionLists, tradingSessionList)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTradingSessionListResponse{TradingSessionList: tradingSessionLists, Pagination: pageRes}, nil
}

func (k Keeper) TradingSessionList(goCtx context.Context, req *types.QueryGetTradingSessionListRequest) (*types.QueryGetTradingSessionListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tradingSessionList, found := k.GetTradingSessionList(ctx, req.TradSesReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTradingSessionListResponse{TradingSessionList: tradingSessionList}, nil
}
