package shapes

import "fmt"
import "math"
import "raytracer/vector"

type Sphere struct {
	Shape
	center vector.Position
	radius float64
}

func (s *Sphere) Intersect(r vector.Ray) (hit bool, length float64) {
	start := r.Position.Copy()
	start.Displace(start.Offset(s.center))

	ray := vector.NewRay(start, r.Direction.Copy())

	radius := s.radius

	a := vector.Dot(&ray.Direction, &ray.Direction)
	b := 2 * vector.Dot(&ray.Position, &ray.Direction)
	c := vector.Dot(&ray.Position, &ray.Position) - radius * radius

	discriminant := math.Sqrt(b * b - 4 * a * c)

	if math.Abs(discriminant) <= (math.SmallestNonzeroFloat64 * 10) {
		return false, math.Inf(1)
	}

	t := (-b - discriminant) / 2 * a

	hitPos := r.Position.Copy()
	move := r.Direction.Copy()
	move.Scale(t)
	hitPos.Displace(move)

	s.Shape.Hit = hitPos

	return true, t
}

func (s *Sphere) String() string {
	return fmt.Sprintf("\t%s\t%f", s.center.String(), s.radius)
}
