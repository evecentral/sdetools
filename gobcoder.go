package sdetools

import (
	"encoding/binary"
)

func boltKey(key int) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(key))
	return bs
}
