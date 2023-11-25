package on399

func GetSum(a int, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	return GetSum((a&b)<<1, a^b)
}
