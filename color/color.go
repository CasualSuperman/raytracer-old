package color

import (
	"bufio"
	"fmt"
	"math"
	"raytracer/debug"
	"raytracer/log"
)

/* Represents the R, G, and B channels of a color. */
type Color struct {
	R, G, B float64
}

/* Reads from the given reader until it finds a color. */
func (c *Color) Read(r *bufio.Reader) (err error) {
	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, err = fmt.Sscanf(string(line), "%f %f %f", &c.R, &c.G, &c.B)
		}
	}

	if err != nil {
		return err
	}

	if count != 3 {
		return fmt.Errorf("Tried to read a color, only got %d values.", count)
	}

	if debug.INPUT {
		log.Println("Read color:", c)
	}

	return nil
}

/* Scale each component of the color by the given factor. */
func (c *Color) Scale(factor float64) {
	c.R *= factor
	c.G *= factor
	c.B *= factor
}

/* Cap each channel of the color at the given value. */
func (c *Color) Cap(max float64) {
	c.R = math.Min(max, c.R)
	c.G = math.Min(max, c.G)
	c.B = math.Min(max, c.B)
}

/* Return the sum of all channels. Useful for checking if the color is black. */
func (c *Color) Magnitude() float64 {
	return c.R + c.G + c.B
}

/* Return a string representation of the color. */
func (c Color) String() string {
	return fmt.Sprintf("R: %.2f, G: %.2f, B: %.2f", c.R, c.G, c.B)
}
