package raygo

import "math"

type Vector struct {
	X, Y, Z float64
}

func (t Vector) Add(v Vector) Vector {
	return Vector{t.X + v.X, t.Y + v.Y, t.Z + v.Z}
}

func (t Vector) Mult(s float64) Vector {
	return Vector{t.X * s, t.Y * s, t.Z * s}
}

func (t Vector) Length() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Vector) Normalize() Vector {
	l := t.Length()
	return Vector{t.X / l, t.Y / l, t.Z / l}
}
