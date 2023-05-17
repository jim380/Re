package types

const (
	// ModuleName defines the module name
	ModuleName = "tokenregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenregistry"

    // Version defines the current version the IBC module supports
Version = "tokenregistry-1"

// PortID is the default port id that module binds to
PortID = "tokenregistry"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("tokenregistry-port-")
)

func KeyPrefix(p string) []byte {
    return []byte(p)
}
