package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "fix"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fix"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	// AccountKeyPrefix is the prefix to retrieve all Account
	AccountKeyPrefix = "Account/value/"
)

// AccountKey returns the store key to retrieve a Account from the index fields
func AccountKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
