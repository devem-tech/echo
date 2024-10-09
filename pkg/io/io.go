package io

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func String(r io.Reader, err error) string {
	if err != nil {
		if errors.Is(err, io.EOF) {
			return ""
		}

		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf(">>> (suppressed) io: read: %w", err))

		return ""
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf(">>> (suppressed) io: read: %w", err))

		return ""
	}

	return string(bytes)
}
