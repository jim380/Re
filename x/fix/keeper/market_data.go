package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetMarketDataCount get the total number of marketData
func (k Keeper) GetMarketDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetMarketDataCountKey()
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMarketDataCount set the total number of marketData
func (k Keeper) SetMarketDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetMarketDataCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetMarketData set a specific marketData in the store
func (k Keeper) SetMarketData(ctx sdk.Context, mdReqID string, marketData types.MarketData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMarketDataKey())
	key := []byte(mdReqID)
	b := k.cdc.MustMarshal(&marketData)
	store.Set(key, b)
}

// GetMarketData returns a marketData from its mdReqID
func (k Keeper) GetMarketData(ctx sdk.Context, mdReqID string) (val types.MarketData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMarketDataKey())
	key := []byte(mdReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMarketData removes a marketData from the store
func (k Keeper) RemoveMarketData(ctx sdk.Context, mdReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMarketDataKey())
	key := []byte(mdReqID)
	store.Delete(key)
}

// GetAllMarketData returns all marketData
func (k Keeper) GetAllMarketData(ctx sdk.Context) []types.MarketData {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMarketDataKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	list := make([]types.MarketData, 0, k.GetMarketDataCount(ctx))

	for ; iterator.Valid(); iterator.Next() {
		var val types.MarketData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
