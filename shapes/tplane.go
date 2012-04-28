package shapes

import "bufio"
import "fmt"
import "raytracer/debug"
import "raytracer/log"
import "raytracer/vector"

type Tplane struct {
	Plane
	xDir          vector.Direction
	width, height float64
	background    Material
}

func init() {
	RegisterShapeFormat(16, readTplane)
}

func readTplane(r *bufio.Reader) (s Shape, err error) {
	if debug.TPLANES {
		log.Println("Reading in a tiled plane.")
	}
	p := new(Tplane)
	s, err = readPlane(r)
	p.Plane = *(s.(*Plane))

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

	if debug.TPLANES {
		log.Println("Loading Plane width")
	}

	line, _, err := r.ReadLine()

	if err != nil {
		return nil, err
	}

	count := 0

	for count == 0 && err == nil {
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

func (p *Tplane) Type() shapeId {
	return 16
}

func (p *Tplane) hitBackground(d *vector.Position) bool {

	x := p.xDir.Copy()
	z := p.Normal.Copy()

	rot := vector.OrthogonalMatrix(&x, nil, &z)

	offset := p.Center.Direction().Copy()
	offset.Invert()

	newHit := d.Copy()
	newHit.Displace(offset)

	rot.Xform(&newHit)

	relX := int(1024 + newHit.X/p.width)
	relY := int(1024 + newHit.Y/p.height)

	return (relX+relY)&1 == 1

	/*
		if relY == 0 {
			if relX > 0 {
				return ((relX + relY) & 1 != 1)
			}
		}
		return ((relX + relY) & 1 == 1)
	*/
}

func (p *Tplane) Ambient(d *vector.Position) vector.Vec3 {
	if p.hitBackground(d) {
		return p.background.Ambient
	}
	return p.shape.Mat.Ambient
}

func (p *Tplane) Diffuse(d *vector.Position) vector.Vec3 {
	if p.hitBackground(d) {
		return p.background.Diffuse
	}
	return p.shape.Mat.Diffuse
}

func (p *Tplane) Specular(d *vector.Position) vector.Vec3 {
	if p.hitBackground(d) {
		return p.background.Specular
	}
	return p.shape.Mat.Specular
}

func (p *Tplane) String() string {
	return fmt.Sprintf("Plane:\n\t%v\n\tcenter:\n\t%v\n\tnormal:\n\t%v",
		p.shape.String(), p.Center.String(), p.Normal.String())
}
