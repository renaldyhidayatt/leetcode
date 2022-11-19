package main

import "fmt"

func trailingZeroes(n int) int {
	count := 0

	for n > 4 {
		n /= 5
		fmt.Println(n)
		count += n
	}
	return count
}

func trailingZeroesMain() {
	fmt.Println(trailingZeroes(3))
}
