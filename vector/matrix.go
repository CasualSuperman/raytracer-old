package vector

type Matrix [3]Vec3

func OrthogonalMatrix(x, y, z vectorer) (m Matrix) {
	switch {
	case x == nil:
		m[1] = y.Vector()
		m[2] = z.Vector()
		m[1].Unit()
		m[2].Unit()

		m[0] = Cross(&m[1], &m[2])
		m[0].Unit()
	case y == nil:
		m[0] = x.Vector()
		m[2] = z.Vector()
		m[0].Unit()
		m[2].Unit()

		m[1] = Cross(&m[0], &m[2])
		m[1].Unit()
	case z == nil:
		m[0] = x.Vector()
		m[1] = y.Vector()
		m[0].Unit()
		m[1].Unit()

		m[2] = Cross(&m[0], &m[1])
		m[2].Unit()
	default:
		m[0] = x.Vector()
		m[1] = y.Vector()
		m[2] = z.Vector()
		m[0].Unit()
		m[1].Unit()
		m[2].Unit()
	}
	return
}

func (m *Matrix) Xform(v *Position) {
	*v = Position{v.X*m[0].X + v.Y*m[0].Y + v.Z*m[0].Z,
		v.X*m[1].X + v.Y*m[1].Y + v.Z*m[1].Z,
		v.X*m[2].X + v.Y*m[2].Y + v.Z*m[2].Z}
}

func (m *Matrix) UnXform(v *Position) {
	*v = Position{v.X*m[0].X + v.Y*m[1].X + v.Z*m[2].X,
		v.X*m[0].Y + v.Y*m[1].Y + v.Z*m[2].Y,
		v.X*m[0].Z + v.Y*m[1].Z + v.Z*m[2].Z}
}
