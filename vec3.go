package main

import "math"

func (v vec3) Length() float64 {
	return math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
}
