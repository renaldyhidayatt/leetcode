package main

import "math"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	for _, n := range nums {
		numSet[n] = true
	}

	maxLen := 0

	for len(numSet) > 0 {
		n := getAnyKey(numSet)
		numSet[n] = false

		l1, l2 := 0, 0
		i := n + 1
		for numSet[i] {
			numSet[i] = false
			i++
			l1++
		}
		i = n - 1
		for numSet[i] {
			numSet[i] = false
			i--
			l2++
		}
		maxLen = maax(maxLen, l1+l2+1)
	}

	return maxLen
}

func getAnyKey(m map[int]bool) int {
	for k := range m {
		return k
	}

	return 0
}

func maax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
