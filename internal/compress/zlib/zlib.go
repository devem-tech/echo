package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Decompress(x string) string {
	r, _ := zlib.NewReader(bytes.NewReader([]byte(x)))

	b, _ := io.ReadAll(r)

	return string(b)
}
