package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/mic/types"
)

// GetMarketIdentificationCodeCount get the total number of marketIdentificationCode
func (k Keeper) GetMarketIdentificationCodeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MarketIdentificationCodeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMarketIdentificationCodeCount set the total number of marketIdentificationCode
func (k Keeper) SetMarketIdentificationCodeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MarketIdentificationCodeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMarketIdentificationCode appends a marketIdentificationCode in the store with a new id and update the count
func (k Keeper) AppendMarketIdentificationCode(
	ctx sdk.Context,
	marketIdentificationCode types.MarketIdentificationCode,
) uint64 {
	// Create the marketIdentificationCode
	count := k.GetMarketIdentificationCodeCount(ctx)

	// Set the ID of the appended value
	marketIdentificationCode.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKey))
	appendedValue := k.cdc.MustMarshal(&marketIdentificationCode)
	store.Set(GetMarketIdentificationCodeIDBytes(marketIdentificationCode.Id), appendedValue)

	// Update marketIdentificationCode count
	k.SetMarketIdentificationCodeCount(ctx, count+1)

	return count
}

// SetMarketIdentificationCode set a specific marketIdentificationCode in the store
func (k Keeper) SetMarketIdentificationCode(ctx sdk.Context, marketIdentificationCode types.MarketIdentificationCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKey))
	b := k.cdc.MustMarshal(&marketIdentificationCode)
	store.Set(GetMarketIdentificationCodeIDBytes(marketIdentificationCode.Id), b)
}

// GetMarketIdentificationCode returns a marketIdentificationCode from its id
func (k Keeper) GetMarketIdentificationCode(ctx sdk.Context, id uint64) (val types.MarketIdentificationCode, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKey))
	b := store.Get(GetMarketIdentificationCodeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMarketIdentificationCode removes a marketIdentificationCode from the store
func (k Keeper) RemoveMarketIdentificationCode(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKey))
	store.Delete(GetMarketIdentificationCodeIDBytes(id))
}

// GetAllMarketIdentificationCode returns all marketIdentificationCode
func (k Keeper) GetAllMarketIdentificationCode(ctx sdk.Context) (list []types.MarketIdentificationCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MarketIdentificationCode
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMarketIdentificationCodeIDBytes returns the byte representation of the ID
func GetMarketIdentificationCodeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMarketIdentificationCodeIDFromBytes returns ID in uint64 format from a byte array
func GetMarketIdentificationCodeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
