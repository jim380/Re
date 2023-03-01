package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetOrdersExecutionReportCount get the total number of ordersExecutionReport
func (k Keeper) GetOrdersExecutionReportCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OrdersExecutionReportCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOrdersExecutionReportCount set the total number of ordersExecutionReport
func (k Keeper) SetOrdersExecutionReportCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OrdersExecutionReportCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOrdersExecutionReport appends a ordersExecutionReport in the store with a new id and update the count
func (k Keeper) AppendOrdersExecutionReport(
	ctx sdk.Context,
	ordersExecutionReport types.OrdersExecutionReport,
) uint64 {
	// Create the ordersExecutionReport
	count := k.GetOrdersExecutionReportCount(ctx)

	// Set the ID of the appended value
	ordersExecutionReport.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersExecutionReportKey))
	appendedValue := k.cdc.MustMarshal(&ordersExecutionReport)
	store.Set(GetOrdersExecutionReportIDBytes(ordersExecutionReport.Id), appendedValue)

	// Update ordersExecutionReport count
	k.SetOrdersExecutionReportCount(ctx, count+1)

	return count
}

// SetOrdersExecutionReport set a specific ordersExecutionReport in the store
func (k Keeper) SetOrdersExecutionReport(ctx sdk.Context, ordersExecutionReport types.OrdersExecutionReport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersExecutionReportKey))
	b := k.cdc.MustMarshal(&ordersExecutionReport)
	store.Set(GetOrdersExecutionReportIDBytes(ordersExecutionReport.Id), b)
}

// GetOrdersExecutionReport returns a ordersExecutionReport from its id
func (k Keeper) GetOrdersExecutionReport(ctx sdk.Context, id uint64) (val types.OrdersExecutionReport, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersExecutionReportKey))
	b := store.Get(GetOrdersExecutionReportIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersExecutionReport removes a ordersExecutionReport from the store
func (k Keeper) RemoveOrdersExecutionReport(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersExecutionReportKey))
	store.Delete(GetOrdersExecutionReportIDBytes(id))
}

// GetAllOrdersExecutionReport returns all ordersExecutionReport
func (k Keeper) GetAllOrdersExecutionReport(ctx sdk.Context) (list []types.OrdersExecutionReport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersExecutionReportKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrdersExecutionReport
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOrdersExecutionReportIDBytes returns the byte representation of the ID
func GetOrdersExecutionReportIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetOrdersExecutionReportIDFromBytes returns ID in uint64 format from a byte array
func GetOrdersExecutionReportIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
