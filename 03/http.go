package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	res, _ := http.ReadResponse(bufio.NewReader(conn), nil)
	// header
	fmt.Println(res.Header)
	// body
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
