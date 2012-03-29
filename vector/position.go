package vector

func (p *Position) Displace(d *Direction) {
	p.X += d.X
	p.Y += d.Y
	p.Z += d.Z
}

/* Returns a Direction from the source to the target */
func (source *Position) Offset(target *Position) Direction {
	return Direction {
		target.X - source.X,
		target.Y - source.Y,
		target.Z - source.Z,
	}
}
