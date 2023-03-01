package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// SetOrders set a specific orders in the store
func (k Keeper) SetOrders(ctx sdk.Context, sessionID string, orders types.Orders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&orders)
	store.Set(key, b)
}

// GetOrders returns a orders from its sessionName
func (k Keeper) GetOrders(ctx sdk.Context, sessionID string) (val types.Orders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrders removes a orders from the store
func (k Keeper) RemoveOrders(ctx sdk.Context, sessionID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	store.Delete([]byte(sessionID))
}

// SetOrdersCancelRequest set a specific ordersCancelRequest in the store
func (k Keeper) SetOrdersCancelRequest(ctx sdk.Context, sessionID string, ordersCancelRequest types.OrdersCancelRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRequestKey))
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&ordersCancelRequest)
	store.Set(key, b)
}

// GetOrdersCancelRequest returns a ordersCancelRequest from its sessionName
func (k Keeper) GetOrdersCancelRequest(ctx sdk.Context, sessionID string) (val types.OrdersCancelRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRequestKey))
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersCancelRequest removes a ordersCancelRequest from the store
func (k Keeper) RemoveOrdersCancelRequest(ctx sdk.Context, sessionID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRequestKey))
	store.Delete([]byte(sessionID))
}


// SetOrdersCancelReject set a specific ordersCancelReject in the store
func (k Keeper) SetOrdersCancelReject(ctx sdk.Context, sessionID string, ordersCancelReject types.OrdersCancelReject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRejectKey))
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&ordersCancelReject)
	store.Set(key, b)
}

// GetOrdersCancelReject returns a ordersCancelReject from its id
func (k Keeper) GetOrdersCancelReject(ctx sdk.Context, sessionID string) (val types.OrdersCancelReject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRejectKey))
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersCancelReject removes a ordersCancelReject from the store
func (k Keeper) RemoveOrdersCancelReject(ctx sdk.Context, sessionID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersCancelRejectKey))
	store.Delete([]byte(sessionID))
}
