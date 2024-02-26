package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	var number int
	for i := len(inputDigits) - 1; i >= 0; i-- {
		d := inputDigits[i]
		if d < 0 || d >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		p := len(inputDigits) - 1 - i
		number += d * int(math.Pow(float64(inputBase), float64(p)))
	}
	if number == 0 {
		return []int{0}, nil
	}
	var digits []int
	for i := number; i > 0; i /= outputBase {
		digits = append([]int{i % outputBase}, digits...)
	}
	return digits, nil
}
