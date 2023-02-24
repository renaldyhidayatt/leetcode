package randomweight

import "math/rand"

type Solution struct {
	prefixSum []int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))

	for i, e := range w {
		if i == 0 {
			prefixSum[i] = e
			continue
		}
		prefixSum[i] = prefixSum[i-1] + e
	}

	return Solution{prefixSum: prefixSum}
}

func (so *Solution) PickIndex() int {
	n := rand.Intn(so.prefixSum[len(so.prefixSum)-1]) + 1
	low, high := 0, len(so.prefixSum)-1
	for low < high {
		mid := low + (high-low)>>1
		if so.prefixSum[mid] == n {
			return mid
		} else if so.prefixSum[mid] < n {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
