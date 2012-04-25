package ppm

import (
	"io"
)

const (
	unreadErrorInstance unreadError = ' '
)

type unreadError byte

func (e unreadError) Error() string {
	return "Cannot unread twice in a row."
}

type byteScanner struct {
	r		io.Reader
	last	byte
	unread	bool
}

func (b *byteScanner) UnreadByte() error {
	if b.unread {
		return unreadErrorInstance
	}

	b.unread = true
}

func (b *byteScanner) ReadByte() (c byte, err error) {
	if b.unread {
		c = b.last
	} else {
		char := []byte{' '}
		_, err := b.r.Read(char)
		b.last, c = char[0], char[0]
	}
	return
}

func newByteScanner(r io.Reader) byteScanner {
	return byteScanner{r, ' ', false}
}
