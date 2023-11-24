package on299

import "sort"

func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	diff := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]-i-1 >= diff {
			diff = nums[i] - i - 1
		} else {
			return nums[i]
		}
	}
	return 0
}
