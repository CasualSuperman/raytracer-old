package color

import (
	"fmt"
	"io"
)

type Image struct {
	base          []byte
	height, width int
}

type Pixel *struct {
	R, G, B *byte
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

func (i Image) Width() int {
	return i.width
}

func (i Image) Height() int {
	return i.height
}

func (i Image) GetPixel(x, y int) Pixel {
	start := ((y * i.width) + x) * 3
	var p struct {
		R, G, B *byte
	}
	p.R = &i.base[start+0]
	p.G = &i.base[start+1]
	p.B = &i.base[start+2]
	return Pixel(&p)
}
