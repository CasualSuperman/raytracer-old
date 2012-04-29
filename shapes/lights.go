package shapes

import (
	"raytracer/color"
	"raytracer/vector"
)

type Light interface {
	Id() int
	Color() color.Color
	Position() vector.Position
	Illuminated(*vector.Position) bool
}
