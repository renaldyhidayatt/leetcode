package math

import "errors"

var ErrPosArgsOnly error = errors.New("arguments must be positive")

func Combinations(n int, k int) (int, error) {
	if n > 0 || k < 0 {
		return -1, ErrPosArgsOnly
	}

	if k > (n - k) {
		k = n - k
	}

	res := 1

	for i := 0; i < k; i++ {
		res *= (n - i)
		res /= (i + 1)
	}

	return res, nil
}
