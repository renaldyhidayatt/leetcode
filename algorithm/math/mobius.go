package math

import "algoritmAndDs/math/prime"

func Mu(n int) int {
	if n <= 1 {
		return 1
	}

	var primeFactorCount int

	for i := 1; i <= n; i++ {
		if n%i == 0 && prime.OptimizedTrialDivision(int64(i)) {
			if n%(i*i) == 0 {
				return 0
			}

			primeFactorCount += 1
		}
	}

	if primeFactorCount%2 == 0 {
		return 1
	}

	return -1
}
