package types

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	MultiChainTxOracleKey      = []byte{0x01}
	MultiChainTxOracleCountKey = []byte{0x02}
)

func GetMultiChainTxOracleKey() []byte {
	return MultiChainTxOracleKey
}

func GetMultiChainTxOracleCountKey() []byte {
	return MultiChainTxOracleCountKey
}
