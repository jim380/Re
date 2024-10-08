package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/fix/types"
)

// GetSessionsCount get the total number of sessions
func (k Keeper) GetSessionsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(string(types.GetSessionsCountKey()))
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
	byteKey := types.KeyPrefix(string(types.GetSessionsCountKey()))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetSessions set a specific sessions in the store
func (k Keeper) SetSessions(ctx sdk.Context, sessionID string, sessions types.Sessions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionsKey())
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&sessions)
	store.Set(key, b)
}

// GetSessions returns a sessions from its id
func (k Keeper) GetSessions(ctx sdk.Context, sessionID string) (val types.Sessions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionsKey())
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSessions removes a sessions from the store
func (k Keeper) RemoveSessions(ctx sdk.Context, sessionID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionsKey())
	store.Delete([]byte(sessionID))
}

// GetAllSessions returns all sessions
func (k Keeper) GetAllSessions(ctx sdk.Context) (list []types.Sessions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionsKey())
	iterator := store.Iterator([]byte{}, []byte{})

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

// SetSessionReject set a specific sessionReject in the store
func (k Keeper) SetSessionReject(ctx sdk.Context, sessionID string, sessionReject types.SessionReject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionRejectKey())
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&sessionReject)
	store.Set(key, b)
}

// GetSessionReject returns a sessionReject from its sessionName
func (k Keeper) GetSessionReject(ctx sdk.Context, sessionID string) (val types.SessionReject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionRejectKey())
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetSessionLogout set a specific sessionLogout in the store
func (k Keeper) SetSessionLogout(ctx sdk.Context, sessionID string, sessionLogout types.SessionLogout) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionLogoutKey())
	key := []byte(sessionID)
	b := k.cdc.MustMarshal(&sessionLogout)
	store.Set(key, b)
}

// GetSessionLogout returns a sessionLogout from its session name
func (k Keeper) GetSessionLogout(ctx sdk.Context, sessionID string) (val types.SessionLogout, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetSessionLogoutKey())
	key := []byte(sessionID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
