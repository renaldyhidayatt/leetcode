package on399

import "sort"

func WiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	array := make([]int, len(nums))

	copy(array, nums)

	sort.Ints(array)

	n := len(nums)
	left := (n+1)/2 - 1
	right := n - 1

	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			nums[i] = array[right]
			right--
		} else {
			nums[i] = array[left]
			left--
		}
	}
}
