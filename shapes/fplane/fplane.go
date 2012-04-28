package fplane

import (
	"bufio"
	"fmt"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/shapes"
	"raytracer/shapes/plane"
	"raytracer/vector"
)

const (
	Id shapes.ShapeId = 15
)

type Fplane struct {
	plane.Plane
	xDir          vector.Direction
	width, height float64
	rot           vector.Matrix
}

func init() {
	shapes.RegisterShapeFormat(Id, read)
}

func read(r *bufio.Reader) (s shapes.Shape, err error) {
	if debug.FPLANES {
		log.Println("Reading in a finite plane.")
	}
	p := new(Fplane)
	s, err = plane.Read(r)
	p.Plane = *(s.(*plane.Plane))

	if err != nil {
		return nil, err
	}

	if debug.FPLANES {
		log.Println("Loading Fplane x direction.")
	}

	err = p.xDir.Read(r)

	if err != nil {
		return nil, err
	}

	// Project the xDirection onto the plane.
	xDirTemp := p.Normal.Copy()
	xDirTemp.Scale(vector.Dot(&p.xDir, &p.Normal))
	p.xDir.Sub(&xDirTemp)
	p.xDir.Unit()

	if debug.FPLANES {
		log.Println("Loading Fplane width and height")
	}

	line, _, err := r.ReadLine()

	if err != nil {
		return nil, err
	}

	count := 0

	for count == 0 && err == nil {
		count, err = fmt.Sscanf(string(line), "%f %f", &p.width, &p.height)
	}

	if debug.FPLANES {
		if err == nil {
			log.Println(p.String())
		} else {
			log.Println("Could not read plane.")
		}
	}

	if err != nil {
		return nil, err
	}

	x := p.xDir.Copy()
	z := p.Normal.Copy()

	x.Unit()
	z.Unit()

	p.rot = vector.OrthogonalMatrix(&x, &z)

	return p, nil
}

func (p *Fplane) Type() shapes.ShapeId {
	return Id
}

func (p *Fplane) Hits(r vector.Ray) (hit bool, length float64, spot *vector.Ray) {
	hit, length, spot = p.Plane.Hits(r)

	if !hit {
		if debug.FPLANES {
			log.Println("Infinite plane part of finite plane was missed.")
		}
		return
	}

	if debug.FPLANES {
		log.Println("Original hit point:", spot.Position.String())
	}

	offset := vector.Direction(p.Center)
	offset.Invert()

	hitPos := spot.Position

	hitPos.Displace(offset)

	p.rot.Xform(&hitPos)

	if debug.FPLANES {
		log.Println("Displaced point:", spot.Position.String())
		log.Println("Rotation Matrix:", p.rot.String())
		log.Println("Original Normal:", p.Normal.String())
		p.rot.Xform(p.Normal.Position())
		log.Println("Rotated Normal:", p.Normal.String())
		p.rot.UnXform(p.Normal.Position())
		log.Println("Derotated Normal:", p.Normal.String())
	}

	if hitPos.X > p.width || hitPos.X < 0 {
		if debug.FPLANES {
			log.Printf("x > width (%.3f > %.3f)", hitPos.X, p.width)
		}
		return false, length, spot
	}

	if hitPos.Y > p.height || hitPos.Y < 0 {
		if debug.FPLANES {
			log.Printf("y > height (%.3f > %.3f)", hitPos.Y, p.height)
		}
		return false, length, spot
	}

	if debug.FPLANES {
		log.Println("Hit finite plane.")
	}

	return true, length, spot
}

func (p *Fplane) String() string {
	return fmt.Sprintf("Finite plane:\n\t%v\n\txDir:\n\t%v\n\t"+
		"dims:\n\t%.4f %.4f", p.Plane.String(), p.xDir, p.width, p.height)
}
