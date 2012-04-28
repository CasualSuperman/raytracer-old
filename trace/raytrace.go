package trace

import (
	"math"
	"raytracer/debug"
	"raytracer/log"
	. "raytracer/shapes"
	. "raytracer/vector"
	"raytracer/view"
)

// Given a model, a ray to trace along, a pixel to store the value in, a
// distance travelled, and the shape we hit last, computes the pixel value of
// what the ray hits.
func rayTrace(m *view.Model, r Ray, p *pixel, dist *float64, last Shape) {
	// Find the closest object
	closest, nextDist, hit := findClosestObject(m.Shapes, r, nil)

	// If we have one
	if closest != nil {
		if debug.RAYTRACE {
			log.Printf("Hit an object. (%d)\n", closest.Id())
		}
		// Increase our ditance travelled.
		*dist += nextDist
		// Get the object's ambient light
		c := closest.Ambient(&hit.Position)

		// Initialize our color to it.
		p[0] = c.X
		p[1] = c.Y
		p[2] = c.Z

		// Get the diffuse lighting at that spot as well.
		diffuseIllumination(m, &closest, &hit, p)

		// Then divide by how far we've come, since light has an inverse falloff
		p[0] /= (*dist)
		p[1] /= (*dist)
		p[2] /= (*dist)
	}
}

// Given a pixel location, map it to a physical spot in the world.
func mapPixToWorld(m *view.Model, row, col int) (r Ray) {
	// Start with our projection.
	p := m.Projection

	pWidth := p.Pixel.Width
	pHeight := p.Pixel.Height

	wWidth := p.World.Width
	wHeight := p.World.Height

	r.Position.X = float64(row) / float64(pWidth-1) * wWidth
	r.Position.X -= wWidth / 2
	// Do this backwards so we flip the image.
	r.Position.Y = float64(col) / float64(pHeight-1) * wHeight
	r.Position.Y -= wHeight / 2
	r.Position.Y *= -1

	r.Position.Z = 0

	// Calculate the vector from the viewpoint to this point.
	r.Direction = *(r.Position.Direction())
	r.Direction.Sub(p.Viewpoint.Direction())
	r.Direction.Unit()

	if debug.RAYTRACE || debug.PIXEL {
		log.Println("Pixel", row, pixelHeight-col-1, "ray is", r)
	}

	// But start the ray at the viewpoint.
	r.Position = p.Viewpoint

	return
}

// Given a list of shapes, a place to start, and possibly a shape to ignore,
// find the closest shape to where we start. Return if we hit one, the distance
// to it if we did hit it, and a ray with its position on the shape where we
// hit, and its direction as the normal on the shape at that point.
func findClosestObject(shapes []Shape, start Ray, base *Shape) (s Shape, d float64, r Ray) {
	// Start at positive infinity
	d = math.Inf(1)
	// The closest shape doesn't exist yet.
	s = nil
	for _, shape := range shapes {
		// Skip this if the shape is the one we're ignoring.
		if base == nil || shape.Id() != (*base).Id() {
			if debug.HITS {
				log.Println("Casting ray at shape", shape.Id())
			}
			hit, dist, dir := shape.Hits(start)
			// If we're closer than our last closest object.
			if hit && dist < d {
				s = shape
				d = dist
				r = *dir
				if debug.HITS {
					log.Println("Hit object", shape.Id())
				}
			}
		}
	}
	return
}