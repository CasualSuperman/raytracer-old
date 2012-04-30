package pshapes

import (
	"math"
	"raytracer/color"
	"raytracer/vector"
)

func sphereRainbows(obj *Psphere, hit *vector.Position, c *color.Color) {
	offset := obj.Sphere.Center.Offset(*hit)

	v1 := offset.X / math.Sqrt(offset.X*offset.X+offset.Y*offset.Y)
	t1 := math.Acos(v1)

	if offset.Y < 0 {
		t1 = 2*math.Pi - t1
	}

	c.R *= (1 + math.Cos(2*t1))
	c.G *= (1 + math.Cos(2*t1+2*math.Pi/3))
	c.B *= (1 + math.Cos(2*t1+4*math.Pi/3))
}

func sphereLines(obj *Psphere, hit *vector.Position, c *color.Color) {
	offset := obj.Sphere.Center.Offset(*hit)

	sum := int(1000 + offset.X*offset.Y*offset.Y/100 +
		offset.X*offset.Y/100)

	if sum&1 == 1 {
		c.R = 0
	} else {
		c.B = 0
	}
}

func rainbows(obj *Pplane, hit *vector.Position, c *color.Color) {
	offset := obj.Plane.Center.Offset(*hit)

	v1 := offset.X / math.Sqrt(offset.X*offset.X+offset.Y*offset.Y)
	t1 := math.Acos(v1)

	if offset.Y < 0 {
		t1 = 2*math.Pi - t1
	}

	c.R *= (1 + math.Cos(2*t1))
	c.G *= (1 + math.Cos(2*t1+2*math.Pi/3))
	c.B *= (1 + math.Cos(2*t1+4*math.Pi/3))
}

func lines(obj *Pplane, hit *vector.Position, c *color.Color) {
	offset := obj.Plane.Center.Offset(*hit)

	sum := int(1000 + offset.X*offset.Y*offset.Y/100 +
		offset.X*offset.Y/100)

	if sum&1 == 1 {
		c.R = 0
	} else {
		c.B = 0
	}
}
