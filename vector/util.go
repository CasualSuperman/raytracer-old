package vector

import "math"

func Dot(v1, v2 vectorer) float64 {
	t1 := v1.Vector()
	t2 := v2.Vector()
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z
}

func Cross(v1, v2 Vec3) (v3 Vec3) {
	v3.X = v1.Y*v2.Z - v1.Z*v2.Y
	v3.Y = v1.Z*v2.X - v1.X*v2.Z
	v3.Z = v1.X*v2.Y - v1.Y*v2.X
	return
}

func IsZero(num float64) bool {
	return num < 0.00001 && num > -0.00001
}

func length(v *Vec3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
