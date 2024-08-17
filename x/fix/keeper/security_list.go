package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityListCount get the total number of securityList
func (k Keeper) GetSecurityListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetSecurityListCountKey()
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
	byteKey := types.GetSecurityListCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSecurityList set a specific securityList in the store
func (k Keeper) SetSecurityList(ctx sdk.Context, securityReqID string, securityList types.SecurityList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityListKey())
	key := []byte(securityReqID)
	b := k.cdc.MustMarshal(&securityList)
	store.Set(key, b)
}

// GetSecurityList returns a securityList from its id
func (k Keeper) GetSecurityList(ctx sdk.Context, securityReqID string) (val types.SecurityList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityListKey())
	key := []byte(securityReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurityList removes a securityList from the store
func (k Keeper) RemoveSecurityList(ctx sdk.Context, securityReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityListKey())
	key := []byte(securityReqID)
	store.Delete(key)
}

// GetAllSecurityList returns all securityList
func (k Keeper) GetAllSecurityList(ctx sdk.Context) (list []types.SecurityList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityListKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
