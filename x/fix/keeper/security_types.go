package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityTypesCount get the total number of securityTypes
func (k Keeper) GetSecurityTypesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetSecurityTypesCountKey()
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSecurityTypesCount set the total number of securityTypes
func (k Keeper) SetSecurityTypesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetSecurityTypesCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSecurityTypes set a specific securityTypes in the store
func (k Keeper) SetSecurityTypes(ctx sdk.Context, securityReqID string, securityTypes types.SecurityTypes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityTypesKey())
	key := []byte(securityReqID)
	b := k.cdc.MustMarshal(&securityTypes)
	store.Set(key, b)
}

// GetSecurityTypes returns a securityTypes using securityReqID
func (k Keeper) GetSecurityTypes(ctx sdk.Context, securityReqID string) (val types.SecurityTypes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityTypesKey())
	key := []byte(securityReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurityTypes removes a securityTypes from the store
func (k Keeper) RemoveSecurityTypes(ctx sdk.Context, securityReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityTypesKey())
	store.Delete([]byte(securityReqID))
}

// GetAllSecurityTypes returns all securityTypes
func (k Keeper) GetAllSecurityTypes(ctx sdk.Context) (list []types.SecurityTypes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityTypesKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityTypes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
