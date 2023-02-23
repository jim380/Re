package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

func (k Keeper) SetAccount(ctx sdk.Context, did string, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(did)
	bz := k.cdc.MustMarshalLengthPrefixed(&account)
	store.Set(key, bz)
}

func (k Keeper) GetAccount(ctx sdk.Context, did string) types.Account {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(did)
	bz := store.Get(key)
	if bz == nil {
		return types.Account{}
	}

	var account types.Account
	k.cdc.MustUnmarshalLengthPrefixed(bz, &account)
	return account
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(ctx sdk.Context, did string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	store.Delete([]byte(did))
}

func (k Keeper) getDIDs(ctx sdk.Context) []string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	dids := make([]string, 0)

	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		did := string(iter.Key())
		dids = append(dids, did)
	}
	return dids
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (list []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAccountIDBytes returns the byte representation of the ID
func GetAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAccountIDFromBytes returns ID in uint64 format from a byte array
func GetAccountIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
