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

type Fplane struct {
	plane.Plane
	xDir          vector.Direction
	width, height float64
}

func init() {
	shapes.RegisterShapeFormat(15, read)
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

	return p, nil
}

func (p *Fplane) Type() shapes.ShapeId {
	return 15
}

func (p *Fplane) Hits(r vector.Ray) (hit bool, length float64, spot *vector.Ray) {
	hit, length, spot = p.Plane.Hits(r)

	if !hit {
		if debug.FPLANES {
			log.Println("Infinite plane part of finite plane was missed.")
		}
		return
	}

	offset := p.Center.Direction().Copy()
	offset.Invert()

	newHit := spot.Position.Copy()
	newHit.Displace(offset)

	x := p.xDir.Copy()
	z := p.Normal.Copy()

	rot := vector.OrthogonalMatrix(&x, nil, &z)

	rot.Xform(&newHit)

	if debug.FPLANES {
		log.Println("Rotation Matrix:", rot)
		log.Println("Displaced point:", newHit)
		log.Println("Original Normal:", p.Normal)
		rot.Xform(p.Normal.Position())
		log.Println("Rotated Normal:", p.Normal)
		rot.UnXform(p.Normal.Position())
		log.Println("Derotated Normal:", p.Normal)
	}

	if newHit.X > p.width || newHit.X < 0 {
		return false, length, spot
	}

	if newHit.Y < -p.height || newHit.Y > 0 {
		return false, length, spot
	}

	return true, length, spot
}

func (p *Fplane) String() string {
	return fmt.Sprintf("Finite plane:\n\t%v\n\t&v\n\txDir:\n\t%v\n\t" +
		"dims:\n\t%.4f %.4f", p.BaseShape.String(), p.Plane.String(), p.xDir,
		p.width, p.height)
}
