package main

func arraySign(nums []int) int {
	result := 1
	positive := 0
	for _, element := range nums {
		if element == 0 {
			return 0
		}
		if element < 0 {
			positive++
		}
		result *= element
	}
	if positive%2 != 0 {
		result = -result
	}
	return result
}
