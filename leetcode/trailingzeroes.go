package main

import "fmt"

func TrailingZeroes(n int) int {
	count := 0

	for n > 4 {
		n /= 5
		fmt.Println(n)
		count += n
	}
	return count
}

func TrailingZeroesMain() {
	fmt.Println(TrailingZeroes(3))
}
