package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/mic/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MarketIdentificationCodeAll(goCtx context.Context, req *types.QueryAllMarketIdentificationCodeRequest) (*types.QueryAllMarketIdentificationCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var marketIdentificationCodes []types.MarketIdentificationCode
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	marketIdentificationCodeStore := prefix.NewStore(store, types.KeyPrefix(types.MarketIdentificationCodeKeyPrefix))

	pageRes, err := query.Paginate(marketIdentificationCodeStore, req.Pagination, func(key []byte, value []byte) error {
		var marketIdentificationCode types.MarketIdentificationCode
		if err := k.cdc.Unmarshal(value, &marketIdentificationCode); err != nil {
			return err
		}

		marketIdentificationCodes = append(marketIdentificationCodes, marketIdentificationCode)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMarketIdentificationCodeResponse{MarketIdentificationCode: marketIdentificationCodes, Pagination: pageRes}, nil
}

func (k Keeper) MarketIdentificationCode(goCtx context.Context, req *types.QueryGetMarketIdentificationCodeRequest) (*types.QueryGetMarketIdentificationCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	marketIdentificationCode, found := k.GetMarketIdentificationCode(ctx, req.MIC)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMarketIdentificationCodeResponse{MarketIdentificationCode: marketIdentificationCode}, nil
}
