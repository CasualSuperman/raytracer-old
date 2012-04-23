package view

import (
	"bufio"
	"fmt"
	"io"
	vec "raytracer/vector"
	"strconv"
)

const (
	TWO_DIMENSIONAL    = 2
	THREE_DIMENSIONAL  = 3
	NUM_REQ_PARAMETERS = 3
)

type Projection struct {
	WinSizePixel [TWO_DIMENSIONAL]int
	WinSizeWorld [TWO_DIMENSIONAL]float64
	Viewpoint     vec.Position
}

func newProjection(args []string, input io.Reader) (p Projection, err error) {
	err = loadProjectionPixels(&p, args)
	in := bufio.NewReader(input)

	if err != nil {
		return
	}

	err = loadProjectionWorld(&p, in)

	if err != nil {
		return p, fmt.Errorf("Unable to load projection, error: %s",
			err.Error())
	}

	err = loadProjectionViewPoint(&p, in)

	if err != nil {
		return p, fmt.Errorf("Unable to load viewpoint, error: %s",
			err.Error())
	}

	return
}

func loadProjectionPixels(proj *Projection, args []string) error {
	if len(args) != NUM_REQ_PARAMETERS {
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

	if height < 0 || width < 0 {
		return fmt.Errorf("Width and height must be positive.")
	}

	(*proj).WinSizePixel[0], (*proj).WinSizePixel[1] = width, height

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

		num, err = fmt.Sscanf(string(line), "%f %f", &(*proj).WinSizeWorld[0],
			&(*proj).WinSizeWorld[1])
	}

	if num != TWO_DIMENSIONAL && err == nil {
		err = fmt.Errorf("Tried to read %d values, read %d instead.",
			TWO_DIMENSIONAL, num)
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

	if num != THREE_DIMENSIONAL && err == nil {
		err = fmt.Errorf("Tried to read %d values, read %d instead.",
			TWO_DIMENSIONAL, num)
	}

	return err
}

func (p *Projection) String() string {
	return fmt.Sprintf("Projection:" +
			   "\n\tPixels: %d %d" +
			   "\n\tWorld size: %f %f" +
			   "\n\tViewpoint: %s",
				p.WinSizePixel[0], p.WinSizePixel[1],
				p.WinSizeWorld[0], p.WinSizeWorld[1],
				p.Viewpoint.String())
}
