package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MultiChainTxOracleAll(goCtx context.Context, req *types.QueryAllMultiChainTxOracleRequest) (*types.QueryAllMultiChainTxOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var multiChainTxOracles []types.MultiChainTxOracle
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	multiChainTxOracleStore := prefix.NewStore(store, types.GetMultiChainTxOracleKey())

	pageRes, err := query.Paginate(multiChainTxOracleStore, req.Pagination, func(key []byte, value []byte) error {
		var multiChainTxOracle types.MultiChainTxOracle
		if err := k.cdc.Unmarshal(value, &multiChainTxOracle); err != nil {
			return err
		}

		multiChainTxOracles = append(multiChainTxOracles, multiChainTxOracle)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMultiChainTxOracleResponse{MultiChainTxOracle: multiChainTxOracles, Pagination: pageRes}, nil
}

func (k Keeper) MultiChainTxOracle(goCtx context.Context, req *types.QueryGetMultiChainTxOracleRequest) (*types.QueryGetMultiChainTxOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	multiChainTxOracle, found := k.GetMultiChainTxOracle(ctx, req.OracleID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMultiChainTxOracleResponse{MultiChainTxOracle: multiChainTxOracle}, nil
}
