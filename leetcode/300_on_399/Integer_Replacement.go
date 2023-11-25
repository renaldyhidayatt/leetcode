package on399

func IntegerReplaceMent(n int) int {
	res := 0

	for n > 1 {
		if (n & 1) == 0 {
			n >>= 1
		} else if (n+1)%4 == 0 && n != 3 {
			n++
		} else {
			n--
		}

		res++
	}

	return res
}
