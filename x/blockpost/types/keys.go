package types

const (
	// ModuleName defines the module name
	ModuleName = "blockpost"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blockpost"
)

var (
	ParamsKey = []byte("p_blockpost")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
