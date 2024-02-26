package power

import "math"

func UsingLog(a float64, b float64) float64 {
	var p float64

	p = 1

	if a < 0 && int(b)&1 != 0 {
		p = -1
	}

	log := math.Log(math.Abs(a))
	exp := math.Exp(b * log)
	result := exp * p

	return math.Round(result)

}
