package pshapes

import (
	"bufio"
	"fmt"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/shapes/sphere"
	"raytracer/vector"
)

const (
	SphereId shapes.ShapeId = 19
)

var (
	sphereShaders = []func(*Psphere, *vector.Position, *color.Color){
		sphereRainbows,
		sphereLines,
	}
)

type Psphere struct {
	sphere.Sphere
	shaderId int
}

func init() {
	shapes.RegisterShapeFormat(SphereId, readSphere)
}

func readSphere(r *bufio.Reader) (shapes.Shape, error) {
	line := []byte{}
	if debug.SPHERES {
		log.Println("Reading in a sphere.")
	}
	s := new(Psphere)

	if err := s.Sphere.Read(r); err != nil {
		return nil, err
	}

	err := fmt.Errorf("")
	err = nil

	count := 0

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		count, _ = fmt.Sscanf(string(line), "%d", &s.shaderId)
	}

	if err != nil {
		return nil, err
	}

	if s.shaderId >= len(sphereShaders) {
		return nil, fmt.Errorf("Shader id %d invalid, only %d shaders.",
			s.shaderId, len(sphereShaders))
	}

	return s, nil
}

func (s *Psphere) Type() shapes.ShapeId {
	return SphereId
}

func (p *Psphere) Ambient(d *vector.Position) color.Color {
	c := p.Sphere.BaseShape.Mat.Ambient
	sphereShaders[p.shaderId](p, d, &c)
	return c
}

func (s *Psphere) String() string {
	return fmt.Sprintf("Psphere:\n\t%v\n\tshader:\n\t%d",
		s.Sphere.String(), s.shaderId)
}
