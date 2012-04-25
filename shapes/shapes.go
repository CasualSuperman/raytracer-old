package shapes

import "io"
import "raytracer/vector"

type shapeReader func(io.Reader)(Shape, error)

var types map[int]shapeReader

func init() {
	types = make(map[int]shapeReader)
}

func RegisterFormat(id int, reader shapeReader) {
	types[id] = reader
}

type Shape interface {
	Intersector
	Next() Shape
}

type Intersector interface {
	Intersect(r vector.Ray) (bool, float64)
}

type ShapeId byte

type shape struct {
	Next Shape
	Id   int
	Type ShapeId

	Mat Material

	Hit vector.Position
}

func (s shape) String() string {
	return ""
}

func (s *shape) Read(r io.Reader) error {
	err := s.Mat.Read(r)
	if err != nil {
		return err
	}

	err = s.Hit.Read(r)

	return err
}
