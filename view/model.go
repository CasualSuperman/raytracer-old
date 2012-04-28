package view

import "bufio"
import "raytracer/shapes"

type Model struct {
	Projection Projection
	Lights     []shapes.Light
	Shapes     []shapes.Shape
}

func Read(args []string, input *bufio.Reader) (Model, error) {
	// Start with a blank model.
	m := Model{}

	// Read in the projection information.
	err := m.Projection.Read(args, input)

	if err != nil {
		return m, err
	}

	// Allocate slices for the lights and shapes.
	m.Lights = make([]shapes.Light, 0)
	m.Shapes = make([]shapes.Shape, 0)

	return m, nil
}
