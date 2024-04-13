package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	readerA := io.NewSectionReader(programming, 5, 1)
	readerS := io.NewSectionReader(system, 0, 1)
	readerC := io.NewSectionReader(computer, 0, 1)
	readerI1 := io.NewSectionReader(programming, 8, 1)
	readerI2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(readerA, readerS, readerC, readerI1, readerI2)
	io.Copy(os.Stdout, stream)
}
