package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSessionsCount get the total number of sessions
func (k Keeper) GetSessionsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SessionsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSessionsCount set the total number of sessions
func (k Keeper) SetSessionsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SessionsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSessions set a specific sessions in the store
func (k Keeper) SetSessions(ctx sdk.Context, sessionName string, sessions types.Sessions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	key := []byte(sessionName)
	b := k.cdc.MustMarshal(&sessions)
	store.Set(key, b)
}

// GetSessions returns a sessions from its id
func (k Keeper) GetSessions(ctx sdk.Context, sessionName string) (val types.Sessions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	key := []byte(sessionName)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSessions removes a sessions from the store
func (k Keeper) RemoveSessions(ctx sdk.Context, sessionName string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	store.Delete([]byte(sessionName))
}

// GetAllSessions returns all sessions
func (k Keeper) GetAllSessions(ctx sdk.Context) (list []types.Sessions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Sessions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSessionsIDBytes returns the byte representation of the ID
func GetSessionsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSessionsIDFromBytes returns ID in uint64 format from a byte array
func GetSessionsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
