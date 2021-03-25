package types

const (
	ModuleName = "oracle"
	StoreKey   = ModuleName
)

var (
	// Keys for store prefixes
	RefKey = []byte{0x01}
)

func GetRefStoreKey(symbol string) []byte {
	return append(RefKey, []byte(symbol)...)
}
