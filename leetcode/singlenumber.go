package main

import "fmt"

func Solution(num []int) int {
	value := 0
	for _, v := range num {
		value ^= v
	}

	return value
}

func SingleNumber() {
	fmt.Println(Solution([]int{4, 1, 2, 1, 2}))
}
