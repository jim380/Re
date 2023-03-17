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


func (k Keeper) SessionRejectAll(goCtx context.Context, req *types.QueryAllSessionRejectRequest) (*types.QueryAllSessionRejectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sessionRejects []types.SessionReject
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sessionRejectStore := prefix.NewStore(store, types.GetSessionRejectKey())

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
	sessionReject, found := k.GetSessionReject(ctx, req.SessionID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSessionRejectResponse{SessionReject: sessionReject}, nil
}

func (k Keeper) SessionLogoutAll(goCtx context.Context, req *types.QueryAllSessionLogoutRequest) (*types.QueryAllSessionLogoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sessionLogouts []types.SessionLogout
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sessionLogoutStore := prefix.NewStore(store, types.GetSessionLogoutKey())

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

