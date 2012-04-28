package vector

import "bufio"
import "fmt"
import "math"
import "raytracer/debug"
import "raytracer/log"

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

func Cross(v1, v2 vectorer) (v3 Vec3) {
	u1 := v1.Vector()
	u2 := v2.Vector()
	u1.Unit()
	u2.Unit()
	v3.X = u1.Y*u2.Z - u1.Z*u2.Y
	v3.Y = u1.Z*u2.X - u1.X*u2.Z
	v3.Z = u1.X*u2.Y - u1.Y*u2.X
	return
}

func IsZero(num float64) bool {
	return num < 0.00001 && num > -0.00001
}

func (v *Vec3) Vector() Vec3 {
	return *v
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(Dot(v, v))
}

func (v *Vec3) Unit() *Vec3 {
	length := v.Length()

	if length != 1 && length != 0 {
		inverse := 1 / length
		v.X *= inverse
		v.Y *= inverse
		v.Z *= inverse
	} else if length == 0 {
		nan := math.NaN()
		v.X = nan
		v.Y = nan
		v.Z = nan
	}
	return v
}

func (v *Vec3) String() string {
	return fmt.Sprintf("{%.4f, %.4f, %.4f}", v.X, v.Y, v.Z)
}

func (v *Vec3) Read(r *bufio.Reader) (err error) {
	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, err = fmt.Sscanf(string(line), "%f %f %f", &v.X, &v.Y, &v.Z)
		}
	}

	if err != nil {
		return err
	}

	if count != 3 {
		return fmt.Errorf("Tried to read a vector, only got %d values.", count)
	}

	if debug.INPUT {
		log.Println("Read vector:", v)
	}

	return nil
}
