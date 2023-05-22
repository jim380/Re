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
	byteKey := types.KeyPrefix(types.OrderMassStatusCountKey)
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
	byteKey := types.KeyPrefix(types.OrderMassStatusCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOrderMassStatus appends a orderMassStatus in the store with a new id and update the count
func (k Keeper) AppendOrderMassStatus(
	ctx sdk.Context,
	orderMassStatus types.OrderMassStatus,
) uint64 {
	// Create the orderMassStatus
	count := k.GetOrderMassStatusCount(ctx)

	// Set the ID of the appended value
	orderMassStatus.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderMassStatusKey))
	appendedValue := k.cdc.MustMarshal(&orderMassStatus)
	store.Set(GetOrderMassStatusIDBytes(orderMassStatus.Id), appendedValue)

	// Update orderMassStatus count
	k.SetOrderMassStatusCount(ctx, count+1)

	return count
}

// SetOrderMassStatus set a specific orderMassStatus in the store
func (k Keeper) SetOrderMassStatus(ctx sdk.Context, orderMassStatus types.OrderMassStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderMassStatusKey))
	b := k.cdc.MustMarshal(&orderMassStatus)
	store.Set(GetOrderMassStatusIDBytes(orderMassStatus.Id), b)
}

// GetOrderMassStatus returns a orderMassStatus from its id
func (k Keeper) GetOrderMassStatus(ctx sdk.Context, id uint64) (val types.OrderMassStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderMassStatusKey))
	b := store.Get(GetOrderMassStatusIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrderMassStatus removes a orderMassStatus from the store
func (k Keeper) RemoveOrderMassStatus(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderMassStatusKey))
	store.Delete(GetOrderMassStatusIDBytes(id))
}

// GetAllOrderMassStatus returns all orderMassStatus
func (k Keeper) GetAllOrderMassStatus(ctx sdk.Context) (list []types.OrderMassStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderMassStatusKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrderMassStatus
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOrderMassStatusIDBytes returns the byte representation of the ID
func GetOrderMassStatusIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetOrderMassStatusIDFromBytes returns ID in uint64 format from a byte array
func GetOrderMassStatusIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
