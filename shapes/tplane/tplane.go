package tplane

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
	Id shapes.ShapeId = 16
)

type Tplane struct {
	plane.Plane
	xDir          vector.Direction
	width, height float64
	background    shapes.Material
	rot           vector.Matrix
}

func init() {
	shapes.RegisterShapeFormat(Id, read)
}

func read(r *bufio.Reader) (s shapes.Shape, err error) {
	line := []byte{}
	if debug.TPLANES {
		log.Println("Reading in a tiled plane.")
	}
	p := new(Tplane)
	s, err = plane.Read(r)
	p.Plane = *(s.(*plane.Plane))

	if err != nil {
		return nil, err
	}

	if debug.TPLANES {
		log.Println("Loading Tplane x direction.")
	}

	err = p.xDir.Read(r)

	if err != nil {
		return nil, err
	}

	p.xDir.Unit()

	x := p.xDir.Copy()
	z := p.Normal.Copy()

	p.rot = vector.OrthogonalMatrix(&x, &z)

	if debug.TPLANES {
		log.Println("Loading Tplane width")
	}

	if err != nil {
		return nil, err
	}

	count := 0

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		count, err = fmt.Sscanf(string(line), "%f %f", &p.width, &p.height)
	}

	if err != nil {
		return nil, err
	}

	if debug.TPLANES {
		log.Println("Reading in background material.")
	}

	err = p.background.Read(r)

	if debug.TPLANES {
		if err == nil {
			log.Println(p.String())
		} else {
			log.Println("Could not read plane.")
		}
	}

	return p, err
}

func (p *Tplane) Type() shapes.ShapeId {
	return Id
}

func (p *Tplane) hitBackground(d *vector.Position) bool {
	offset := vector.Direction(p.Center)
	offset.Invert()

	newHit := *d
	newHit.Displace(offset)

	p.rot.Xform(&newHit)

	relX := int(2<<6 + newHit.X/p.width)
	relY := int(2<<6 + newHit.Y/p.height)

	return (relX+relY)&1 == 0
}

func (p *Tplane) Ambient(d *vector.Position) color.Color {
	if p.hitBackground(d) {
		return p.background.Ambient
	}
	return p.BaseShape.Mat.Ambient
}

func (p *Tplane) Diffuse(d *vector.Position) color.Color {
	if p.hitBackground(d) {
		return p.background.Diffuse
	}
	return p.BaseShape.Mat.Diffuse
}

func (p *Tplane) Specular(d *vector.Position) color.Color {
	if p.hitBackground(d) {
		return p.background.Specular
	}
	return p.BaseShape.Mat.Specular
}

func (p *Tplane) String() string {
	return fmt.Sprintf("Tiled plane:\n\t%v\n\tcenter:\n\t%v\n\tnormal:\n\t%v"+
		"\n\tBackground material:\n%v", p.Plane.String(), p.Center.String(),
		p.Normal.String(), p.background.String())
}
