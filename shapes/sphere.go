package shapes

import "fmt"
import "io"
import "math"
import "raytracer/vector"

type Sphere struct {
	shape
	center vector.Position
	radius float64
}

func init() {
	RegisterFormat(13, read)
}

func read(r io.Reader) (Shape, error) {
	s := new(Sphere)
	err := s.shape.Read(r)
	if err != nil {
		return nil, err
	}

	s.center.Read(r)
	num, err := fmt.Fscanf(r, "%f", &s.radius)

	for num == 0 && err == nil {
		num, err = fmt.Fscanf(r, "%f", &s.radius)
	}

	return s, err
}

func (s *Sphere) Hits(r vector.Ray) (hit bool, length float64, spot vector.Ray) {
	start := r.Position.Copy()
	start.Displace(start.Offset(s.center))

	radius := s.radius

	a := vector.Dot(&r.Direction, &r.Direction)
	b := 2 * vector.Dot(&r.Position, &r.Direction)
	c := vector.Dot(&r.Position, &r.Position) - radius*radius

	discriminant := b*b - 4*a*c

	if math.Abs(discriminant) <= math.SmallestNonzeroFloat32 {
		return false, math.Inf(1), spot
	}

	length = (-b - math.Sqrt(discriminant)) / 2 * a

	move := r.Direction.Copy()
	hitPos := r.Position.Copy()
	move.Scale(length)
	hitPos.Displace(move)

	// Set the hit position to where we hit it.
	spot.Position = hitPos

	// Start the direction at the hit position, and subtract the center of the
	// sphere.
	spot.Direction = *hitPos.Direction()
	spot.Direction.Sub(s.center.Direction())

	return
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere:\n\t%v\n\t%v\n\t%v", s.shape.String(), s.center,
		s.radius)
}
