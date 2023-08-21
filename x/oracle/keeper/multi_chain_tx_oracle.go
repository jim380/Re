package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/oracle/types"
)

// GetMultiChainTxOracleCount get the total number of multiChainTxOracle
func (k Keeper) GetMultiChainTxOracleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MultiChainTxOracleCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMultiChainTxOracleCount set the total number of multiChainTxOracle
func (k Keeper) SetMultiChainTxOracleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MultiChainTxOracleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMultiChainTxOracle appends a multiChainTxOracle in the store with a new id and update the count
func (k Keeper) AppendMultiChainTxOracle(
	ctx sdk.Context,
	multiChainTxOracle types.MultiChainTxOracle,
) uint64 {
	// Create the multiChainTxOracle
	count := k.GetMultiChainTxOracleCount(ctx)

	// Set the ID of the appended value
	multiChainTxOracle.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MultiChainTxOracleKey))
	appendedValue := k.cdc.MustMarshal(&multiChainTxOracle)
	store.Set(GetMultiChainTxOracleIDBytes(multiChainTxOracle.Id), appendedValue)

	// Update multiChainTxOracle count
	k.SetMultiChainTxOracleCount(ctx, count+1)

	return count
}

// SetMultiChainTxOracle set a specific multiChainTxOracle in the store
func (k Keeper) SetMultiChainTxOracle(ctx sdk.Context, multiChainTxOracle types.MultiChainTxOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MultiChainTxOracleKey))
	b := k.cdc.MustMarshal(&multiChainTxOracle)
	store.Set(GetMultiChainTxOracleIDBytes(multiChainTxOracle.Id), b)
}

// GetMultiChainTxOracle returns a multiChainTxOracle from its id
func (k Keeper) GetMultiChainTxOracle(ctx sdk.Context, id uint64) (val types.MultiChainTxOracle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MultiChainTxOracleKey))
	b := store.Get(GetMultiChainTxOracleIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMultiChainTxOracle removes a multiChainTxOracle from the store
func (k Keeper) RemoveMultiChainTxOracle(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MultiChainTxOracleKey))
	store.Delete(GetMultiChainTxOracleIDBytes(id))
}

// GetAllMultiChainTxOracle returns all multiChainTxOracle
func (k Keeper) GetAllMultiChainTxOracle(ctx sdk.Context) (list []types.MultiChainTxOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MultiChainTxOracleKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MultiChainTxOracle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMultiChainTxOracleIDBytes returns the byte representation of the ID
func GetMultiChainTxOracleIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMultiChainTxOracleIDFromBytes returns ID in uint64 format from a byte array
func GetMultiChainTxOracleIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
