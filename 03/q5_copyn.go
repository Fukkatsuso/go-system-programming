package main

import (
	"io"
	"os"
	"strings"
)

func CopyN(dest io.Writer, src io.Reader, length int) {
	srcLimit := io.LimitReader(src, int64(length))
	io.Copy(dest, srcLimit)
}

func main() {
	reader := strings.NewReader("hello world, copyn")

	length := 11
	CopyN(os.Stdout, reader, length)
	// io.CopyN(os.Stdout, reader, int64(length))
}
