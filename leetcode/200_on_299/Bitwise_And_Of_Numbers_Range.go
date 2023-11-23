package on299

func RangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1)
	}

	return n
}
