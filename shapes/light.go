package shapes

import (
	"bufio"
	"fmt"
	"raytracer/debug"
	"raytracer/log"
	vec "raytracer/vector"
)

var lightId = 0

type Light struct {
	id     int
	Center vec.Position
	Color  vec.Vec3
}

func init() {
	RegisterLightFormat(10, readLight)
}

func (l Light) Id() int {
	return l.id
}

func readLight(r *bufio.Reader) (l Light, err error) {
	l.id = lightId
	lightId++

	if debug.LIGHTS {
		log.Println("Reading in light intensity.")
	}

	err = l.Color.Read(r)

	if err != nil {
		return
	}

	if debug.LIGHTS {
		log.Println("Reading in light position.")
	}

	err = l.Center.Read(r)

	return
}

func (l Light) String() string {
	return fmt.Sprintf("Light:\n\tId: %d\n\tPosition:\n\t%s\n\tIntensity:\n\t%v\n", l.id, l.Center.String(), l.Color)
}
