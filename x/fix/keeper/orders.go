package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// SetOrders set a specific orders in the store
func (k Keeper) SetOrders(ctx sdk.Context, clOrdID string, orders types.Orders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersKey())
	key := []byte(clOrdID)
	b := k.cdc.MustMarshal(&orders)
	store.Set(key, b)
}

// GetOrders returns an order using clOrdID
func (k Keeper) GetOrders(ctx sdk.Context, clOrdID string) (val types.Orders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersKey())
	key := []byte(clOrdID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrders removes an order from the store
func (k Keeper) RemoveOrders(ctx sdk.Context, clOrdID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersKey())
	store.Delete([]byte(clOrdID))
}

// SetOrdersCancelRequest set a specific ordersCancelRequest in the store
func (k Keeper) SetOrdersCancelRequest(ctx sdk.Context, clOrdID string, ordersCancelRequest types.OrdersCancelRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRequestKey())
	key := []byte(clOrdID)
	b := k.cdc.MustMarshal(&ordersCancelRequest)
	store.Set(key, b)
}

// GetOrdersCancelRequest returns an ordersCancelRequest using clOrdID
func (k Keeper) GetOrdersCancelRequest(ctx sdk.Context, clOrdID string) (val types.OrdersCancelRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRequestKey())
	key := []byte(clOrdID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersCancelRequest removes an ordersCancelRequest from the store
func (k Keeper) RemoveOrdersCancelRequest(ctx sdk.Context, clOrdID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRequestKey())
	store.Delete([]byte(clOrdID))
}

// SetOrdersCancelReject set a specific ordersCancelReject in the store
func (k Keeper) SetOrdersCancelReject(ctx sdk.Context, clOrdID string, ordersCancelReject types.OrdersCancelReject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRejectKey())
	key := []byte(clOrdID)
	b := k.cdc.MustMarshal(&ordersCancelReject)
	store.Set(key, b)
}

// GetOrdersCancelReject returns an ordersCancelReject using clOrdID
func (k Keeper) GetOrdersCancelReject(ctx sdk.Context, clOrdID string) (val types.OrdersCancelReject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRejectKey())
	key := []byte(clOrdID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersCancelReject removes an ordersCancelReject from the store
func (k Keeper) RemoveOrdersCancelReject(ctx sdk.Context, clOrdID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersCancelRejectKey())
	store.Delete([]byte(clOrdID))
}

// SetOrdersExecutionReport set a specific ordersExecutionReport in the store
func (k Keeper) SetOrdersExecutionReport(ctx sdk.Context, clOrdID string, ordersExecutionReport types.OrdersExecutionReport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersExecutionReportKey())
	key := []byte(clOrdID)
	b := k.cdc.MustMarshal(&ordersExecutionReport)
	store.Set(key, b)
}

// GetOrdersExecutionReport returns an ordersExecutionReport using clOrdID
func (k Keeper) GetOrdersExecutionReport(ctx sdk.Context, clOrdID string) (val types.OrdersExecutionReport, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersExecutionReportKey())
	key := []byte(clOrdID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersExecutionReport removes an ordersExecutionReport from the store
func (k Keeper) RemoveOrdersExecutionReport(ctx sdk.Context, clOrdID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetOrdersExecutionReportKey())
	store.Delete([]byte(clOrdID))
}
