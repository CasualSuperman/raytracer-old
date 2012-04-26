package view

import "bufio"
import "raytracer/shapes"

type Model struct {
	Projection Projection
	Lights     []shapes.Light
	Shapes     []shapes.Shape
}

func New() Model {
	return Model{}
}

func (m *Model) LoadProjection(args []string, input *bufio.Reader) (err error) {
	m.Projection, err = newProjection(args, input)
	return
}
