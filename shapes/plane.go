package shapes

import "raytracer/vector"

type sphere struct {
	center vector.Position
	radius float64
}

func (s *Sphere) Intersect(r vector.Ray) (hit bool, length float64) {
}

func String() string {
	
}
