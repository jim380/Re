package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/oracle/types"
)

// GetMultiChainTxOracleCount get the total number of multiChainTxOracle
func (k Keeper) GetMultiChainTxOracleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetMultiChainTxOracleCountKey()
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
	byteKey := types.GetMultiChainTxOracleCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetMultiChainTxOracle set a specific multiChainTxOracle in the store
func (k Keeper) SetMultiChainTxOracle(ctx sdk.Context, oracleID string, multiChainTxOracle types.MultiChainTxOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMultiChainTxOracleKey())
	key := []byte(oracleID)
	b := k.cdc.MustMarshal(&multiChainTxOracle)
	store.Set(key, b)
}

// GetMultiChainTxOracle returns a multiChainTxOracle from its id
func (k Keeper) GetMultiChainTxOracle(ctx sdk.Context, oracleID string) (val types.MultiChainTxOracle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMultiChainTxOracleKey())
	key := []byte(oracleID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMultiChainTxOracle removes a multiChainTxOracle from the store
func (k Keeper) RemoveMultiChainTxOracle(ctx sdk.Context, oracleID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMultiChainTxOracleKey())
	key := []byte(oracleID)
	store.Delete(key)
}

// GetAllMultiChainTxOracle returns all multiChainTxOracle
func (k Keeper) GetAllMultiChainTxOracle(ctx sdk.Context) (list []types.MultiChainTxOracle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetMultiChainTxOracleKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MultiChainTxOracle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
