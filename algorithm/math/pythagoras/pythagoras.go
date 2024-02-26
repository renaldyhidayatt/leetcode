package pythagoras

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func Distance(a, b Vector) float64 {
	res := math.Pow(b.x-a.x, 2.0) + math.Pow(b.y-a.y, 2.0) + math.Pow(b.z-a.z, 2.0)
	return math.Sqrt(res)
}
