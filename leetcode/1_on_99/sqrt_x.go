package on99

func MySqrt(x int) int {
	r := x

	for r*r > x {
		r = (x + x/r) / 2
	}

	return r
}
