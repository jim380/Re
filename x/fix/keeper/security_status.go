package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityStatusCount get the total number of securityStatus
func (k Keeper) GetSecurityStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityStatusCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSecurityStatusCount set the total number of securityStatus
func (k Keeper) SetSecurityStatusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityStatusCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSecurityStatus appends a securityStatus in the store with a new id and update the count
func (k Keeper) AppendSecurityStatus(
	ctx sdk.Context,
	securityStatus types.SecurityStatus,
) uint64 {
	// Create the securityStatus
	count := k.GetSecurityStatusCount(ctx)

	// Set the ID of the appended value
	securityStatus.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityStatusKey))
	appendedValue := k.cdc.MustMarshal(&securityStatus)
	store.Set(GetSecurityStatusIDBytes(securityStatus.Id), appendedValue)

	// Update securityStatus count
	k.SetSecurityStatusCount(ctx, count+1)

	return count
}

// SetSecurityStatus set a specific securityStatus in the store
func (k Keeper) SetSecurityStatus(ctx sdk.Context, securityStatus types.SecurityStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityStatusKey))
	b := k.cdc.MustMarshal(&securityStatus)
	store.Set(GetSecurityStatusIDBytes(securityStatus.Id), b)
}

// GetSecurityStatus returns a securityStatus from its id
func (k Keeper) GetSecurityStatus(ctx sdk.Context, id uint64) (val types.SecurityStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityStatusKey))
	b := store.Get(GetSecurityStatusIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurityStatus removes a securityStatus from the store
func (k Keeper) RemoveSecurityStatus(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityStatusKey))
	store.Delete(GetSecurityStatusIDBytes(id))
}

// GetAllSecurityStatus returns all securityStatus
func (k Keeper) GetAllSecurityStatus(ctx sdk.Context) (list []types.SecurityStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityStatusKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSecurityStatusIDBytes returns the byte representation of the ID
func GetSecurityStatusIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSecurityStatusIDFromBytes returns ID in uint64 format from a byte array
func GetSecurityStatusIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
