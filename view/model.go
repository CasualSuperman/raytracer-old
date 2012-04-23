package view

import "io"
import "raytracer/shapes"

type Model struct {
	Projection Projection
	Lights     []shapes.Light
	Shapes     []shapes.Shape
}

func New() Model {
	return Model{}
}

func (m *Model) LoadProjection(args []string, input io.Reader) (err error) {
	m.Projection, err = newProjection(args, input)
	return
}
