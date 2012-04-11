package shapes

import "raytracer/vector"

type Shape interface {
	Intersector
	Next() *Shape
}

type Intersector interface {
	Intersect(r vector.Ray) (bool, float64)
}

type ShapeId byte

type shape struct {
	Next *Shape
	Id   int
	Type ShapeId

	Mat Material

	Hit vector.Position
}

type Material struct {
	Ambient, Diffuse, Specular [3]float64
}
