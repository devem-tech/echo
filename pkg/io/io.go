package io

import (
	"io"
)

func String(r io.Reader, _ error) string {
	bytes, _ := io.ReadAll(r)

	return string(bytes)
}
