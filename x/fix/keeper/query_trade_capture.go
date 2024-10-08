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

func (k Keeper) TradeCaptureAll(goCtx context.Context, req *types.QueryAllTradeCaptureRequest) (*types.QueryAllTradeCaptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tradeCaptures []types.TradeCapture
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tradeCaptureStore := prefix.NewStore(store, types.GetTradeCaptureKey())

	pageRes, err := query.Paginate(tradeCaptureStore, req.Pagination, func(key []byte, value []byte) error {
		var tradeCapture types.TradeCapture
		if err := k.cdc.Unmarshal(value, &tradeCapture); err != nil {
			return err
		}
		if tradeCapture.TradeCaptureReport.Header.ChainID == req.ChainID {
			tradeCaptures = append(tradeCaptures, tradeCapture)
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTradeCaptureResponse{TradeCapture: tradeCaptures, Pagination: pageRes}, nil
}

func (k Keeper) TradeCapture(goCtx context.Context, req *types.QueryGetTradeCaptureRequest) (*types.QueryGetTradeCaptureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tradeCapture, found := k.GetTradeCapture(ctx, req.TradeReportID)
	if !found {
		return nil, sdkerrorstypes.ErrKeyNotFound
	}
	if tradeCapture.TradeCaptureReport.Header.ChainID != req.ChainID {
		return nil, sdkerrors.Wrapf(types.ErrWrongChainID, "chainID: %s", req.ChainID)
	}

	return &types.QueryGetTradeCaptureResponse{TradeCapture: tradeCapture}, nil
}
