package trace

import (
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/vector"
	"raytracer/view"
)

func diffuseIllumination(m *view.Model, obj *shapes.Shape, hit *vector.Ray, color *pixel) {
	if debug.LIGHTS {
		log.Println()
	}

	diff := (*obj).Diffuse(&hit.Position)

	if !vector.IsZero(diff.Length()) {
		for _, light := range m.Lights {
			if debug.LIGHTS {
				log.Println("Testing against light", light.Id())
			}
			if debug.DIFFUSE {
				log.Println()
			}
			processLight(m.Shapes, obj, hit, &light, color)
		}
	}
}

func processLight(s []shapes.Shape, obj *shapes.Shape, hit *vector.Ray, l *shapes.Light, color *pixel) {

	directionToLight := hit.Position.Offset(l.Center)

	lightDistance := directionToLight.Length()

	unitToLight := directionToLight
	unitToLight.Unit()

	r := vector.Ray{hit.Position, unitToLight}

	shapeColor := (*obj).Diffuse(&hit.Position)

	cos := vector.Dot(&hit.Direction, &r.Direction)

	if debug.DIFFUSE {
		log.Println("Hit object type was    ", (*obj).Type())
		log.Println("Hit object id was      ", (*obj).Id())
		log.Println("Hit point was          ", hit.Position)
		log.Println("Normal at hitpoint     ", hit.Direction)
		log.Println("Light object id        ", l.Id())
		log.Println("Light center           ", l.Center)
		log.Println("Distance to light is   ", lightDistance)
		log.Println("Unit vector to light is", r.Direction)
		log.Println("cosine is              ", cos)
		log.Println("Emissivity of the light", l.Color)
		log.Println("Diffuse Reflectivity   ", shapeColor)
		log.Println("Current ivec           ", color)
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
				log.Printf("It was closer than the light (%f < %f).\n", nextDist, lightDistance)
			}
			return
		}
	}

	base := [3]float64{0, 0, 0}

	base[0] = l.Color.X * shapeColor.X * cos / lightDistance
	base[1] = l.Color.Y * shapeColor.Y * cos / lightDistance
	base[2] = l.Color.Z * shapeColor.Z * cos / lightDistance

	if debug.DIFFUSE {
		log.Println("Scaled reflectivity   ", base)
	}

	color[0] += base[0]
	color[1] += base[1]
	color[2] += base[2]
}
