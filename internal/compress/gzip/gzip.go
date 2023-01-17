package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Decompress(x string) string {
	r, _ := gzip.NewReader(bytes.NewReader([]byte(x)))

	b, _ := io.ReadAll(r)

	return string(b)
}
