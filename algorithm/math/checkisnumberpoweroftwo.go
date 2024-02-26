package math

import "math"

func IsPowOfTwoUseLog(number float64) bool {
	if number == 0 || math.Round(number) == math.MaxInt64 {
		return false
	}

	log := math.Log2(number)

	return log == math.Round(log)
}
