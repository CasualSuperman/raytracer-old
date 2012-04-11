package view

import "fmt"
import "os"
import "strconv"

const (
	TWO_DIMENSIONAL = 2
	THREE_DIMENSIONAL = 3
	NUM_REQ_PARAMETERS = 3
)

type Projection struct {
	Win_size_pixel [TWO_DIMENSIONAL]int
	Win_size_world [TWO_DIMENSIONAL]float64
	View_point     [THREE_DIMENSIONAL]float64
}

func NewProjection() (Projection, error) {
	var proj Projection

	if len(os.Args) != NUM_REQ_PARAMETERS {
		return proj, fmt.Errorf("usage:\n\t%s width height", os.Args[0])
	}

	height, errH := strconv.Atoi(os.Args[2])
	width, errW := strconv.Atoi(os.Args[1])

	if errH != nil || errW != nil {
		return proj, fmt.Errorf("Unable to parse height or width.")
	}

	proj.Win_size_pixel[0], proj.Win_size_pixel[1] = width, height

	return proj, nil
}