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

func (k Keeper) MarketDataAll(goCtx context.Context, req *types.QueryAllMarketDataRequest) (*types.QueryAllMarketDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var marketDatas []types.MarketData
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	marketDataStore := prefix.NewStore(store, types.GetMarketDataKey())

	pageRes, err := query.Paginate(marketDataStore, req.Pagination, func(key []byte, value []byte) error {
		var marketData types.MarketData
		if err := k.cdc.Unmarshal(value, &marketData); err != nil {
			return err
		}
		if marketData.MarketDataRequest.Header.ChainID == req.ChainID {
			marketDatas = append(marketDatas, marketData)
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMarketDataResponse{MarketData: marketDatas, Pagination: pageRes}, nil
}

func (k Keeper) MarketData(goCtx context.Context, req *types.QueryGetMarketDataRequest) (*types.QueryGetMarketDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	marketData, found := k.GetMarketData(ctx, req.MdReqID)
	if !found {
		return nil, sdkerrorstypes.ErrKeyNotFound
	}
	if marketData.MarketDataRequest.Header.ChainID != req.ChainID {
		return nil, sdkerrors.Wrapf(types.ErrWrongChainID, "chainID: %s", req.ChainID)
	}

	return &types.QueryGetMarketDataResponse{MarketData: marketData}, nil
}
