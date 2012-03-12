package main

import (
	"fmt"
	vec "raytracer/vector"
)

func main() {
	v1 := vec.Vec3{3, 4, 0}
	fmt.Println("Length:", v1.Length())
}
