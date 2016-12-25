package sdetools

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
)

func gobToBytes(v interface{}) []byte {
	bb := bytes.NewBuffer([]byte{})
	e := gob.NewEncoder(bb)
	err := e.Encode(v)
	if err != nil {
		log.Fatal(err)
	}

	return bb.Bytes()
}

func boltKey(key int) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(key))
	return bs
}
