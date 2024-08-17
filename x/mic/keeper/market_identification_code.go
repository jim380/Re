package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/mic/types"
)

// SetMarketIdentificationCode set a specific marketIdentificationCode in the store
func (k Keeper) SetMarketIdentificationCode(ctx sdk.Context, marketIdentificationCode types.MarketIdentificationCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKeyPrefix))
	b := k.cdc.MustMarshal(&marketIdentificationCode)
	store.Set(types.GetMarketIdentificationCodeKey(marketIdentificationCode.MIC), b)
}

// GetMarketIdentificationCode returns a marketIdentificationCode from its mic
func (k Keeper) GetMarketIdentificationCode(ctx sdk.Context, mic string) (val types.MarketIdentificationCode, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKeyPrefix))

	b := store.Get(types.GetMarketIdentificationCodeKey(mic))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMarketIdentificationCode removes a marketIdentificationCode from the store
func (k Keeper) RemoveMarketIdentificationCode(ctx sdk.Context, mic string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKeyPrefix))
	store.Delete(types.GetMarketIdentificationCodeKey(mic))
}

// GetAllMarketIdentificationCode returns all marketIdentificationCode
func (k Keeper) GetAllMarketIdentificationCode(ctx sdk.Context) (list []types.MarketIdentificationCode) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketIdentificationCodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MarketIdentificationCode
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
