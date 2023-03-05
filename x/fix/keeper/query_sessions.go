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

func (k Keeper) SessionsAll(goCtx context.Context, req *types.QueryAllSessionsRequest) (*types.QueryAllSessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sessionss []types.Sessions
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sessionsStore := prefix.NewStore(store, types.GetSessionsKey())

	pageRes, err := query.Paginate(sessionsStore, req.Pagination, func(key []byte, value []byte) error {
		var sessions types.Sessions
		if err := k.cdc.Unmarshal(value, &sessions); err != nil {
			return err
		}

		sessionss = append(sessionss, sessions)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSessionsResponse{Sessions: sessionss, Pagination: pageRes}, nil
}

func (k Keeper) Sessions(goCtx context.Context, req *types.QueryGetSessionsRequest) (*types.QueryGetSessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sessions, found := k.GetSessions(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSessionsResponse{Sessions: sessions}, nil
}
