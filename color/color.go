package color

import (
	"fmt"
	"io"
)

type Image struct {
	base []byte
	height, width int
}

func NewImage(x, y int) (i Image) {
	i.base = make([]byte, x*y*3)
	i.height = y
	i.width = x
	return
}

func (i Image) PPM(w io.Writer) {
	fmt.Fprintln(w, "P6")
	fmt.Fprintf(w, "%d %d\n", i.width, i.height)
	fmt.Fprintln(w, "255")
	w.Write(i.base)
}

func (i Image) SetPixel(x, y int, r, g, b uint8) {
	i.base[((x * i.width) + y) * 3 + 0] = r
	i.base[((x * i.width) + y) * 3 + 1] = g
	i.base[((x * i.width) + y) * 3 + 2] = b
}
