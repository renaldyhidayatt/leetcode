package on399

func CountNumbersWithUniqueDigits1(n int) int {
	res := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
	if n >= 10 {
		return res[10]
	}
	return res[n]
}
