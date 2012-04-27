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
	"runtime"
)

type pixel [3]float64
type pixelSegment struct {
	Min, Max struct {
		X, Y int
	}
}

var goroutines = runtime.NumCPU()

func makePixelSegment(done chan bool, m *view.Model, i color.Image, s *pixelSegment) {
	for y := s.Min.Y; y < s.Max.Y; y++ {
		for x := s.Min.X; x < s.Max.X; x++ {
			if debug.IMAGE {
				log.Printf("Calculating pixel (%d, %d)\n", x, y)
			}
			makePixel(m, x, y, i)
		}
	}
	done <- true
}

func MakeImage(m *view.Model) {
	image := color.New(m.Projection.WinSizePixel[0],
		m.Projection.WinSizePixel[1])

	if debug.IMAGE {
		log.Println(*m)
	}

	if debug.ANY {
		// If any debugging is turned on, we want to do this in a single thread.
		// That prevents the log messages from showing up inside each other.
		for y := image.Height() - 1; y >= 0; y-- {
			for x := 0; x < image.Width(); x++ {
				makePixel(m, x, y, image)
			}
		}
	} else {
		// Otherwise, split up the work between the cores!
		work := make([]pixelSegment, goroutines*goroutines)

		for i := 0; i < goroutines; i++ {
			for j := 0; j < goroutines; j++ {
				segment := &work[i*goroutines+j]

				segment.Min.X = image.Width() / goroutines * (i + 0)
				segment.Max.X = image.Width() / goroutines * (i + 1)

				segment.Min.Y = image.Height() / goroutines * (j + 0)
				segment.Max.Y = image.Height() / goroutines * (j + 1)
			}
		}

		done := make(chan bool)

		// Kick off as many threads as we have cores.
		for i := 0; i < goroutines; i++ {
			go makePixelSegment(done, m, image, &work[i])
		}

		// Start a new goroutine every time we get a result back, keep the CPU busy
		for i := goroutines; i < goroutines*goroutines; i++ {
			<-done
			go makePixelSegment(done, m, image, &work[i])
		}

		// Wait for the last few to finish
		for i := 0; i < goroutines; i++ {
			<-done
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

	i.SetPixel(x, y, uint8(p[0]*255), uint8(p[1]*255), uint8(p[2]*255))
}

func rayTrace(m *view.Model, r vector.Ray, p *pixel, dist *float64, last shapes.Shape) {
	closest, nextDist, hit := findClosestObject(m.Shapes, r, nil)

	if closest != nil {
		if debug.RAYTRACE {
			log.Printf("Hit an object. (%d)\n", closest.Id())
		}
		*dist += nextDist
		c := closest.Ambient(&hit.Position)

		p[0] = c.X
		p[1] = c.Y
		p[2] = c.Z

		diffuseIllumination(m, &closest, &hit, p)

		p[0] /= (*dist)
		p[1] /= (*dist)
		p[2] /= (*dist)
	}
}

func mapPixToWorld(m *view.Model, row, col int) (r vector.Ray) {
	p := m.Projection

	pixelWidth := p.WinSizePixel[0]
	pixelHeight := p.WinSizePixel[1]

	worldWidth := p.WinSizeWorld[0]
	worldHeight := p.WinSizeWorld[1]

	r.Position.X = float64(row)/float64(pixelWidth-1)*worldWidth - (worldWidth / 2)

	r.Position.Y = float64(col)/float64(-pixelHeight-1)*worldHeight + (worldHeight / 2)

	r.Position.Z = 0

	r.Direction = *(r.Position.Direction())
	r.Direction.Sub(p.Viewpoint.Direction())
	r.Direction.Unit()

	r.Position = p.Viewpoint

	if debug.RAYTRACE || debug.PIXEL {
		log.Printf("Pixel %d, %d is ray %s\n", row, pixelHeight-col-1, r.String())
	}

	return
}

func findClosestObject(shapes []shapes.Shape, start vector.Ray, base *shapes.Shape) (s shapes.Shape, d float64, r vector.Ray) {
	d = math.Inf(1)
	s = nil
	for _, shape := range shapes {
		if base == nil || shape.Id() != (*base).Id() {
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
				r = *dir
			}
		}
	}
	return
}
