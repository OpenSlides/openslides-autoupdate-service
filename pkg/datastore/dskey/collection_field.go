package dskey

type collectionField struct {
	collection string
	field      string
}

//go:generate  sh -c "go run gen_collection_fields/main.go > gen_collection_fields.go"

func splitUInt64(i uint64) (int, int) {
	return int(i & 0xffffffff), int(i >> 32)
}

func joinInt(i1, i2 int) uint64 {
	return (uint64(i2) << 32) | uint64(i1)
}
