package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetTradeCaptureCount get the total number of tradeCapture
func (k Keeper) GetTradeCaptureCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradeCaptureCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTradeCaptureCount set the total number of tradeCapture
func (k Keeper) SetTradeCaptureCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TradeCaptureCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTradeCapture appends a tradeCapture in the store with a new id and update the count
func (k Keeper) AppendTradeCapture(
	ctx sdk.Context,
	tradeCapture types.TradeCapture,
) uint64 {
	// Create the tradeCapture
	count := k.GetTradeCaptureCount(ctx)

	// Set the ID of the appended value
	tradeCapture.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeCaptureKey))
	appendedValue := k.cdc.MustMarshal(&tradeCapture)
	store.Set(GetTradeCaptureIDBytes(tradeCapture.Id), appendedValue)

	// Update tradeCapture count
	k.SetTradeCaptureCount(ctx, count+1)

	return count
}

// SetTradeCapture set a specific tradeCapture in the store
func (k Keeper) SetTradeCapture(ctx sdk.Context, tradeCapture types.TradeCapture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeCaptureKey))
	b := k.cdc.MustMarshal(&tradeCapture)
	store.Set(GetTradeCaptureIDBytes(tradeCapture.Id), b)
}

// GetTradeCapture returns a tradeCapture from its id
func (k Keeper) GetTradeCapture(ctx sdk.Context, id uint64) (val types.TradeCapture, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeCaptureKey))
	b := store.Get(GetTradeCaptureIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradeCapture removes a tradeCapture from the store
func (k Keeper) RemoveTradeCapture(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeCaptureKey))
	store.Delete(GetTradeCaptureIDBytes(id))
}

// GetAllTradeCapture returns all tradeCapture
func (k Keeper) GetAllTradeCapture(ctx sdk.Context) (list []types.TradeCapture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeCaptureKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradeCapture
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTradeCaptureIDBytes returns the byte representation of the ID
func GetTradeCaptureIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTradeCaptureIDFromBytes returns ID in uint64 format from a byte array
func GetTradeCaptureIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
