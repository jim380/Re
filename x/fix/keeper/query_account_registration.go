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

func (k Keeper) AccountRegistrationAll(goCtx context.Context, req *types.QueryAllAccountRegistrationRequest) (*types.QueryAllAccountRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountRegistrations []types.AccountRegistration
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	accountRegistrationStore := prefix.NewStore(store, types.GetAccountRegistrationKey())

	pageRes, err := query.Paginate(accountRegistrationStore, req.Pagination, func(key []byte, value []byte) error {
		var accountRegistration types.AccountRegistration
		if err := k.cdc.Unmarshal(value, &accountRegistration); err != nil {
			return err
		}

		accountRegistrations = append(accountRegistrations, accountRegistration)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountRegistrationResponse{AccountRegistration: accountRegistrations, Pagination: pageRes}, nil
}

func (k Keeper) AccountRegistration(goCtx context.Context, req *types.QueryGetAccountRegistrationRequest) (*types.QueryGetAccountRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	accountRegistration, found := k.GetAccountRegistration(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAccountRegistrationResponse{AccountRegistration: accountRegistration}, nil
}
