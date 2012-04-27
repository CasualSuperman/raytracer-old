package vector

import "bufio"
import "fmt"
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

func IsZero(num float64) bool {
	return num < 0.00001 && num > -0.00001
}

func (v *Vec3) Vector() Vec3 {
	return *v
}

func (v *Vec3) String() string {
	return fmt.Sprintf("{%.3f, %.3f, %.3f}", v.X, v.Y, v.Z)
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
