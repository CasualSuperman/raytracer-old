package vector

import "fmt"
import "raytracer/log"

type Matrix [3]Vec3

func OrthogonalMatrix(x, z vectorer) (m Matrix) {
	m[0] = x.Vector()
	m[2] = z.Vector()
	m[0].Unit()
	m[2].Unit()

	m[1] = Cross(m[2], m[0])
	m[1].Unit()
	log.Println("Finished Matrix:", m)
	return
}

func (m *Matrix) String() string {
	return fmt.Sprintf("\n\t%s\n\t%s\n\t%s", m[0].String(), m[1].String(),
		m[2].String())
}

func (m *Matrix) Xform(v *Position) {
	*v = Position{
		v.X*m[0].X + v.Y*m[0].Y + v.Z*m[0].Z,
		v.X*m[1].X + v.Y*m[1].Y + v.Z*m[1].Z,
		v.X*m[2].X + v.Y*m[2].Y + v.Z*m[2].Z}
}

func (m *Matrix) UnXform(v *Position) {
	*v = Position{
		v.X*m[0].X + v.Y*m[1].X + v.Z*m[2].X,
		v.X*m[0].Y + v.Y*m[1].Y + v.Z*m[2].Y,
		v.X*m[0].Z + v.Y*m[1].Z + v.Z*m[2].Z}
}
