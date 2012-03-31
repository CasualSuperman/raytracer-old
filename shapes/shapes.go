package shapes

import "raytracer/vector"

type Intersecter interface {
	Intersect(r vector.Ray) (bool, float64)
}

type ShapeId byte

type Shape struct {
	Next *Shape
	Id   int
	Type ShapeId

	Mat Material

	shape Intersecter

	Hit vector.Position
}

type Material struct {
	Ambient, Diffuse, Specular [3]float64
}

func (s *Shape) Intersect(r vector.Ray) (bool, float64) {
	return s.shape.Intersect(r)
}
