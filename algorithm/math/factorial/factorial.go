package factorial

func Iterative(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func Recursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Recursive(n-1)
	}
}

func UsingTree(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	return prodTree(2, n)
}

func prodTree(l int, r int) int {
	if l > r {
		return 1
	}

	if l == r {
		return l
	}

	if r-l == 1 {
		return l * r
	}

	m := (l + r) / 2

	return prodTree(l, m) * prodTree(m+1, r)
}
