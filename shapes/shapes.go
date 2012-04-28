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

// A map of types that we know how to read in.
var shapes = make(map[shapeId]shapeReader)
var lights = make(map[shapeId]lightReader)

// The shape interface.
type Shape interface {
	// Returns the shape's global ID.
	Id() int
	// Return the type of shape that it is.
	Type() shapeId
	// Embed the Intersector interface.
	Intersector
	// Returns the specular at the given point.
	Specular(*vector.Position) vector.Vec3
	// Returns the ambient at the given point.
	Ambient(*vector.Position) vector.Vec3
	// Returns the diffuse at the given point.
	Diffuse(*vector.Position) vector.Vec3
}

// Things that can be intersected will have these methods.
type Intersector interface {
	/* Takes a normalized ray, and returns a boolean that says if it hit, and a
	 * float64 that is the length of the ray when it hits the object.
	 */
	Hits(r vector.Ray) (bool, float64, *vector.Ray)
}

// An alias of int as shapeId, but with actual type safety
type shapeId int

// The basic information all shapes have.
type shape struct {
	// The global shape id.
	id int
	// The shape's material.
	Mat Material
}

// Register a given shapeId with our list of shape formats we understand. It
// will be read in using the provided shapeReader.
func RegisterShapeFormat(id shapeId, reader shapeReader) {
	shapes[id] = reader
}

// Register a given shapeId with our list of light formats we understand. It
// will be read in using the provided lightReader.
func RegisterLightFormat(id shapeId, reader lightReader) {
	lights[id] = reader
}

// Read in shapes and lights from the given input stream, adding them to the
// provded lights and shapes slices. Will stop if any format is not understood
// or a reader returns an error, passing that error back up.
func Read(r *bufio.Reader, s *[]Shape, l *[]Light) (err error) {
	err = nil
	scanning := true

	// While we're trying to scan.
	for scanning {
		count, num, line := 0, 0, []byte{}
		// Read lines until we hit a line that works.
		for err == nil && count == 0 {
			line, _, err = r.ReadLine()
			if err == nil {
				count, _ = fmt.Sscanf(string(line), "%d", &num)
			}
		}

		if err != nil {
			if err != io.EOF {
				// Some other error.
				return fmt.Errorf("Unable to read shape id.")
			} else {
				// We're done here.
				scanning = false
			}
		} else {
			// We can continue. See if we have a compatible reader with the
			// shapeId we got.
			shapeReader, shapeExists := shapes[shapeId(num)]
			lightReader, lightExists := lights[shapeId(num)]

			// If we don't have one, let them know the shape is unsupported.
			if !shapeExists && !lightExists {
				if debug.SHAPES {
					log.Printf("id, read, shapes, lights: %v, %v, %v, %v\n",
						num, count, shapes, lights)
				}
				return fmt.Errorf("Unknown type id %d.", num)
			}

			// If we have a compatible shapeReader
			if shapeExists {
				// Give it our stream and get back the shape.
				shape, err := shapeReader(r)

				// If there was an error, send it up.
				if err != nil {
					return err
				}

				// Otherwise, append it to our list of shapes.
				*s = append(*s, shape)

				// If we have a lightReader with the right Id.
			} else if lightExists {
				// Give it our stream and get back the light.
				light, err := lightReader(r)

				// Quit if we got an error.
				if err != nil {
					return err
				}

				// Otherwise, add it to the lights.
				*l = append(*l, light)
			}
		}
		// Now loop back to finding an input ID if we 're still scanning.
	}

	// If we broke out of the loop, we finished cleanly.
	return nil
}

// Our internal counter for shapes.
var shapeCounter = 0

// Pretty-print shape information.
func (s shape) String() string {
	return fmt.Sprintf("id: %d\n\tMaterial: \n%s", s.id, s.Mat.String())
}

// Return our id.
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
		if err == nil {
			log.Println("Read in material.")
		} else {
			log.Println("Error reading in material.")
		}
	}

	// Return our error, if we have one.
	return err
}
