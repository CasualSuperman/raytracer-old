package trace

import (
	"math"
	"os"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/vector"
	"raytracer/view"
)

type pixel [3]float64

func MakeImage(m *view.Model) {
	image := color.New(m.Projection.WinSizePixel[0],
					   m.Projection.WinSizePixel[1])

	if debug.IMAGE {
		log.Println(*m)
	}

	for y := 0; y < image.Height(); y++ {
		for x := 0; x < image.Width(); x++ {
			if debug.IMAGE {
				log.Printf("Calculating pixel (%d, %d)\n", x, y)
			}
			makePixel(m, x, y, image)
		}
	}

	image.PPM(os.Stdout)
}

func makePixel(m *view.Model, x, y int, i color.Image) {
	base := mapPixToWorld(m, x, y)
	p := pixel{0, 0, 0}
	dist := 0.0

	rayTrace(m, base, &p, &dist, nil)

	// Cap the pixel's intensity
	p[0] = math.Min(1, p[0])
	p[1] = math.Min(1, p[1])
	p[2] = math.Min(1, p[2])

	i.SetPixel(x, y, uint8(p[0] * 255), uint8(p[1] * 255), uint8(p[2] * 255))
}

func rayTrace(m *view.Model, r vector.Ray, p *pixel, dist *float64, last shapes.Shape) {
	closest, nextDist, hit := findClosestObject(m.Shapes, r, nil)

	if closest != nil {
		if debug.RAYTRACE {
			log.Printf("Hit an object. (%d)\n", closest.Id())
		}
		*dist += nextDist
		c := closest.Ambient(&hit.Position)

		p[0] = c.X/(*dist)
		p[1] = c.Y/(*dist)
		p[2] = c.Z/(*dist)
	}
}

func mapPixToWorld(m *view.Model, x, y int) (r vector.Ray) {
	p := m.Projection

	pixelWidth := p.WinSizePixel[0]
	pixelHeight := p.WinSizePixel[1]

	worldWidth := p.WinSizeWorld[0]
	worldHeight := p.WinSizeWorld[1]

	r.Position.X = float64(x) / float64(pixelWidth - 1) * worldWidth
	r.Position.X -= worldWidth / 2

	r.Position.Y = float64(y) / float64(-pixelHeight - 1) * worldHeight
	r.Position.Y += worldHeight / 2

	r.Position.Z = 0

	r.Direction = *(r.Position.Direction())
	r.Direction.Sub(p.Viewpoint.Direction())
	r.Direction.Unit()

	r.Position = p.Viewpoint

	if debug.RAYTRACE {
		log.Printf("Pixel %d, %d is ray %s\n", x, y, r.String())
	}

	return
}

func findClosestObject(shapes []shapes.Shape, start vector.Ray, unused shapes.Shape) (s shapes.Shape, d float64, r vector.Ray) {
	d = math.Inf(1)
	s = nil
	for _, shape := range shapes {
		if debug.HITS {
			log.Println("Casting ray at shape", shape.Id())
		}
		hit, dist, dir := shape.Hits(start)
		if hit && dist < d {
			if debug.HITS {
				log.Println("Hit object")
			}
			s = shape
			d = dist
			r = dir
		}
	}
	return
}
