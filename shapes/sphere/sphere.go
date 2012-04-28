package sphere

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
	Id shapes.ShapeId = 13
)

type Sphere struct {
	shapes.BaseShape
	Center vector.Position
	Radius float64
}

func init() {
	shapes.RegisterShapeFormat(Id, read)
}

func read(r *bufio.Reader) (shapes.Shape, error) {
	if debug.SPHERES {
		log.Println("Reading in a sphere.")
	}
	s := new(Sphere)
	err := s.BaseShape.Read(r)
	if err != nil {
		return nil, err
	}

	if debug.SPHERES {
		log.Println("Loading Sphere center")
	}

	s.Center.Read(r)

	if debug.SPHERES {
		log.Println("Loading Sphere radius")
	}

	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f", &s.Radius)
		}
	}

	if debug.SPHERES {
		if err == nil {
			log.Println(s.String())
		} else {
			log.Println("Could not read sphere.")
		}
	}

	return s, nil
}

func (s *Sphere) Type() shapes.ShapeId {
	return Id
}

func (s *Sphere) Hits(r vector.Ray) (hit bool, length float64, spot *vector.Ray) {
	spot = &vector.Ray{s.Center, r.Direction}

	radius := s.Radius

	newPos := r.Position.Copy()

	offset := s.Center.Copy()
	newPos.Displace(offset.Offset(vector.Origin()))

	if debug.SPHERES {
		log.Println("Viewpoint adjusted to origin:", newPos)
	}

	a := vector.Dot(&r.Direction, &r.Direction)
	b := 2 * vector.Dot(&newPos, &r.Direction)
	c := vector.Dot(&newPos, &newPos) - radius*radius

	if debug.SPHERES {
		log.Printf("a: %f, b: %f, c: %f\n", a, b, c)
	}

	discriminant := b*b - 4*a*c

	if discriminant <= 0 {
		return false, length, spot
	}

	length = (-b - math.Sqrt(discriminant)) / (2 * a)

	if debug.SPHERES {
		log.Printf("Sphere discriminant is %f, (T = %f)\n", discriminant, length)
	}

	if length < 0 && !vector.IsZero(length) {
		if debug.SPHERES {
			log.Println("Negative length.")
		}
		return false, length, spot
	}

	move := r.Direction.Copy()
	hitPos := r.Position.Copy()
	move.Scale(length)
	hitPos.Displace(move)

	// Set the hit position to where we hit it.
	spot.Position = hitPos

	// Start the direction at the hit position, and subtract the center of the
	// sphere.
	spot.Direction = s.Center.Offset(hitPos)
	spot.Direction.Unit()

	if debug.SPHERES {
		log.Println("Hit Location:", hitPos)
		log.Println("Sphere Center:", s.Center)
		log.Println("Hit Normal:", spot.Direction)
	}

	hit = true

	return
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere:\n\t%v\n\tcenter:\n\t%v\n\tradius:\n\t%v",
		s.BaseShape.String(), s.Center.String(), s.Radius)
}
