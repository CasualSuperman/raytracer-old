package shapes

import (
	"fmt"
	"io"
	"raytracer/vector"
)

// A function pointer type.
type shapeReader func(io.Reader) (Shape, error)

// A list of types that we know how to read in.
var types = make(map[shapeId]shapeReader)

// Shapes can be intersected
type Shape interface {
	Intersector
}

// Things that can be intersected will have these methods.
type Intersector interface {
	/* Takes a normalized ray, and returns a boolean that says if it hit, and a
	 * float64 that is the length of the ray when it hits the object.
	 */
	Hits(r vector.Ray) (bool, float64, vector.Ray)
}

// An alias of bytes to shapeId, but with actual type safety
type shapeId byte

type shape struct {
	// The global shape id.
	Id int
	// The shape type.
	Type shapeId
	// The shape's material.
	Mat Material
	// Where we were last hit.
	Hit vector.Position
}

// Register a format with our list of readable formats.
func RegisterFormat(id shapeId, reader shapeReader) {
	types[id] = reader
}

// Our internal counter for shapes.
var shapeCounter = 0

// Pretty-print shape information.
func (s shape) String() string {
	return fmt.Sprintf("\tid: %d\n\tMaterial: %s\n", s.Id, s.Mat.String())
}

// Read in a shape from the given io.Reader, return an error on failure.
func (s *shape) Read(r io.Reader) error {
	// Give our shape an Id and increment it.
	s.Id = shapeCounter
	shapeCounter++

	// Read in our material.
	err := s.Mat.Read(r)
	return err
}
