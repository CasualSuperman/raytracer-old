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

func (s *Sphere) Read(r io.Reader) error {
	err := s.shape.Read(r)
	if err != nil {
		return err
	}

	s.center.Read(r)
	num, err := fmt.Sscanf("%f", r, &s.radius)

	for num == 0 && err == nil {
		num, err = fmt.Sscanf("%f", r, &s.radius)
	}

	return err
}

func (s *Sphere) Intersect(r vector.Ray) (hit bool, length float64) {
	start := r.Position.Copy()
	start.Displace(start.Offset(s.center))

	ray := vector.NewRay(start, r.Direction.Copy())

	radius := s.radius

	a := vector.Dot(&ray.Direction, &ray.Direction)
	b := 2 * vector.Dot(&ray.Position, &ray.Direction)
	c := vector.Dot(&ray.Position, &ray.Position) - radius*radius

	discriminant := b*b - 4*a*c

	if math.Abs(discriminant) <= math.SmallestNonzeroFloat32 {
		return false, math.Inf(1)
	}

	t := (-b - math.Sqrt(discriminant)) / 2 * a

	move := r.Direction.Copy()
	hitPos := r.Position.Copy()
	move.Scale(t)
	hitPos.Displace(move)

	s.Hit = hitPos

	return true, t
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere:\n\t%v\n\t%v\n\t%v", s.shape.String(), s.center,
		s.radius)
}
