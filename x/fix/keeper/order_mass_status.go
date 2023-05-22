package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetOrderMassStatusCount get the total number of orderMassStatus
func (k Keeper) GetOrderMassStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetOrderMassStatusCountKey()
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOrderMassStatusCount set the total number of orderMassStatus
func (k Keeper) SetOrderMassStatusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetOrderMassStatusCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetOrderMassStatus set a specific orderMassStatus in the store
func (k Keeper) SetOrderMassStatus(ctx sdk.Context, massStatusReqID string, orderMassStatus types.OrderMassStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrderMassStatusKey())
	key := []byte(massStatusReqID)
	b := k.cdc.MustMarshal(&orderMassStatus)
	store.Set(key, b)
}

// GetOrderMassStatus returns a orderMassStatus from its id
func (k Keeper) GetOrderMassStatus(ctx sdk.Context, massStatusReqID string) (val types.OrderMassStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrderMassStatusKey())
	key := []byte(massStatusReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrderMassStatus removes a orderMassStatus from the store
func (k Keeper) RemoveOrderMassStatus(ctx sdk.Context, massStatusReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrderMassStatusKey())
	key := []byte(massStatusReqID)
	store.Delete(key)
}

// GetAllOrderMassStatus returns all orderMassStatus
func (k Keeper) GetAllOrderMassStatus(ctx sdk.Context) (list []types.OrderMassStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrderMassStatusKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrderMassStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
