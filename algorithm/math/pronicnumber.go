package math

import stdMath "math"

func PronicNumber(n int) bool {
	if n < 0 || n%2 == 1 {
		return false
	}

	x := int(stdMath.Sqrt(float64(n)))
	return n == x*(x+1)
}
