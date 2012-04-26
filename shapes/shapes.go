package shapes

import (
	"bufio"
	"fmt"
	"io"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/vector"
)

// A function pointer type.
type shapeReader func(*bufio.Reader) (Shape, error)

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
type shapeId int

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

func Read(r *bufio.Reader) (shapes []Shape, err error) {
	err = nil
	scanning := true

	for scanning {
		count, num, line := 0, 0, []byte{}
		// Read lines until we hit a line that works.
		for err == nil && count == 0 {
			line, _, err = r.ReadLine()
			if err == nil {
				count, _ = fmt.Sscanf(string(line), "%d", &num)
			}
		}

		if err == nil {
			// We can continue.
			reader, exists := types[shapeId(num)]

			if !exists {
				if debug.SHAPES {
					log.Printf("Exists, id, readCount, shapes: %v, %v, %v,\n%v\n", exists, num, count, types)
				}
				return nil, fmt.Errorf("Unknown type id %d.", num)
			}

			shape, err := reader(r)

			if err != nil {
				return nil, err
			}

			shapes = append(shapes, shape)
		} else {
			// We ran into an error.
			if err != io.EOF {
				// Some other error.
				return nil, fmt.Errorf("Unable to read shape id.")
			} else {
				// We're done here.
				scanning = false
			}
		}
	}

	return shapes, nil
}

// Our internal counter for shapes.
var shapeCounter = 0

// Pretty-print shape information.
func (s shape) String() string {
	return fmt.Sprintf("id: %d\n\tMaterial: \n%s", s.Id, s.Mat.String())
}

// Read in a shape from the given io.Reader, return an error on failure.
func (s *shape) Read(r *bufio.Reader) error {
	if debug.SHAPES {
		log.Println("Reading in a shape.")
	}
	// Give our shape an Id and increment it.
	s.Id = shapeCounter
	shapeCounter++

	// Read in our material.
	err := s.Mat.Read(r)
	if debug.SHAPES {
		log.Println("Read in material.")
	}

	return err
}
