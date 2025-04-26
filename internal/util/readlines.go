package util

import (
	"io"
	"strings"
)

// Read lines, preserving the newline
func ReadLines(r io.Reader) ([]string, error) {
	var buf [4096]byte
	var builder strings.Builder
	result := []string{}

	for {
		n, err := r.Read(buf[:])
		for _, c := range buf[:n] {
			builder.WriteByte(c)
			if c == '\n' {
				result = append(result, builder.String())
				builder.Reset()
			}
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			if builder.Len() > 0 {
				result = append(result, builder.String())
			}
			return result, err
		}
	}
}
