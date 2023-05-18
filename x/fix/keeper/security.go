package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityCount get the total number of security
func (k Keeper) GetSecurityCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
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
func (k Keeper) SetSecurityCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecurityCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSecurity set a specific security in the store
func (k Keeper) SetSecurity(ctx sdk.Context, securityReqID string, security types.Security) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	key := []byte(securityReqID)
	b := k.cdc.MustMarshal(&security)
	store.Set(key, b)
}

// GetSecurity returns a security from its id
func (k Keeper) GetSecurity(ctx sdk.Context, securityReqID string) (val types.Security, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	key := []byte(securityReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurity removes a security from the store
func (k Keeper) RemoveSecurity(ctx sdk.Context, securityReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecurityKey))
	key := []byte(securityReqID)
	store.Delete(key)
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
