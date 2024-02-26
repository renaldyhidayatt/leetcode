package math

import (
	"algoritmAndDs/constraints"
	"errors"
)

var ErrEmptySlice = errors.New("empty slice provided")

func Mode[T constraints.Number](numbers []T) (T, error) {
	countMap := make(map[T]int)

	n := len(numbers)

	if n == 0 {
		return 0, ErrEmptySlice
	}

	for _, number := range numbers {
		countMap[number]++
	}

	var mode T

	count := 0

	for k, v := range countMap {
		if v > count {
			count = v
			mode = k
		}
	}

	return mode, nil
}
