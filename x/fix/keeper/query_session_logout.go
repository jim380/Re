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

func (k Keeper) SessionLogoutAll(goCtx context.Context, req *types.QueryAllSessionLogoutRequest) (*types.QueryAllSessionLogoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sessionLogouts []types.SessionLogout
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sessionLogoutStore := prefix.NewStore(store, types.KeyPrefix(types.SessionLogoutKey))

	pageRes, err := query.Paginate(sessionLogoutStore, req.Pagination, func(key []byte, value []byte) error {
		var sessionLogout types.SessionLogout
		if err := k.cdc.Unmarshal(value, &sessionLogout); err != nil {
			return err
		}

		sessionLogouts = append(sessionLogouts, sessionLogout)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSessionLogoutResponse{SessionLogout: sessionLogouts, Pagination: pageRes}, nil
}

func (k Keeper) SessionLogout(goCtx context.Context, req *types.QueryGetSessionLogoutRequest) (*types.QueryGetSessionLogoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sessionLogout, found := k.GetSessionLogout(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSessionLogoutResponse{SessionLogout: sessionLogout}, nil
}
