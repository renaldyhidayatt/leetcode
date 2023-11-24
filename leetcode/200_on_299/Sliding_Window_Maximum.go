package on299

func MaxSlidingWindow(a []int, k int) []int {
	res := make([]int, 0, k)

	n := len(a)

	if n == 0 {
		return []int{}
	}

	for i := 0; i <= n-k; i++ {
		max := a[i]
		for j := 1; j < k; j++ {
			if max < a[i+j] {
				max = a[i+j]
			}
		}
		res = append(res, max)
	}

	return res

}
