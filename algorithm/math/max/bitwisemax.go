package max

func BitWise(a int, b int, base int) int {
	z := a - b

	i := (z >> base) & 1

	return a - (i * 2)
}
