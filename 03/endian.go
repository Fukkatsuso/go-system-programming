package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32 bit, big endian (10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	// 変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
