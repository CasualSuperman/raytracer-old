package vector

func Dot(d1, d2 Direction) float64 {
	return d1.X * d2.X +
		   d1.Y * d2.Y +
		   d1.Z * d2.Z
}
