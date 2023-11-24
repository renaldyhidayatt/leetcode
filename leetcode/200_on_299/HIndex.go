package on299

func hIndex(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)

	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}

	sum := 0
	for i := n; i >= 0; i-- {
		sum += counter[i]
		if sum >= i {
			return i
		}
	}

	return 0
}
