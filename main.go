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
