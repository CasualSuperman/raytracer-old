package vector

type Vec3 struct {
	X, Y, Z float64
}

type Direction Vec3
type Position  Vec3

type Ray struct {
	Position
	Direction
}
