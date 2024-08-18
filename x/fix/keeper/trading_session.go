package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetTradingSessionCount get the total number of tradingSession
func (k Keeper) GetTradingSessionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetTradingSessionCountKey()
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
	byteKey := types.GetTradingSessionCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetTradingSession set a specific tradingSession in the store
func (k Keeper) SetTradingSession(ctx sdk.Context, tradSesReqID string, tradingSession types.TradingSession) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionKey())
	key := []byte(tradSesReqID)
	b := k.cdc.MustMarshal(&tradingSession)
	store.Set(key, b)
}

// GetTradingSession returns a tradingSession using tradSesReqID
func (k Keeper) GetTradingSession(ctx sdk.Context, tradSesReqID string) (val types.TradingSession, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionKey())
	key := []byte(tradSesReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradingSession removes a tradingSession from the store
func (k Keeper) RemoveTradingSession(ctx sdk.Context, tradSesReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionKey())
	key := []byte(tradSesReqID)
	store.Delete(key)
}

// GetAllTradingSession returns all tradingSession
func (k Keeper) GetAllTradingSession(ctx sdk.Context) (list []types.TradingSession) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetTradingSessionKey())
	iterator := store.Iterator([]byte{}, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradingSession
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
