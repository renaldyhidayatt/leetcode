package catalan

import f "algoritmAndDs/math/factorial"

func CatalanNumber(n int) int {
	return f.Iterative(n*2) / (f.Iterative(n) * f.Iterative(n+1))
}
