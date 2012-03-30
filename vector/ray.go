package vector

func NewRay(p Position, d Direction) Ray {
	return Ray{p, d}
}

func (r *Ray) String() string {
	return "P: " + r.Position.String() + " D: " + r.Direction.String()
}
