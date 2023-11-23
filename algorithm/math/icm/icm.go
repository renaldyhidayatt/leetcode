package icm

import (
	"algoritmAndDs/math/gcd"
	"math"
)

func Lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd.Iterative(a, b)))
}
