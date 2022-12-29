package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	ls := len(nums)
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < ls-2; i++ {
		j, k := i+1, ls-1
		for j < k {
			temp := nums[i] + nums[j] + nums[k]
			if abs(target-temp) < abs(target-res) {
				res = temp
			}
			if temp < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
