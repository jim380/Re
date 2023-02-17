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

func (k Keeper) AccountAll(c context.Context, req *types.QueryAllAccountRequest) (*types.QueryAllAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accounts []types.Account
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, types.GetAccountKey())

	pageRes, err := query.Paginate(accountStore, req.Pagination, func(key []byte, value []byte) error {
		var account types.Account
		if err := k.cdc.Unmarshal(value, &account); err != nil {
			return err
		}

		accounts = append(accounts, account)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountResponse{Account: accounts, Pagination: pageRes}, nil
}

func (k Keeper) Account(c context.Context, req *types.QueryGetAccountRequest) (*types.QueryGetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	account := k.GetAccount(ctx, req.Did)
	if account.Empty() {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAccountResponse{Account: account}, nil
}
