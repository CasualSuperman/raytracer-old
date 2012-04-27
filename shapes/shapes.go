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
type lightReader func(*bufio.Reader) (Light, error)

// A list of types that we know how to read in.
var shapes = make(map[shapeId]shapeReader)
var lights = make(map[shapeId]lightReader)

// Shapes can be intersected
type Shape interface {
	Specular(*vector.Position) vector.Vec3
	Ambient(*vector.Position) vector.Vec3
	Diffuse(*vector.Position) vector.Vec3
	Intersector
	Id() int
	Type() shapeId
}

// Things that can be intersected will have these methods.
type Intersector interface {
	/* Takes a normalized ray, and returns a boolean that says if it hit, and a
	 * float64 that is the length of the ray when it hits the object.
	 */
	Hits(r vector.Ray) (bool, float64, *vector.Ray)
}

// An alias of bytes to shapeId, but with actual type safety
type shapeId int

type shape struct {
	// The global shape id.
	id int
	// The shape type.
	Type shapeId
	// The shape's material.
	Mat Material
	// Where we were last hit.
	Hit vector.Position
}

// Register a format with our list of readable formats.
func RegisterShapeFormat(id shapeId, reader shapeReader) {
	shapes[id] = reader
}

func RegisterLightFormat(id shapeId, reader lightReader) {
	lights[id] = reader
}

func Read(r *bufio.Reader, s *[]Shape, l *[]Light) (err error) {
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
			shapeReader, shapeExists := shapes[shapeId(num)]
			lightReader, lightExists := lights[shapeId(num)]

			if !shapeExists && !lightExists {
				if debug.SHAPES {
					log.Printf("id, readCount, shapes, lights: %v, %v, %v,\n%v\n", num, count, shapes, lights)
				}
				return fmt.Errorf("Unknown type id %d.", num)
			}

			if shapeExists {
				shape, err := shapeReader(r)

				if err != nil {
					return err
				}

				*s = append(*s, shape)
			} else if lightExists {
				light, err := lightReader(r)

				if err != nil {
					return err
				}

				*l = append(*l, light)
			}
		} else {
			// We ran into an error.
			if err != io.EOF {
				// Some other error.
				return fmt.Errorf("Unable to read shape id.")
			} else {
				// We're done here.
				scanning = false
			}
		}
	}

	return nil
}

// Our internal counter for shapes.
var shapeCounter = 0

// Pretty-print shape information.
func (s shape) String() string {
	return fmt.Sprintf("id: %d\n\tMaterial: \n%s", s.id, s.Mat.String())
}

func (s shape) Id() int {
	return s.id
}

// Read in a shape from the given io.Reader, return an error on failure.
func (s *shape) Read(r *bufio.Reader) error {
	if debug.SHAPES {
		log.Println("Reading in a shape.")
	}
	// Give our shape an Id and increment it.
	s.id = shapeCounter
	shapeCounter++

	// Read in our material.
	err := s.Mat.Read(r)
	if debug.SHAPES {
		log.Println("Read in material.")
	}

	return err
}
