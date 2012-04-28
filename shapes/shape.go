package shapes

import (
	"bufio"
	"fmt"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/vector"
)

// Our internal counter for shapes.
var shapeCounter = 0

// The basic information all shapes have.
type BaseShape struct {
	// The global shape id.
	id int
	// The shape's material.
	Mat Material
}

// Pretty-print shape information.
func (s *BaseShape) String() string {
	return fmt.Sprintf("id: %d\n\tMaterial: \n%s", s.id, s.Mat.String())
}

// Return our id.
func (s *BaseShape) Id() int {
	return s.id
}

// Read in a shape from the given io.Reader, return an error on failure.
func (s *BaseShape) Read(r *bufio.Reader) error {
	if debug.SHAPES {
		log.Println("Reading in a shape.")
	}
	// Give our shape an Id and increment it.
	s.id = shapeCounter
	shapeCounter++

	// Read in our material.
	err := s.Mat.Read(r)
	if debug.SHAPES {
		if err == nil {
			log.Println("Read in material.")
		} else {
			log.Println("Error reading in material.")
		}
	}

	// Return our error, if we have one.
	return err
}

// Give shapes default ambient functions.
func (s *BaseShape) Ambient(d *vector.Position) color.Color {
	return s.Mat.Ambient
}

// Give shapes default diffuse functions.
func (s *BaseShape) Diffuse(d *vector.Position) color.Color {
	return s.Mat.Diffuse
}

// Give shapes default specular functions.
func (s *BaseShape) Specular(d *vector.Position) color.Color {
	return s.Mat.Specular
}
