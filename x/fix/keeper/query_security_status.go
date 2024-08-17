package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jim380/Re/x/fix/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SecurityStatusAll(goCtx context.Context, req *types.QueryAllSecurityStatusRequest) (*types.QueryAllSecurityStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var securityStatuss []types.SecurityStatus
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	securityStatusStore := prefix.NewStore(store, types.GetSecurityStatusKey())

	pageRes, err := query.Paginate(securityStatusStore, req.Pagination, func(key []byte, value []byte) error {
		var securityStatus types.SecurityStatus
		if err := k.cdc.Unmarshal(value, &securityStatus); err != nil {
			return err
		}
		if securityStatus.SecurityStatusRequest.Header.ChainID == req.ChainID {
			securityStatuss = append(securityStatuss, securityStatus)
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecurityStatusResponse{SecurityStatus: securityStatuss, Pagination: pageRes}, nil
}

func (k Keeper) SecurityStatus(goCtx context.Context, req *types.QueryGetSecurityStatusRequest) (*types.QueryGetSecurityStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	securityStatus, found := k.GetSecurityStatus(ctx, req.SecurityStatusReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	if securityStatus.SecurityStatusRequest.Header.ChainID != req.ChainID {
		return nil, sdkerrors.Wrapf(types.ErrWrongChainID, "chainID: %s", req.ChainID)
	}

	return &types.QueryGetSecurityStatusResponse{SecurityStatus: securityStatus}, nil
}
