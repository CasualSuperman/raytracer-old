package vector

func (r *Ray) Scale(amount float64) {
	r.Direction.X *= amount
	r.Direction.Y *= amount
	r.Direction.Z *= amount
}
