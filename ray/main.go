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

	shapes, err := shapes.Read(stdin)

	if err != nil {
		log.Fatalln("Unable to read shapes.", err)
	}

	for _, shape := range shapes {
		log.Println(shape)
	}
}
