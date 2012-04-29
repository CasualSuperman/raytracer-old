package shapes

import (
	"bufio"
	"fmt"
	"math"
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

func (l *Spotlight) Illuminated(p *vector.Position) bool {
	start := l.BaseLight.Position()
	centerToHit := start.Offset(*p)
	centerToHit.Unit()


	if debug.SPOTLIGHTS {
		log.Println("Spotlight to point:", centerToHit)
		log.Println("cos(theta)", l.theta)
		log.Println("Dot Product:", vector.Dot(&centerToHit, &l.aim))
	}

	return l.theta < vector.Dot(&centerToHit, &l.aim)
}

func readLight(r *bufio.Reader) (l shapes.Light, err error) {
	var light Spotlight
	err = light.BaseLight.Read(r)

	if err != nil {
		return l, err
	}

	err = light.aim.Read(r)

	pos := light.BaseLight.Position()
	tempPos := light.aim.Position()

	light.aim = pos.Offset(*tempPos)
	light.aim.Unit()

	if err != nil {
		return l, err
	}

	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f", &light.theta)
		}
	}

	if debug.LIGHTS {
		if err == nil {
			log.Println(light.String())
		} else {
			log.Println("Could not read spotlight.")
		}
	}

	light.theta = math.Cos(light.theta * math.Pi / 180)

	return &light, err
}

func (l Spotlight) String() string {
	return fmt.Sprintf("Spotlight:\n\tId: %s\n\tDirection:\n\t%s\n\t" +
		"Theta:\n\t%v\n", l.BaseLight.String(), l.aim.String(), l.theta)
}
