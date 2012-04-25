package shapes

import (
	vec "raytracer/vector"
	"reflect"
)

type Light struct {
	vec.Position
}

func init() {
	// This passes the type literal of Sphere to the function.
	RegisterFormat(10, reflect.TypeOf((*Light)(nil)).Elem())
}


