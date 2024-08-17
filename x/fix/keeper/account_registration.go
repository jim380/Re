package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetAccountRegistrationCount get the total number of accountRegistration
func (k Keeper) GetAccountRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetAccountRegistrationCountKey()
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAccountRegistrationCount set the total number of accountRegistration
func (k Keeper) SetAccountRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.GetAccountRegistrationCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetAccountRegistration set a specific accountRegistration in the store
func (k Keeper) SetAccountRegistration(ctx sdk.Context, address string, accountRegistration types.AccountRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountRegistrationKey())
	key := []byte(address)
	b := k.cdc.MustMarshal(&accountRegistration)
	store.Set(key, b)
}

// GetAccountRegistration returns an accountRegistration using address
func (k Keeper) GetAccountRegistration(ctx sdk.Context, address string) (val types.AccountRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountRegistrationKey())
	key := []byte(address)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountRegistration removes a accountRegistration from the store
func (k Keeper) RemoveAccountRegistration(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountRegistrationKey())
	key := []byte(address)
	store.Delete(key)
}

// GetAllAccountRegistration returns all accountRegistration
func (k Keeper) GetAllAccountRegistration(ctx sdk.Context) []types.AccountRegistration {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountRegistrationKey())
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	list := make([]types.AccountRegistration, 0, k.GetAccountRegistrationCount(ctx))

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
