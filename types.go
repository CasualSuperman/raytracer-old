package main

type vec3 [3]float64
type vec2 [2]float64

type vector interface {
	Length() float64
	String() string
}
