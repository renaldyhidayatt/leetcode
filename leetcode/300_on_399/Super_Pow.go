package on399

func SuperPow(a int, b []int) int {
	res := 1

	for i := 0; i < len(b); i++ {
		res = (qPow(res, 10) * qPow(a, b[i])) % 1337
	}

	return res
}

func qPow(x, n int) int {
	res := 1

	x %= 1337

	for n > 0 {
		if (n & 1) == 1 {
			res = (res * x) % 1337
		}

		x = (x * x) % 1337
		n >>= 1
	}

	return res
}
