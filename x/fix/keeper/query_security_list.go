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

func (k Keeper) SecurityListAll(goCtx context.Context, req *types.QueryAllSecurityListRequest) (*types.QueryAllSecurityListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var securityLists []types.SecurityList
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	securityListStore := prefix.NewStore(store, types.GetSecurityListKey())

	pageRes, err := query.Paginate(securityListStore, req.Pagination, func(key []byte, value []byte) error {
		var securityList types.SecurityList
		if err := k.cdc.Unmarshal(value, &securityList); err != nil {
			return err
		}

		securityLists = append(securityLists, securityList)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecurityListResponse{SecurityList: securityLists, Pagination: pageRes}, nil
}

func (k Keeper) SecurityList(goCtx context.Context, req *types.QueryGetSecurityListRequest) (*types.QueryGetSecurityListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	securityList, found := k.GetSecurityList(ctx, req.SecurityReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSecurityListResponse{SecurityList: securityList}, nil
}
