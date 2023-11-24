package on299

func CombinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}

	c, res := []int{}, [][]int{}
	FindCombinationSum3(k, n, 1, c, &res)

	return res
}

func FindCombinationSum3(k, target, index int, c []int, res *[][]int) {
	if target == 0 {
		if len(c) == k {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}

	for i := index; i < 10; i++ {
		if target >= i {
			c = append(c, i)
			FindCombinationSum3(k, target-i, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
