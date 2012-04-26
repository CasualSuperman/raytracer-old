package ppm

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"io"
)

const (
	magicPlainPPM = "P3"
	magicPPM      = "P6"
)

type decoder struct {
	r             io.ByteScanner
	img           image.Image
	width, height int
	maxVal        uint16
}

type decoderP6 decoder

func decodeP6(r io.Reader) (image.Image, error) {
	reader := newByteScanner(r)

	p6 := newP6Decoder(reader)

	err := p6.ReadHeader()

	if err != nil {
		return nil, err
	}
}

func (d *decoderP6) readMagicBytes() error {
	char, err := d.r.ReadByte()

	if err != nil || (char != 'P' && char != 'p') {
		return err || FormatError("Incorrect image format.")
	}

	char, err = d.r.ReadByte()

	if err != nil || char != '6' {
		return err || FormatError("Incorrect image format.")
	}

	return nil
}

func (d *decoderP6) readWidth() (int, error) {
	width := ""

	char, err := d.r.ReadByte()

	for !isWhitespace(char) && err == nil {
		width += char
		char, err = d.r.ReadByte()
	}

	if err != nil {
		return err
	}

	return strconv.Atoi(width)
}

func (d *decoderP6) readHeight() (int, error) {
	height := ""
	char := ""

	char, err = d.r.ReadByte()

	for !isWhitespace(char) && err == nil {
		height += char
		char, err = d.r.ReadByte()
	}

	if err != nil {
		return err
	}

	return strconv.Atoi(height)
}

func (d *decoderP6) readMaxVal() (uint16, error) {
	max := ""
	char := ""

	char, err = d.r.ReadByte()

	for !isWhitespace(char) && err == nil {
		max += char
		char, err = d.r.ReadByte()
	}

	if err != nil {
		return err
	}

	val, err := strconv.Atoi(max)

	if (val >= (1 << 16)) || val < 0 {
		return 0, FormatError("Invalid maximum pixel size.")
	}

	return val, nil
}

func (d *decoderP6) ReadHeader() (err error) {
	err = d.readMagicBytes()

	if err != nil {
		return
	}

	consumeWhitespace(d.r)

	d.width, err = d.readWidth()

	if err != nil {
		return
	}

	consumeWhitespace(d.r)

	d.height, err = d.readHeight()

	if err != nil {
		return
	}

	consumeWhitespace(d.r)

	d.maxVal, err = d.readMaxVal()

	return
}
