package view

import "raytracer/shapes"

type Model struct {
	Projection Projection
	Lights     []shapes.Light
	Shapes     []shapes.Shape
}

func NewModel(proj Projection, lights []shapes.Light, shapes []shapes.Shape) Model {
	return Model{proj, lights, shapes}
}
