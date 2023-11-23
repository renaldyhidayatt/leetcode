package gcd

func ExtendedRecursive(a, b int64) (int64, int64, int64) {
	if b > 0 {
		d, y, x := ExtendedRecursive(b, a%b)
		y -= (a / b) * x
		return d, x, y
	}

	return a, 1, 0
}
