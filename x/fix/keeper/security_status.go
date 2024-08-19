package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSecurityStatusCount get the total number of securityStatus
func (k Keeper) GetSecurityStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetSecurityStatusCountKey()
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
	byteKey := types.GetSecurityStatusCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSecurityStatus set a specific securityStatus in the store
func (k Keeper) SetSecurityStatus(ctx sdk.Context, securityStatusReqID string, securityStatus types.SecurityStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityStatusKey())
	b := k.cdc.MustMarshal(&securityStatus)
	key := []byte(securityStatusReqID)
	store.Set(key, b)
}

// GetSecurityStatus returns a securityStatus using securityStatusReqID
func (k Keeper) GetSecurityStatus(ctx sdk.Context, securityStatusReqID string) (val types.SecurityStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityStatusKey())
	key := []byte(securityStatusReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecurityStatus removes a securityStatus from the store
func (k Keeper) RemoveSecurityStatus(ctx sdk.Context, securityStatusReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityStatusKey())
	store.Delete([]byte(securityStatusReqID))
}

// GetAllSecurityStatus returns all securityStatus
func (k Keeper) GetAllSecurityStatus(ctx sdk.Context) (list []types.SecurityStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSecurityStatusKey())
	iterator := store.Iterator([]byte{}, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecurityStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
