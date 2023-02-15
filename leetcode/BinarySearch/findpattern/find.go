package findpattern

import "math"

func find142pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	num3, stack := math.MinInt64, []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < num3 {
			return true
		}

		for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
			num3 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false
}
