package main

import (
	"bufio"
	"fmt"
	"os"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/trace"
	"raytracer/view"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	model := view.New()

	err := model.LoadProjection(os.Args, stdin)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", model.Projection.String())

	err = shapes.Read(stdin, &model.Shapes, &model.Lights)

	if err != nil {
		log.Fatalln("Unable to read shapes.", err)
	}

	for _, shape := range model.Shapes {
		log.Println(shape)
	}

	trace.MakeImage(&model)
}
