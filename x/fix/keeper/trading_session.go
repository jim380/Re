package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetTradingSessionCount get the total number of tradingSession
func (k Keeper) GetTradingSessionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradingSessionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTradingSessionCount set the total number of tradingSession
func (k Keeper) SetTradingSessionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradingSessionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTradingSession appends a tradingSession in the store with a new id and update the count
func (k Keeper) AppendTradingSession(
	ctx sdk.Context,
	tradingSession types.TradingSession,
) uint64 {
	// Create the tradingSession
	count := k.GetTradingSessionCount(ctx)

	// Set the ID of the appended value
	tradingSession.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionKey))
	appendedValue := k.cdc.MustMarshal(&tradingSession)
	store.Set(GetTradingSessionIDBytes(tradingSession.Id), appendedValue)

	// Update tradingSession count
	k.SetTradingSessionCount(ctx, count+1)

	return count
}

// SetTradingSession set a specific tradingSession in the store
func (k Keeper) SetTradingSession(ctx sdk.Context, tradingSession types.TradingSession) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionKey))
	b := k.cdc.MustMarshal(&tradingSession)
	store.Set(GetTradingSessionIDBytes(tradingSession.Id), b)
}

// GetTradingSession returns a tradingSession from its id
func (k Keeper) GetTradingSession(ctx sdk.Context, id uint64) (val types.TradingSession, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionKey))
	b := store.Get(GetTradingSessionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradingSession removes a tradingSession from the store
func (k Keeper) RemoveTradingSession(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionKey))
	store.Delete(GetTradingSessionIDBytes(id))
}

// GetAllTradingSession returns all tradingSession
func (k Keeper) GetAllTradingSession(ctx sdk.Context) (list []types.TradingSession) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradingSessionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradingSession
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTradingSessionIDBytes returns the byte representation of the ID
func GetTradingSessionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTradingSessionIDFromBytes returns ID in uint64 format from a byte array
func GetTradingSessionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
