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

// SetAccountCompanyName sets account's company name to store
func (k Keeper) SetAccountCompanyName(ctx sdk.Context, companyName string, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(companyName)
	bz := k.cdc.MustMarshalLengthPrefixed(&account)
	store.Set(key, bz)
}

// GetAccountCompanyName gets account's company name from store
func (k Keeper) GetAccountCompanyName(ctx sdk.Context, companyName string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(companyName)
	bz := store.Get(key)
	if bz == nil {
		return " "
	}

	var account types.Account
	k.cdc.MustUnmarshalLengthPrefixed(bz, &account)
	return account.CompanyName
}

// SetAccountWebsite sets account's website to store
func (k Keeper) SetAccountWebsite(ctx sdk.Context, website string, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(website)
	bz := k.cdc.MustMarshalLengthPrefixed(&account)
	store.Set(key, bz)
}

// GetAccountWebsite gets account's website from store
func (k Keeper) GetAccountWebsite(ctx sdk.Context, website string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	key := []byte(website)
	bz := store.Get(key)
	if bz == nil {
		return " "
	}

	var account types.Account
	k.cdc.MustUnmarshalLengthPrefixed(bz, &account)
	return account.Website
}

// SetAccountCreator sets Account's Creator address.
func (k Keeper) SetAccountCreator(ctx sdk.Context, creator sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := creator
	store.Set(types.GetAccountKey(), bz)
}

// GetAccountCreator gets Account's Creator address.
func (k Keeper) GetAccountCreator(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.GetAccountKey())
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey())
	store.Delete(GetAccountIDBytes(id))
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