package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{"Hello": "World"}

	// json encode
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(source)

	// json gzip writer
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	// gzipWriter.Header.Name = "response.txt.gz"

	// response gzip, and print json
	writer := io.MultiWriter(gzipWriter, os.Stdout)
	writer.Write(buffer.Bytes())
	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
