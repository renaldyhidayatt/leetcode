package triangle

import "math"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
