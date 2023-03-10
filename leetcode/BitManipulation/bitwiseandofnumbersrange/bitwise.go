package bitwiseandofnumbersrange

func rangeBitwiseAnd1(m int, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m != n {
		m >>= 1
		n >>= 1
		moved++
	}
	return m << uint32(moved)
}
