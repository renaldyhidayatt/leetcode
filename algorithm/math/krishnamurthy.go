package math

import "algoritmAndDs/constraints"

func IsKrishnamurthy[T constraints.Integer](n T) bool {
	if n <= 0 {
		return false
	}

	digitFact := make([]T, 10)
	digitFact[0] = 1

	for i := 1; i < 10; i++ {
		digitFact[i] = digitFact[i-1] * T(i)
	}

	nTemp := n

	for n > 0 {
		nTemp -= digitFact[n%10]

		n /= 10
	}

	return nTemp == 0
}
