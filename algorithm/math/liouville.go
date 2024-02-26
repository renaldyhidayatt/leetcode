package math

import (
	"algoritmAndDs/math/prime"
	"errors"
)

var ErrNonZeroArgsOnly error = errors.New("arguments cannot be zero")

func LiovilleLambda(n int) (int, error) {
	switch {
	case n < 0:
		return 0, ErrPosArgsOnly
	case n == 0:
		return 0, ErrNonZeroArgsOnly

	case len(prime.Factorize(int64(n)))%2 == 0:
		return 1, nil

	default:
		return -1, nil
	}
}
