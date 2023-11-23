package geometry

import (
	"errors"
	"math"
)

type EuclideanPoint []float64

var ErrDimMismatch = errors.New("mismatched dismensions")

func EuclideanDistance(p1 EuclideanPoint, p2 EuclideanPoint) (float64, error) {
	n := len(p1)

	if len(p2) != n {
		return -1, ErrDimMismatch
	}

	var total float64 = 0

	for i, x_i := range p1 {

		diff := math.Abs(x_i - p2[i])
		total += diff * diff
	}

	return math.Sqrt(total), nil
}
