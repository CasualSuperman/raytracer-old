package color

import (
	"bufio"
	"fmt"
	"raytracer/debug"
	"raytracer/log"
)

type Color struct {
	R, G, B float64
}

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

func (c *Color) Scale(factor float64) {
	c.R *= factor
	c.G *= factor
	c.B *= factor
}

func (c *Color) Magnitude() float64 {
	return c.R + c.G + c.B
}

func (c Color) String() string {
	return fmt.Sprintf("R: %.2f, G: %.2f, B: %.2f", c.R, c.G, c.B)
}
