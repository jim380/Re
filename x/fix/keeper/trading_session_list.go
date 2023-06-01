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
	byteKey := types.GetTradingSessionListCountKey()
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
	byteKey := types.GetTradingSessionListCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetTradingSessionList set a specific tradingSessionList in the store
func (k Keeper) SetTradingSessionList(ctx sdk.Context, tradSesReqID string, tradingSessionList types.TradingSessionList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionListKey())
	key := []byte(tradSesReqID)
	b := k.cdc.MustMarshal(&tradingSessionList)
	store.Set(key, b)
}

// GetTradingSessionList returns a tradingSessionList using tradSesReqID
func (k Keeper) GetTradingSessionList(ctx sdk.Context, tradSesReqID string) (val types.TradingSessionList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionListKey())
	key := []byte(tradSesReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradingSessionList removes a tradingSessionList from the store
func (k Keeper) RemoveTradingSessionList(ctx sdk.Context, tradSesReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionListKey())
	key := []byte(tradSesReqID)
	store.Delete(key)
}

// GetAllTradingSessionList returns all tradingSessionList
func (k Keeper) GetAllTradingSessionList(ctx sdk.Context) (list []types.TradingSessionList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionListKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradingSessionList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
