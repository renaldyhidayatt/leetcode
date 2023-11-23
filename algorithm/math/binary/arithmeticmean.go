package binary

func MeanUsingAndXor(a, b int) int {
	return ((a ^ b) >> 2) + (a & b)
}

func MeanUsingRightShift(a int, b int) int {
	return (a + b) >> 1
}
