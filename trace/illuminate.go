package trace

import (
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/vector"
	"raytracer/view"
)

func diffuseIllumination(m *view.Model, obj *shapes.Shape, hit *vector.Ray, color *pixel) {
	for _, light := range m.Lights {
		processLight(m.Shapes, obj, hit, &light, color)
	}
}

func processLight(s []shapes.Shape, obj *shapes.Shape, hit *vector.Ray, l *shapes.Light, color *pixel) {

	directionToLight := hit.Position.Offset(l.Center)

	if vector.Dot(&hit.Direction, &directionToLight) <= 0 {
		if debug.DIFFUSE {
			log.Println("Object self-occluded.")
		}
		return
	}

	r := vector.Ray{hit.Position, directionToLight}

	lightDistance := directionToLight.Length()

	r.Direction.Unit()

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

	shapeColor := (*obj).Diffuse(&hit.Position)

	cos := vector.Dot(&hit.Direction, &r.Direction)

	if debug.DIFFUSE {
		log.Println("Hit object id was ", (*obj).Id())
		log.Println("Hit point was", hit.Position)
		log.Println("Normal at hitpoint", hit.Direction)
		log.Println("Light object id", l.Id())
		log.Println("Light center", l.Center)
		log.Println("Distance to light is", lightDistance)
		log.Println("Unit vector to light is", r.Direction)
		log.Println("cosine is", cos)
		log.Println("Emissivity of the light", l.Color)
		log.Println("Diffuse Reflectivity", shapeColor)
		log.Println("Current ivec", color)
	}

	color[0] += l.Color.X * shapeColor.X * cos / lightDistance
	color[1] += l.Color.Y * shapeColor.Y * cos / lightDistance
	color[2] += l.Color.Z * shapeColor.Z * cos / lightDistance
}
