package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	ls := len(nums)
	for i := 0; i < ls-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, ls-1
		for j < k {
			curr := nums[i] + nums[j] + nums[k]
			if curr == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if curr < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
