package on99

func LargestRectangleArea(heights []int) int {
	maxArea := 0

	n := len(heights) + 2

	getHeigh := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}

		return heights[i-1]
	}

	st := make([]int, 0, n/2)

	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeigh(st[len(st)-1]) > getHeigh(i) {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeigh(idx)*(i-st[len(st)-1]-1))
		}
		st = append(st, i)
	}

	return maxArea
}
