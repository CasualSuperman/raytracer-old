package vector

type Vec3 struct {
	X, Y, Z float64
}

type Direction Vec3
type Position Vec3

type Ray struct {
	Position
	Direction
}

type vectorer interface {
	Vector() Vec3
}

func Dot(v1, v2 vectorer) float64 {
	t1 := v1.Vector()
	t2 := v2.Vector()
	return t1.X*t2.X +
		t1.Y*t2.Y +
		t1.Z*t2.Z
}

func (v *Vec3) Vector() Vec3 {
	return *v
}
