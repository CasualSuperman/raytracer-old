package shapes

import (
	"bufio"
	"fmt"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	vec "raytracer/vector"
)

const (
	Id ShapeId = 10
)

// Internal light ID counter.
var lightId = 0

// All lights have these things.
type BaseLight struct {
	id     int
	center vec.Position
	color  color.Color
}

func init() {
	RegisterLightFormat(Id, readLight)
}

func (l BaseLight) Id() int {
	return l.id
}

func (l *BaseLight) Illuminated(v *vec.Position) bool {
	return true
}

func (l *BaseLight) Color() color.Color {
	return l.color
}

func (l *BaseLight) Position() vec.Position {
	return l.center
}

func (l *BaseLight) Read(r *bufio.Reader) (err error) {
	light, err := readLight(r)
	*l = *(light.(*BaseLight))
	return err
}

func readLight(r *bufio.Reader) (l Light, err error) {
	var light BaseLight
	light.id = lightId
	lightId++

	if debug.LIGHTS {
		log.Println("Reading in light intensity.")
	}

	err = light.color.Read(r)

	if err != nil {
		return
	}

	if debug.LIGHTS {
		log.Println("Reading in light position.")
	}

	err = light.center.Read(r)

	return &light, err
}

func (l BaseLight) String() string {
	return fmt.Sprintf("Light:\n\tId: %d\n\tPosition:\n\t%s\n\tIntensity:\n\t%v\n", l.id, l.center.String(), l.color.String())
}
