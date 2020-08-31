// Package triangle should have a package comment that summarizes what it's about.
package triangle

import "math"

type Kind string

const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if !isValidTriangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}

func isValidTriangle(a, b, c float64) bool {
	switch {
	case a <= 0 || math.IsNaN(a) || math.IsInf(a, 0):
		return false
	case b <= 0 || math.IsNaN(b) || math.IsInf(b, 0):
		return false
	case c <= 0 || math.IsNaN(c) || math.IsInf(c, 0):
		return false
	case a+b < c:
		return false
	case b+c < a:
		return false
	case c+a < b:
		return false
	default:
		return true
	}
}
