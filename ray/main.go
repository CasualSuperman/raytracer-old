package main

import (
	"fmt"
//	vec "raytracer/vector"
//	"raytracer/shapes"
	"raytracer/view"
	"os"
)

const (
	EXIT_SUCCESS = iota
	EXIT_BAD_PARAMS
)

func main() {
	proj, err := view.NewProjection(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(EXIT_BAD_PARAMS)
	}

	view.NewModel(proj, nil, nil)
}
