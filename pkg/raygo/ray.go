package raygo

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Color() Vector {
	unitDir := r.Direction.Normalize()
	t := 0.5 * (unitDir.Y + 1.0)

	return Vector{1.0, 1.0, 1.0}.Mult(1.0 - t).Add(Vector{0.5, 0.7, 1.0}.Mult(t))
}
