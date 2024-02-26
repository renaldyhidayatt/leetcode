package nthprime

import (
	"errors"
	"math"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}

	nth := 0

	count := 0

	for {
		count++

		if isPrime(count) {
			nth++
		}

		if nth == n {
			return count, nil
		}
	}

}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ { // corrected loop condition
		if n%i == 0 {
			return false
		}
	}

	return true
}
