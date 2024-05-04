package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)

	smap.Delete("test")

	value, ok := smap.Load("hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)

	// これはすでに保存されているので無視
	smap.LoadOrStore(1, 3)
	// これは保存されていないので挿入
	smap.LoadOrStore(2, 4)

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}
