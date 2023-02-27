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

func (k Keeper) SessionRejectAll(goCtx context.Context, req *types.QueryAllSessionRejectRequest) (*types.QueryAllSessionRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sessionRejects []types.SessionReject
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sessionRejectStore := prefix.NewStore(store, types.KeyPrefix(types.SessionRejectKey))

	pageRes, err := query.Paginate(sessionRejectStore, req.Pagination, func(key []byte, value []byte) error {
		var sessionReject types.SessionReject
		if err := k.cdc.Unmarshal(value, &sessionReject); err != nil {
			return err
		}

		sessionRejects = append(sessionRejects, sessionReject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSessionRejectResponse{SessionReject: sessionRejects, Pagination: pageRes}, nil
}

func (k Keeper) SessionReject(goCtx context.Context, req *types.QueryGetSessionRejectRequest) (*types.QueryGetSessionRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sessionReject, found := k.GetSessionReject(ctx, req.SessionName)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSessionRejectResponse{SessionReject: sessionReject}, nil
}
