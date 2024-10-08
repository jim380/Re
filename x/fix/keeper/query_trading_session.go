package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/fix/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TradingSessionAll(goCtx context.Context, req *types.QueryAllTradingSessionRequest) (*types.QueryAllTradingSessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tradingSessions []types.TradingSession
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tradingSessionStore := prefix.NewStore(store, types.GetTradingSessionKey())

	pageRes, err := query.Paginate(tradingSessionStore, req.Pagination, func(key []byte, value []byte) error {
		var tradingSession types.TradingSession
		if err := k.cdc.Unmarshal(value, &tradingSession); err != nil {
			return err
		}
		if tradingSession.TradingSessionStatusRequest.Header.ChainID == req.ChainID {
			tradingSessions = append(tradingSessions, tradingSession)
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTradingSessionResponse{TradingSession: tradingSessions, Pagination: pageRes}, nil
}

func (k Keeper) TradingSession(goCtx context.Context, req *types.QueryGetTradingSessionRequest) (*types.QueryGetTradingSessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tradingSession, found := k.GetTradingSession(ctx, req.TradSesReqID)
	if !found {
		return nil, sdkerrorstypes.ErrKeyNotFound
	}
	if tradingSession.TradingSessionStatusRequest.Header.ChainID != req.ChainID {
		return nil, sdkerrors.Wrapf(types.ErrWrongChainID, "chainID: %s", req.ChainID)
	}

	return &types.QueryGetTradingSessionResponse{TradingSession: tradingSession}, nil
}
