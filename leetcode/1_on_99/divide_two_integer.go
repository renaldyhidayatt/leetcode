package on99

import "math"

func Divide(divided int, divisor int) int {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	result := 0
	sign := -1

	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}

	dvd, dvs := abs(divided), abs(divisor)

	for dvd >= dvs {
		temp := dvs

		m := 1

		for temp<<1 <= dvd {
			temp <<= 1

			m <<= 1
		}

		dvd -= temp
		result += m
	}

	return sign * result
}
