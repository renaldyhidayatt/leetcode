package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	begin := 0
	for index := range nums {
		if index == 0 || nums[index] != nums[index-1] {
			// generate all
			begin = 0
		}
		size := len(res)
		// use existing subsets to generate new subsets
		for j := begin; j < size; j++ {
			curr := make([]int, len(res[j]))
			copy(curr, res[j])
			curr = append(curr, nums[index])
			res = append(res, curr)
		}
		// avoid duplicate subsets
		begin = size
	}
	return res
}
