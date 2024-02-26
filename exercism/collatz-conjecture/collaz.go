package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("Invalid input")
	}

	return collatz(n), nil
}

func collatz(n int) int {
	if n == 1 {
		return 0
	} else if n%2 == 0 {
		return 1 + collatz(n/2)
	} else {
		return 1 + collatz(3*n+1)
	}
}
