package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetTradingSessionListCount get the total number of tradingSessionList
func (k Keeper) GetTradingSessionListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradingSessionListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTradingSessionListCount set the total number of tradingSessionList
func (k Keeper) SetTradingSessionListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradingSessionListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTradingSessionList appends a tradingSessionList in the store with a new id and update the count
func (k Keeper) AppendTradingSessionList(
	ctx sdk.Context,
	tradingSessionList types.TradingSessionList,
) uint64 {
	// Create the tradingSessionList
	count := k.GetTradingSessionListCount(ctx)

	// Set the ID of the appended value
	tradingSessionList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionListKey))
	appendedValue := k.cdc.MustMarshal(&tradingSessionList)
	store.Set(GetTradingSessionListIDBytes(tradingSessionList.Id), appendedValue)

	// Update tradingSessionList count
	k.SetTradingSessionListCount(ctx, count+1)

	return count
}

// SetTradingSessionList set a specific tradingSessionList in the store
func (k Keeper) SetTradingSessionList(ctx sdk.Context, tradingSessionList types.TradingSessionList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionListKey))
	b := k.cdc.MustMarshal(&tradingSessionList)
	store.Set(GetTradingSessionListIDBytes(tradingSessionList.Id), b)
}

// GetTradingSessionList returns a tradingSessionList from its id
func (k Keeper) GetTradingSessionList(ctx sdk.Context, id uint64) (val types.TradingSessionList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionListKey))
	b := store.Get(GetTradingSessionListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradingSessionList removes a tradingSessionList from the store
func (k Keeper) RemoveTradingSessionList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionListKey))
	store.Delete(GetTradingSessionListIDBytes(id))
}

// GetAllTradingSessionList returns all tradingSessionList
func (k Keeper) GetAllTradingSessionList(ctx sdk.Context) (list []types.TradingSessionList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradingSessionList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTradingSessionListIDBytes returns the byte representation of the ID
func GetTradingSessionListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTradingSessionListIDFromBytes returns ID in uint64 format from a byte array
func GetTradingSessionListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
