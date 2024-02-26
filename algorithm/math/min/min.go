package min

import "algoritmAndDs/constraints"

func Int[T constraints.Integer](values ...T) T {
	min := values[0]

	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
