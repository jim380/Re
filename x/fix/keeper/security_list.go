package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityListCount get the total number of securityList
func (k Keeper) GetSecurityListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSecurityListCount set the total number of securityList
func (k Keeper) SetSecurityListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSecurityList appends a securityList in the store with a new id and update the count
func (k Keeper) AppendSecurityList(
	ctx sdk.Context,
	securityList types.SecurityList,
) uint64 {
	// Create the securityList
	count := k.GetSecurityListCount(ctx)

	// Set the ID of the appended value
	securityList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityListKey))
	appendedValue := k.cdc.MustMarshal(&securityList)
	store.Set(GetSecurityListIDBytes(securityList.Id), appendedValue)

	// Update securityList count
	k.SetSecurityListCount(ctx, count+1)

	return count
}

// SetSecurityList set a specific securityList in the store
func (k Keeper) SetSecurityList(ctx sdk.Context, securityList types.SecurityList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityListKey))
	b := k.cdc.MustMarshal(&securityList)
	store.Set(GetSecurityListIDBytes(securityList.Id), b)
}

// GetSecurityList returns a securityList from its id
func (k Keeper) GetSecurityList(ctx sdk.Context, id uint64) (val types.SecurityList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityListKey))
	b := store.Get(GetSecurityListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurityList removes a securityList from the store
func (k Keeper) RemoveSecurityList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityListKey))
	store.Delete(GetSecurityListIDBytes(id))
}

// GetAllSecurityList returns all securityList
func (k Keeper) GetAllSecurityList(ctx sdk.Context) (list []types.SecurityList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSecurityListIDBytes returns the byte representation of the ID
func GetSecurityListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSecurityListIDFromBytes returns ID in uint64 format from a byte array
func GetSecurityListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
