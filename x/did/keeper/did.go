package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jim380/Re/x/did/types"
)

func (k Keeper) SetDIDDocument(ctx sdk.Context, did string, doc types.DIDDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetDIDDocumentKey())
	key := []byte(did)
	bz := k.cdc.MustMarshalLengthPrefixed(&doc)
	store.Set(key, bz)
}

func (k Keeper) GetDIDDocument(ctx sdk.Context, did string) types.DIDDocumentWithSeq {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetDIDDocumentKey())
	key := []byte(did)
	bz := store.Get(key)
	if bz == nil {
		return types.DIDDocumentWithSeq{}
	}

	var doc types.DIDDocumentWithSeq
	k.cdc.MustUnmarshalLengthPrefixed(bz, &doc)
	return doc
}

func (k Keeper) ListDIDs(ctx sdk.Context) []string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetDIDDocumentKey())
	dids := make([]string, 0)

	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		did := string(iter.Key())
		dids = append(dids, did)
	}
	return dids
}

func (k Keeper) SetDeactivatedDIDDocument(ctx sdk.Context, creator string, doc types.DIDDocumentWithSeq) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetDIDDocumentDeactivatedKey())
	key := []byte(creator)
	bz := k.cdc.MustMarshalLengthPrefixed(&doc)
	store.Set(key, bz)
}

func (k Keeper) GetDeactivatedDIDDocument(ctx sdk.Context, creator string) types.DIDDocumentWithSeq {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetDIDDocumentDeactivatedKey())
	key := []byte(creator)
	bz := store.Get(key)
	if bz == nil {
		return types.DIDDocumentWithSeq{}
	}

	var doc types.DIDDocumentWithSeq
	k.cdc.MustUnmarshalLengthPrefixed(bz, &doc)
	return doc
}
