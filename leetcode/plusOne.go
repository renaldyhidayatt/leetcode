package main

func plusOne(digits []int) []int {
	ls := len(digits)
	for index := ls - 1; index >= 0; index-- {
		if digits[index] < 9 {
			digits[index]++
			return digits
		} else {
			digits[index] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
