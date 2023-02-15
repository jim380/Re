package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetAccountCount get the total number of account
func (k Keeper) GetAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AccountCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAccountCount set the total number of account
func (k Keeper) SetAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AccountCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAccount appends a account in the store with a new id and update the count
func (k Keeper) AppendAccount(
	ctx sdk.Context,
	account types.Account,
) uint64 {
	// Create the account
	count := k.GetAccountCount(ctx)

	// Set the ID of the appended value
	account.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	appendedValue := k.cdc.MustMarshal(&account)
	store.Set(GetAccountIDBytes(account.Id), appendedValue)

	// Update account count
	k.SetAccountCount(ctx, count+1)

	return count
}

// SetAccount set a specific account in the store
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	b := k.cdc.MustMarshal(&account)
	store.Set(GetAccountIDBytes(account.Id), b)
}

// GetAccount returns a account from its id
func (k Keeper) GetAccount(ctx sdk.Context, id uint64) (val types.Account, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	b := store.Get(GetAccountIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	store.Delete(GetAccountIDBytes(id))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (list []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAccountIDBytes returns the byte representation of the ID
func GetAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAccountIDFromBytes returns ID in uint64 format from a byte array
func GetAccountIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
