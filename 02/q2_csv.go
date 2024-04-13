package main

import (
	"encoding/csv"
	"os"
)

func main() {
	encoder := csv.NewWriter(os.Stdout)
	encoder.Write([]string{"one", "two", "three"})
	encoder.Flush()
}
