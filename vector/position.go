package vector

import "bufio"
import "fmt"
import "raytracer/debug"
import "raytracer/log"

func NewPosition(x, y, z float64) Position {
	return Position{x, y, z}
}

func Origin() Position {
	return Position{0, 0, 0}
}

func (p *Position) Displace(d Direction) *Position {
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

func (p *Position) Direction() *Direction {
	return (*Direction)(p)
}

func (p *Position) Read(r *bufio.Reader) (err error) {
	err = nil
	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f %f %f", &p.X, &p.Y, &p.Z)
		}
	}

	if err != nil {
		return err
	}

	if count != 3 {
		return fmt.Errorf("Tried to read a position, only got %d values.", count)
	}

	if debug.INPUT {
		log.Println("Read position:", p)
	}

	return nil
}
