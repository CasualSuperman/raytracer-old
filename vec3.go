package main

import "math"

func (v vec3) Length() float64 {
	size := v[0] * v[0]
	size += v[1] * v[1]
	size += v[2] * v[2]
	size = math.Sqrt(size)
	return size
}
