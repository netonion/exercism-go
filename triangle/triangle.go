package triangle

import (
	"math"
)

// Kind represents the kind of the triangle.
type Kind int

// Kinds of triangle
const (
	NaT Kind = iota // not a triangle.
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the the kind of a triangle based on its sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isInvalid(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}

func isInvalid(a, b, c float64) bool {
	sum := a + b + c
	for _, s := range []float64{a, b, c} {
		if math.IsNaN(s) || s <= 0 || math.IsInf(s, 0) || s > sum/2 {
			return true
		}
	}
	return false
}
