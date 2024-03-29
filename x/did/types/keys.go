package types

const (
	// ModuleName defines the module name
	ModuleName = "did"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_did"
)

var (
	DIDDocumentKey            = []byte{0x00}
	DIDDocumentDeactivatedKey = []byte{0x01}
)

func GetDIDDocumentKey() []byte {
	return DIDDocumentKey
}

func GetDIDDocumentDeactivatedKey() []byte {
	return DIDDocumentDeactivatedKey
}
