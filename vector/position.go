package vector

import "fmt"

func NewPosition(x, y, z float64) Position {
	return Position{x, y, z}
}

func Origin() Position {
	return Position{0, 0, 0}
}

func (p *Position) Displace(d *Direction) {
	p.X += d.X
	p.Y += d.Y
	p.Z += d.Z
}

/* Returns a Direction from the source to the target */
func (source *Position) Offset(target *Position) Direction {
	return Direction{
		target.X - source.X,
		target.Y - source.Y,
		target.Z - source.Z,
	}
}

func (p *Position) String() string {
	return fmt.Sprintln("<%f, %f, %f>", p.X, p.Y, p.Z)
}
