package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("duration[sec]?")
	}

	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	duration := time.Duration(d) * time.Second

	ch := time.After(duration)
	after := <-ch
	fmt.Printf("%d seconds passed: %v", d, after)
}
