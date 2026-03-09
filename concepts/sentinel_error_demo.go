// Demo: sentinel value such as zip.ErrFormat
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello, this is not zip data")
	reader := bytes.NewReader(data)
	_, err := zip.NewReader(reader, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Not a zip archive")
	}
}
