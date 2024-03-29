package trace

import (
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/vector"
	"raytracer/view"
)

// Finds the diffuse illumination at a point on an object.
// m: The viewModel of the scene.
// obj: The shape we hit.
// hit: A ray with its position on the object and its direction on the normal.
// color: The color to return the diffuse illumination on.
func diffuseIllumination(m *view.Model, obj shapes.Shape, hit *vector.Ray,
	p *color.Color) {

	if debug.LIGHTS {
		log.Println()
	}

	diff := obj.Diffuse(&hit.Position)

	if diff.Magnitude() > 0 {
		for _, light := range m.Lights {
			if debug.LIGHTS {
				log.Println("Testing against light", light.Id())
			}
			if debug.DIFFUSE {
				log.Println()
			}
			if light.Illuminated(&hit.Position) {
				processLight(m.Shapes, obj, hit, light, p)
			}
		}
	}
}

// Called for each light when processing illumination. Takes a list of shapes
// to look at, a shape to potentially ignore, a hit location, a light to test
// occlusion for, and a color to apply the light to.
func processLight(s []shapes.Shape, obj shapes.Shape, hit *vector.Ray,
	l shapes.Light, p *color.Color) {

	directionToLight := hit.Position.Offset(l.Position())

	lightDistance := directionToLight.Length()

	unitToLight := directionToLight
	unitToLight.Unit()

	r := vector.Ray{hit.Position, unitToLight}

	shapeColor := obj.Diffuse(&hit.Position)

	cos := vector.Dot(&hit.Direction, &r.Direction)

	if debug.DIFFUSE {
		log.Println("Hit object type was    ", obj.Type())
		log.Println("Hit object id was      ", obj.Id())
		log.Println("Hit point was          ", hit.Position)
		log.Println("Normal at hitpoint     ", hit.Direction)
		log.Println("Light object id        ", l.Id())
		log.Println("Light center           ", l.Position())
		log.Println("Distance to light is   ", lightDistance)
		log.Println("Unit vector to light is", r.Direction)
		log.Println("cosine is              ", cos)
		log.Println("Emissivity of the light", l.Color())
		log.Println("Diffuse Reflectivity   ", shapeColor)
		log.Println("Current ivec           ", p)
	}

	if vector.Dot(&hit.Direction, &directionToLight) <= 0 {
		if debug.DIFFUSE {
			log.Println("Object self-occluded.")
		}
		return
	}

	closest, nextDist, _ := findClosestObject(s, r, obj)

	if closest != nil {
		if debug.DIFFUSE {
			log.Println("Hit an object with id", closest.Id())
		}
		if nextDist < lightDistance {
			if debug.DIFFUSE {
				log.Printf("It was closer than the light (%f < %f).\n",
					nextDist, lightDistance)
			}
			return
		}
	}

	base := color.Color{0, 0, 0}

	base.R = l.Color().R * shapeColor.R * cos / lightDistance
	base.G = l.Color().G * shapeColor.G * cos / lightDistance
	base.B = l.Color().B * shapeColor.B * cos / lightDistance

	if FOG_ENABLED {
		addFog(&base, lightDistance)
	}

	if debug.DIFFUSE {
		log.Println("Scaled reflectivity   ", base)
	}

	p.R += base.R
	p.G += base.G
	p.B += base.B
}
