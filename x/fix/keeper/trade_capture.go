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
	byteKey := types.GetTradeCaptureCountKey()
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
	byteKey := types.GetTradeCaptureCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetTradeCapture set a specific tradeCapture in the store
func (k Keeper) SetTradeCapture(ctx sdk.Context, tradeReportID string, tradeCapture types.TradeCapture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradeCaptureKey())
	key := []byte(tradeReportID)
	b := k.cdc.MustMarshal(&tradeCapture)
	store.Set(key, b)
}

// GetTradeCapture returns a tradeCapture from its id
func (k Keeper) GetTradeCapture(ctx sdk.Context, tradeReportID string) (val types.TradeCapture, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradeCaptureKey())
	key := []byte(tradeReportID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradeCapture removes a tradeCapture from the store
func (k Keeper) RemoveTradeCapture(ctx sdk.Context, tradeReportID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradeCaptureKey())
	key := []byte(tradeReportID)
	store.Delete(key)
}

// GetAllTradeCapture returns all tradeCapture
func (k Keeper) GetAllTradeCapture(ctx sdk.Context) (list []types.TradeCapture) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradeCaptureKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradeCapture
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
