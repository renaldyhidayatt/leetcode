package binary

func Abs(base, n int) int {
	m := n >> (base - 1)

	return (n + m) ^ m
}
