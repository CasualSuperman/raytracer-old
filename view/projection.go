package view

import "bufio"
import "fmt"
import "io"
import "os"
import "strconv"

const (
	TWO_DIMENSIONAL    = 2
	THREE_DIMENSIONAL  = 3
	NUM_REQ_PARAMETERS = 3
)

type Projection struct {
	Win_size_pixel [TWO_DIMENSIONAL]int
	Win_size_world [TWO_DIMENSIONAL]float64
	View_point     [THREE_DIMENSIONAL]float64
}

func loadProjectionPixels(proj *Projection) error {
	if len(os.Args) != NUM_REQ_PARAMETERS {
		return fmt.Errorf("usage:\n\t%s width height", os.Args[0])
	}

	width, errW := strconv.Atoi(os.Args[1])
	height, errH := strconv.Atoi(os.Args[2])

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

	(*proj).Win_size_pixel[0], (*proj).Win_size_pixel[1] = width, height

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

		num, err = fmt.Sscanf(string(line), "%f %f", &(*proj).Win_size_world[0],
			&(*proj).Win_size_world[1])
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

		num, err = fmt.Sscanf(string(line), "%f %f %f", &(*proj).View_point[0],
			&(*proj).View_point[1], &(*proj).View_point[2])
	}

	if num != THREE_DIMENSIONAL && err == nil {
		err = fmt.Errorf("Tried to read %d values, read %d instead.",
			TWO_DIMENSIONAL, num)
	}

	return err
}

func NewProjection(input io.Reader) (p Projection, err error) {
	err = loadProjectionPixels(&p)
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

	fmt.Fprintln(os.Stderr, p)

	return
}
