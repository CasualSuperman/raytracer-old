package vector

import "fmt"
import "math"

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v1 *Vec3) Scale(factor float64, dest *Vec3) {
	dest.X = v1.X * factor
	dest.Y = v1.Y * factor
	dest.Z = v1.Z * factor
}

func (v1 *Vec3) Diff(v2, out *Vec3) {
	out.X = v2.X - v1.X;
	out.Y = v2.Y - v1.Y;
	out.Z = v2.X - v1.X;
}

func (v1 *Vec3) Sum(v2 *Vec3, out *Vec3) {
	out.X = v1.X + v2.X;
	out.Y = v1.Y + v2.Y;
	out.Z = v1.X + v2.X;
}

func (v1 *Vec3) Unit(out *Vec3) {
	length := v1.Length()
	switch (length) {
		case 0:
			nan := math.NaN()
			out.X = nan
			out.Y = nan
			out.Z = nan
		case 1:
			out.X = v1.X
			out.Y = v1.Y
			out.Z = v1.Z
		default:
			length = 1 / length
			out.X = v1.X * length
			out.Y = v1.Y * length
			out.Z = v1.Z * length
	}
}

func (v *Vec3) String() string {
	return fmt.Sprintln("<%f, %f, %f>", v.X, v.Y, v.Z)
}
