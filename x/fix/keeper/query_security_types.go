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

func (k Keeper) SecurityTypesAll(goCtx context.Context, req *types.QueryAllSecurityTypesRequest) (*types.QueryAllSecurityTypesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var securityTypess []types.SecurityTypes
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	securityTypesStore := prefix.NewStore(store, types.GetSecurityTypesKey())

	pageRes, err := query.Paginate(securityTypesStore, req.Pagination, func(key []byte, value []byte) error {
		var securityTypes types.SecurityTypes
		if err := k.cdc.Unmarshal(value, &securityTypes); err != nil {
			return err
		}
		if securityTypes.SecurityTypesRequest.Header.ChainID == req.ChainID {
			securityTypess = append(securityTypess, securityTypes)
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecurityTypesResponse{SecurityTypes: securityTypess, Pagination: pageRes}, nil
}

func (k Keeper) SecurityTypes(goCtx context.Context, req *types.QueryGetSecurityTypesRequest) (*types.QueryGetSecurityTypesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	securityTypes, found := k.GetSecurityTypes(ctx, req.SecurityReqID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	if securityTypes.SecurityTypesRequest.Header.ChainID != req.ChainID {
		return nil, sdkerrors.Wrapf(types.ErrWrongChainID, "chainID: %s", req.ChainID)
	}

	return &types.QueryGetSecurityTypesResponse{SecurityTypes: securityTypes}, nil
}
