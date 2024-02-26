package dynamic

import "algoritmAndDs/math/max"

func LongestIncreasingSubsequence(elements []int) int {
	n := len(elements)
	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}
	for i := range lis {
		for j := 0; j < i; j++ {
			if elements[i] > elements[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	res := 0
	for _, value := range lis {
		res = max.Int(res, value)
	}
	return res
}
