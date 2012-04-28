/*
A raytracer written in go.
Supports planes, spheres, finite planes, tiled planes, diffuse and specular
lighting.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/trace"
	"raytracer/view"
	"runtime"
)

// Import these packages for their side-effects (namely, type registration in
// the shapes module.)
import _ "raytracer/shapes/fplane"
import _ "raytracer/shapes/tplane"
import _ "raytracer/shapes/plane"
import _ "raytracer/shapes/sphere"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	stdin := bufio.NewReader(os.Stdin)

	model, err := view.Read(os.Args, stdin)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", model.Projection.String())

	err = shapes.Read(stdin, &model.Shapes, &model.Lights)

	if err != nil {
		log.Fatalln("Unable to read shapes.", err)
	}

	for _, light := range model.Lights {
		log.Println(light)
	}

	for _, shape := range model.Shapes {
		log.Println(shape)
	}

	trace.MakeImage(&model)
}
