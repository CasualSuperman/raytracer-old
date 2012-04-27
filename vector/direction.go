package vector

import "bufio"
import "fmt"
import "math"
import "raytracer/debug"
import "raytracer/log"

func NewDirection(x, y, z float64) Direction {
	return Direction{x, y, z}
}

func (d *Direction) Scale(amount float64) *Direction {
	d.X *= amount
	d.Y *= amount
	d.Z *= amount
	return d
}

func (d *Direction) Invert() *Direction {
	d.X *= -1
	d.Y *= -1
	d.Z *= -1
	return d
}

func (d *Direction) Unit() *Direction {
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
	return d
}

func (d *Direction) Length() float64 {
	return math.Sqrt(Dot(d, d))
}

func (d Direction) Copy() Direction {
	return Direction{d.X, d.Y, d.Z}
}

func (d1 *Direction) Add(d2 *Direction) *Direction {
	d1.X += d2.X
	d1.Y += d2.Y
	d1.Z += d2.Z
	return d1
}

func (d1 *Direction) Sub(d2 *Direction) *Direction {
	d1.X -= d2.X
	d1.Y -= d2.Y
	d1.Z -= d2.Z
	return d1
}

func (d *Direction) String() string {
	return fmt.Sprintf("<%.3f, %.3f, %.3f>", d.X, d.Y, d.Z)
}

func (d *Direction) Vector() Vec3 {
	return Vec3(*d)
}

func (d *Direction) Position() *Position {
	return (*Position)(d)
}

func (d *Direction) Read(r *bufio.Reader) (err error) {
	err = nil
	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f %f %f", &d.X, &d.Y, &d.Z)
		}
	}

	if err != nil {
		return err
	}

	if count != 3 {
		return fmt.Errorf("Tried to read a position, only got %d values.", count)
	}

	if debug.INPUT {
		log.Println("Read direction:", d)
	}

	return nil
}
