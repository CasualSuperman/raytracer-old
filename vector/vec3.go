package vector

import (
	"bufio"
	"fmt"
	"math"
	"raytracer/debug"
	"raytracer/log"
)

func (v *Vec3) Vector() Vec3 {
	return *v
}

func (v *Vec3) Length() float64 {
	return length(v)
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
