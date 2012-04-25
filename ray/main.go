package main

import (
	"fmt"
	"raytracer/view"
	"os"
)

const (
	EXIT_SUCCESS = iota
	EXIT_BAD_PARAMS
)

func main() {
	model := view.New()

	err := model.LoadProjection(os.Args, os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(EXIT_BAD_PARAMS)
	}

	fmt.Fprintf(os.Stderr, "%s\n", model.Projection.String())
}
