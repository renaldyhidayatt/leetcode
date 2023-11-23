package on199

func TrailingZeroes(n int) int {
	if n/5 == 0 {
		return 0
	}

	return n/5 + TrailingZeroes(n/5)
}
