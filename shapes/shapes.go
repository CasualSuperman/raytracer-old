package shapes

import "io"
import "raytracer/vector"
import "reflect"

var types map[int]reflect.Type

func init() {
	types = make(map[int]reflect.Type)
}

func RegisterFormat(id int, object reflect.Type) {
	types[id] = object
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
