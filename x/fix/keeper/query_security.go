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

func (k Keeper) SecurityAll(goCtx context.Context, req *types.QueryAllSecurityRequest) (*types.QueryAllSecurityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var securitys []types.Security
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	securityStore := prefix.NewStore(store, types.GetSecurityKey())

	pageRes, err := query.Paginate(securityStore, req.Pagination, func(key []byte, value []byte) error {
		var security types.Security
		if err := k.cdc.Unmarshal(value, &security); err != nil {
			return err
		}

		securitys = append(securitys, security)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecurityResponse{Security: securitys, Pagination: pageRes}, nil
}

func (k Keeper) Security(goCtx context.Context, req *types.QueryGetSecurityRequest) (*types.QueryGetSecurityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	security, found := k.GetSecurity(ctx, req.SecurityReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSecurityResponse{Security: security}, nil
}
