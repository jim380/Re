package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// SetOrders set a specific orders in the store
func (k Keeper) SetOrders(ctx sdk.Context, sessionName string, orders types.Orders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	key := []byte(sessionName)
	b := k.cdc.MustMarshal(&orders)
	store.Set(key, b)
}

// GetOrders returns a orders from its sessionName
func (k Keeper) GetOrders(ctx sdk.Context, sessionName string) (val types.Orders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	key := []byte(sessionName)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrders removes a orders from the store
func (k Keeper) RemoveOrders(ctx sdk.Context, sessionName string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	store.Delete([]byte(sessionName))
}
