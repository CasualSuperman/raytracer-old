package color

import (
	"fmt"
	"io"
)

/* Used to create an image. Has a height and a width. */
type Image struct {
	base          []byte
	height, width int
}

/* A type that can point into the Image's internal byte array. */
type Pixel *struct {
	R, G, B *byte
}

/* Create a new image with the given dimensions. */
func New(x, y int) (i Image) {
	i.base = make([]byte, x*y*3)
	i.height = y
	i.width = x
	return
}

/* Write the image to the given io.Write in PPM format. */
func (i Image) PPM(w io.Writer) {
	fmt.Fprintln(w, "P6")
	fmt.Fprintf(w, "%d %d\n", i.width, i.height)
	fmt.Fprintln(w, "255")
	w.Write(i.base)
}

/* Return the image's width. */
func (i Image) Width() int {
	return i.width
}

/* Return the image's height. */
func (i Image) Height() int {
	return i.height
}

// Get a pixel with references into the image, modifying this pixel updates the
// image.
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
