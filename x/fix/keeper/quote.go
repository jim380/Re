package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// SetQuote set a specific quote in the store with quoteReqID
func (k Keeper) SetQuote(ctx sdk.Context, quoteReqID string, quote types.Quote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetQuoteKey())
	key := []byte(quoteReqID)
	b := k.cdc.MustMarshal(&quote)
	store.Set(key, b)
}

// GetQuote returns a quote from the store with quoteReqID
func (k Keeper) GetQuote(ctx sdk.Context, quoteReqID string) (val types.Quote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetQuoteKey())
	key := []byte(quoteReqID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveQuote removes a quote from the store
func (k Keeper) RemoveQuote(ctx sdk.Context, quoteReqID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetQuoteKey())
	key := []byte(quoteReqID)
	store.Delete(key)
}

// GetAllQuote returns all quote
func (k Keeper) GetAllQuote(ctx sdk.Context) (list []types.Quote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetQuoteKey())
	iterator := store.Iterator([]byte{}, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Quote
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
