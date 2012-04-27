package color

import (
	"fmt"
	"io"
)

type Pixel []byte

type Image struct {
	base          []byte
	height, width int
}

func New(x, y int) (i Image) {
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
	i.base[((y*i.width)+x)*3+0] = r
	i.base[((y*i.width)+x)*3+1] = g
	i.base[((y*i.width)+x)*3+2] = b
}

func (i Image) Width() int {
	return i.width
}

func (i Image) Height() int {
	return i.height
}

func (i Image) GetPixel(x, y int) Pixel {
	return i.base[((y*i.width)+x)*3 : ((y*i.width)+x+1)*3]
}
