package main

import "sort"

func subsets(nums []int) [][]int {
	// Sort and iteratively generate n subset with n-1 subset, O(n^2) and O(2^n)
	sort.Ints(nums)
	res := [][]int{{}}
	for index := range nums {
		size := len(res)
		// use existing subsets to generate new subsets
		for j := 0; j < size; j++ {
			curr := make([]int, len(res[j]))
			copy(curr, res[j])
			curr = append(curr, nums[index])
			res = append(res, curr)
		}
	}
	return res
}
