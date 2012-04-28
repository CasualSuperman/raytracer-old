package view

import (
	"bufio"
	"fmt"
	vec "raytracer/vector"
	"strconv"
)

const (
	// The number of parameters we require.
	requiredParameters = 3
)

type Projection struct {
	Pixel struct {
		Height, Width int
	}
	World struct {
		Height, Width float64
	}
	Viewpoint vec.Position
}

func (p *Projection) Read(args []string, in *bufio.Reader) (err error) {
	err = loadProjectionPixels(p, args)

	if err != nil {
		return
	}

	err = loadProjectionWorld(p, in)

	if err != nil {
		return fmt.Errorf("Unable to load projection, error: %s",
			err.Error())
	}

	err = loadProjectionViewPoint(p, in)

	if err != nil {
		return fmt.Errorf("Unable to load viewpoint, error: %s",
			err.Error())
	}

	return
}

func loadProjectionPixels(proj *Projection, args []string) error {
	if len(args) != requiredParameters {
		return fmt.Errorf("usage:\n\t%s width height", args[0])
	}

	width, errW := strconv.Atoi(args[1])
	height, errH := strconv.Atoi(args[2])

	if errH != nil || errW != nil {
		if errH != nil && errW != nil {
			return fmt.Errorf("Unable to parse height or width.")
		} else if errH != nil {
			return fmt.Errorf("Unable to parse height.")
		} else {
			return fmt.Errorf("Unable to parse width.")
		}
	}

	if height <= 0 || width <= 0 {
		return fmt.Errorf("Width and height must be positive.")
	}

	(*proj).Pixel.Width, (*proj).Pixel.Height = width, height

	return nil
}

func loadProjectionWorld(proj *Projection, input *bufio.Reader) (err error) {
	num := 0

	// Loop until we get some number of values read.
	for num == 0 {
		line, prefix, err := input.ReadLine()
		var next []byte

		for prefix && err != nil {
			next, prefix, err = input.ReadLine()
			line = append(line, next...)
		}

		if err != nil {
			return err
		}

		num, err = fmt.Sscanf(string(line), "%f %f", &(*proj).World.Width,
			&(*proj).World.Height)
	}

	if num != 2 && err == nil {
		err = fmt.Errorf("Tried to read 2 values, read %d instead.", num)
	}

	return err
}

func loadProjectionViewPoint(proj *Projection, input *bufio.Reader) (err error) {
	num := 0

	// Loop until we get some number of values read.
	for num == 0 {
		var next []byte
		line, prefix, err := input.ReadLine()

		for prefix && err != nil {
			next, prefix, err = input.ReadLine()
			line = append(line, next...)
		}

		if err != nil {
			return err
		}

		num, err = fmt.Sscanf(string(line), "%f %f %f", &(*proj).Viewpoint.X,
			&(*proj).Viewpoint.Y, &(*proj).Viewpoint.Z)
	}

	if num != 3 && err == nil {
		err = fmt.Errorf("Tried to read 3 values, read %d instead.", num)
	}

	return err
}

func (p *Projection) String() string {
	return fmt.Sprintf("Projection:"+
		"\n\tPixels: %v"+
		"\n\tWorld size: %v"+
		"\n\tViewpoint: %s",
		p.Pixel, p.World,
		p.Viewpoint.String())
}
