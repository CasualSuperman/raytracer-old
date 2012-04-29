package shapes

import (
	"bufio"
	"fmt"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/vector"
)

const (
	Id shapes.ShapeId = 11
)

// All lights have these things.
type Spotlight struct {
	shapes.BaseLight
	aim   vector.Direction
	theta float64
}

func init() {
	shapes.RegisterLightFormat(Id, readLight)
}

func (l Spotlight) Id() int {
	return l.BaseLight.Id()
}

func readLight(r *bufio.Reader) (l shapes.Light, err error) {
	var light Spotlight
	err = light.BaseLight.Read(r)

	if err != nil {
		return l, err
	}

	err = light.aim.Read(r)

	if err != nil {
		return l, err
	}

	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f", &light.aim)
		}
	}

	if debug.LIGHTS {
		if err == nil {
			log.Println(light.String())
		} else {
			log.Println("Could not read sphere.")
		}
	}

	return &light, err
}

func (l Spotlight) String() string {
	return fmt.Sprintf("Spotight:\n\tId: %d\n\tDirection:\n\t%s\n\tTheta:\n\t%v\n", l.BaseLight.String(), l.aim.String(), l.theta)
}
