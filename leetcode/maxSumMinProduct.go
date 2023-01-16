package main

import "container/list"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumMinProduct(nums []int) int {
	nums = append(nums, 0)
	sums := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	s := list.New()
	m := 0
	for i := 0; i < len(nums); i++ {
		for s.Len() > 0 && nums[s.Back().Value.(int)] > nums[i] {
			min := nums[s.Back().Value.(int)]
			s.Remove(s.Back())
			start := 0
			if s.Len() > 0 {
				start = s.Back().Value.(int) + 1
			}
			end := i
			m = max(m, min*(sums[end]-sums[start]))
		}
		s.PushBack(i)
	}

	return m % 1000000007
}
