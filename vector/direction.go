package vector

import "fmt"
import "math"

func NewDirection(x, y, z float64) Direction {
	return Direction{x, y, z}
}

func Dot(d1, d2 Direction) float64 {
	return d1.X*d2.X +
		d1.Y*d2.Y +
		d1.Z*d2.Z
}

func (d *Direction) Scale(amount float64) {
	d.X *= amount
	d.Y *= amount
	d.Z *= amount
}

func (d *Direction) Invert() {
	d.X *= -1
	d.Y *= -1
	d.Z *= -1
}

func (d *Direction) Unit() {
	length := d.Length()

	if length != 1 && length != 0 {
		inverse := 1 / length
		d.X *= inverse
		d.Y *= inverse
		d.Z *= inverse
	} else if length == 0 {
		nan := math.NaN()
		d.X = nan
		d.Y = nan
		d.Z = nan
	}
}

func (d *Direction) Length() float64 {
	return math.Sqrt(Dot(*d, *d))
}

func (d *Direction) Copy() Direction {
	return Direction{d.X, d.Y, d.Z}
}

func (d1 *Direction) Add(d2 *Direction) {
	d1.X += d2.X
	d1.Y += d2.Y
	d1.Z += d2.Z
}

func (d1 *Direction) Sub(d2 *Direction) {
	d1.X -= d2.X
	d1.Y -= d2.Y
	d1.Z -= d2.Z
}

func (d *Direction) String() string {
	return fmt.Sprintln("<%f, %f, %f>", d.X, d.Y, d.Z)
}
