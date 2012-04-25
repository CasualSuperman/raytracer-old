package ppm

import (
	"io"
)

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t'
}

func consumeWhitespace(r io.ByteScanner) (int, error) {
	total := 0
	char, err := r.ReadByte()
	for isWhitespace(char) && err == nil  {
		total++
		read, err = r.ReadByte()
	}
	if !isWhitespace(char) {
		r.UnreadByte()
	}
	return total, err
}
