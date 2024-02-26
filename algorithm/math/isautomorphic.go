package math

import "algoritmAndDs/constraints"

func IsAutomorphic[T constraints.Integer](n T) bool {
	if n < 0 {
		return false
	}

	n_sq := n * n

	for n > 0 {
		if (n % 10) != (n_sq % 10) {
			return false
		}

		n /= 10
		n_sq /= 10
	}

	return true
}
