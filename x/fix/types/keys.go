package types

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

var (
	AccountKey      = []byte{0x00}
	AccountCountKey = []byte{0x01}
)

func GetAccountKey() []byte {
	return AccountKey
}

func GetAccountCountKey() []byte {
	return AccountCountKey
}
