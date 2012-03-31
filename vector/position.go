package vector

import "fmt"

func NewPosition(x, y, z float64) Position {
	return Position{x, y, z}
}

func Origin() Position {
	return Position{0, 0, 0}
}

func (p *Position) Displace(d Direction) *Position{
	p.X += d.X
	p.Y += d.Y
	p.Z += d.Z
	return p
}

/* Returns a Direction from the source to the target */
func (source *Position) Offset(target Position) Direction {
	return Direction{
		target.X - source.X,
		target.Y - source.Y,
		target.Z - source.Z,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("<%f, %f, %f>", p.X, p.Y, p.Z)
}

func (p Position) Copy() Position {
	return Position{p.X, p.Y, p.Z}
}

func (p *Position) Vector() Vec3 {
	return Vec3(*p)
}
