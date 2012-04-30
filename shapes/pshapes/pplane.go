package pshapes

import (
	"bufio"
	"fmt"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/shapes/plane"
	"raytracer/vector"
)

const (
	PlaneId shapes.ShapeId = 20
)

var (
	planeShaders = []func(*Pplane, *vector.Position, *color.Color){
		rainbows,
		lines,
	}
)

type Pplane struct {
	plane.Plane
	shaderId int
}

func init() {
	shapes.RegisterShapeFormat(PlaneId, readPlane)
}

func readPlane(r *bufio.Reader) (s shapes.Shape, err error) {
	line := []byte{}
	if debug.PPLANES {
		log.Println("Reading in a procedural plane.")
	}
	p := new(Pplane)
	s, err = plane.Read(r)
	p.Plane = *(s.(*plane.Plane))

	if err != nil {
		return nil, err
	}

	if debug.PPLANES {
		log.Println("Loading Pplane shader ID.")
	}

	count := 0

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		count, _ = fmt.Sscanf(string(line), "%d", &p.shaderId)
	}

	if err != nil {
		return nil, err
	}

	if p.shaderId >= len(planeShaders) {
		return nil, fmt.Errorf("Shader id %d invalid, only %d shaders.",
			p.shaderId, len(planeShaders))
	}

	if debug.PPLANES {
		if err == nil {
			log.Println(p.String())
		} else {
			log.Println("Could not read plane.")
		}
	}

	return p, nil
}

func (p *Pplane) Type() shapes.ShapeId {
	return PlaneId
}

func (p *Pplane) Ambient(d *vector.Position) color.Color {
	c := p.Plane.BaseShape.Mat.Ambient
	planeShaders[p.shaderId](p, d, &c)
	return c
}

func (p *Pplane) String() string {
	return fmt.Sprintf("Procedural plane:\n\t%v\n\tShader id: %d",
		p.Plane.String(), p.shaderId)
}
