package cylinder

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
	Id shapes.ShapeId = 22
)

type Cylinder struct {
	shapes.BaseShape
	Center     vector.Position
	CenterLine vector.Direction
	Radius     float64
	Height     float64
	rot        vector.Matrix
}

func init() {
	shapes.RegisterShapeFormat(Id, read)
}

func read(r *bufio.Reader) (shapes.Shape, error) {
	if debug.CYLINDERS {
		log.Println("Reading in a cylinder.")
	}
	s := new(Cylinder)
	err := s.BaseShape.Read(r)
	if err != nil {
		return nil, err
	}

	if debug.CYLINDERS {
		log.Println("Loading Cylinder center")
	}

	if err = s.Center.Read(r); err != nil {
		return nil, err
	}

	if debug.CYLINDERS {
		log.Println("Loading Cylinder centerline")
	}

	if err = s.CenterLine.Read(r); err != nil {
		return nil, err
	}

	s.CenterLine.Unit()

	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f %f", &s.Radius, &s.Height)
		}
	}

	if debug.CYLINDERS {
		if err == nil {
			log.Println(s.String())
		} else {
			log.Println("Could not read cylinder.")
		}
	}

	if count != 2 {
		return nil, fmt.Errorf("Unable to read height.")
	}

	if s.CenterLine.X == 0 && s.CenterLine.Y == 1 && s.CenterLine.Z == 0 {
		s.rot = vector.Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	} else {
		y := vector.Direction{0, 1, 0}
		x := vector.Cross(s.CenterLine.Vector(), y.Vector())
		x.Unit()
		if x.X < 0 {
			x.Scale(-1)
		}
		z := vector.Cross(x.Vector(), s.CenterLine.Vector())
		z.Unit()

		s.rot = vector.Matrix{x, s.CenterLine.Vector(), z}
	}

	if debug.CYLINDERS {
		log.Println("Rotation matrix:", s.rot.String())
	}

	return s, nil
}

func (s *Cylinder) Type() shapes.ShapeId {
	return Id
}

func (s *Cylinder) Hits(r vector.Ray) (hit bool, length float64, spot *vector.Ray) {
	if debug.CYLINDERS {
		log.Println("Incoming ray:", r)
		log.Println("cylinder orientation:", s.CenterLine)
		log.Println("cylinder center:", s.Center)
	}
	ray := r
	offset := vector.Direction(s.Center)
	offset.Invert()

	ray.Position.Displace(offset)

	s.rot.Xform(&ray.Position)
	s.rot.Xform((*vector.Position)(&ray.Direction))

	if debug.CYLINDERS {
		log.Println("Transformed ray:", ray.String())
	}

	d := ray.Direction
	p := ray.Position

	d.Unit()

	a := d.X*d.X + d.Z*d.Z
	b := 2 * (p.X*d.X + p.Z*d.Z)
	c := (p.X*p.X - p.Z*p.Z) - s.Radius*s.Radius

	determinant := b*b - 4*a*c

	if determinant < 0 {
		if debug.CYLINDERS {
			log.Println("Ray missed cylinder, determinant < 0.")
		}
		return false, 0, nil
	}

	length = -b - math.Sqrt(determinant)
	length /= 2 * a

	if length < 0 {
		if debug.CYLINDERS {
			log.Println("Ray missed cylinder, length < 0.")
		}
		return false, 0, nil
	}

	rotHit := p
	displacement := d
	displacement.Scale(length)
	rotHit.Displace(displacement)

	if rotHit.Y < 0 || rotHit.Y > s.Height {
		if debug.CYLINDERS {
			log.Println("Ray missed cylinder, too high or too low.")
			log.Println("Height:", rotHit.Y)
		}
		return false, 0, nil
	}

	if debug.CYLINDERS {
		log.Println("Ray hit cylinder, height:", rotHit.Y, "cylinder height:", s.Height)
	}

	ray.Direction = vector.Direction{rotHit.X, 0, rotHit.Z}
	ray.Direction.Unit()

	ray.Position = rotHit

	s.rot.UnXform(&ray.Position)
	s.rot.UnXform((*vector.Position)(&ray.Direction))

	ray.Position.Displace(vector.Direction(s.Center))

	return true, length, &ray
}

func (s *Cylinder) String() string {
	return fmt.Sprintf("Cylinder:\n\t%v\n\tcenter:\n\t%v\n\tCenterline: %s\n\tradius: %f\n\theight: %f",
		s.BaseShape.String(), s.Center.String(), s.CenterLine.String(), s.Radius, s.Height)
}
