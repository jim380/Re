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

// AppendSessions appends a sessions in the store with a new id and update the count
func (k Keeper) AppendSessions(
	ctx sdk.Context,
	sessions types.Sessions,
) uint64 {
	// Create the sessions
	count := k.GetSessionsCount(ctx)

	// Set the ID of the appended value
	sessions.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	appendedValue := k.cdc.MustMarshal(&sessions)
	store.Set(GetSessionsIDBytes(sessions.Id), appendedValue)

	// Update sessions count
	k.SetSessionsCount(ctx, count+1)

	return count
}

// SetSessions set a specific sessions in the store
func (k Keeper) SetSessions(ctx sdk.Context, sessions types.Sessions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	b := k.cdc.MustMarshal(&sessions)
	store.Set(GetSessionsIDBytes(sessions.Id), b)
}

// GetSessions returns a sessions from its id
func (k Keeper) GetSessions(ctx sdk.Context, id uint64) (val types.Sessions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	b := store.Get(GetSessionsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSessions removes a sessions from the store
func (k Keeper) RemoveSessions(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SessionsKey))
	store.Delete(GetSessionsIDBytes(id))
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
