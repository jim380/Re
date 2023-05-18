package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetSecurityCount get the total number of security
func (k Keeper) GetSecurityCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSecurityCount set the total number of security
func (k Keeper) SetSecurityCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSecurity appends a security in the store with a new id and update the count
func (k Keeper) AppendSecurity(
    ctx sdk.Context,
    security types.Security,
) uint64 {
	// Create the security
    count := k.GetSecurityCount(ctx)

    // Set the ID of the appended value
    security.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
    appendedValue := k.cdc.MustMarshal(&security)
    store.Set(GetSecurityIDBytes(security.Id), appendedValue)

    // Update security count
    k.SetSecurityCount(ctx, count+1)

    return count
}

// SetSecurity set a specific security in the store
func (k Keeper) SetSecurity(ctx sdk.Context, security types.Security) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	b := k.cdc.MustMarshal(&security)
	store.Set(GetSecurityIDBytes(security.Id), b)
}

// GetSecurity returns a security from its id
func (k Keeper) GetSecurity(ctx sdk.Context, id uint64) (val types.Security, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	b := store.Get(GetSecurityIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurity removes a security from the store
func (k Keeper) RemoveSecurity(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	store.Delete(GetSecurityIDBytes(id))
}

// GetAllSecurity returns all security
func (k Keeper) GetAllSecurity(ctx sdk.Context) (list []types.Security) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Security
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetSecurityIDBytes returns the byte representation of the ID
func GetSecurityIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSecurityIDFromBytes returns ID in uint64 format from a byte array
func GetSecurityIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
