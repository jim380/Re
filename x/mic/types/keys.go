package types

const (
	// ModuleName defines the module name
	ModuleName = "mic"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mic"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var MarketIdentificationCodeKeyPrefix = "MarketIdentificationCode/value/"

func GetMarketIdentificationCodeKey(mic string) []byte {
	return append([]byte(mic), []byte("/")...)
}
