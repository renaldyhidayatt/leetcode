package on99

import "sort"

func FourSum(nums []int, target int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)

	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSum(nums []int, left, right int, target int, k int, cur []int, res *[][]int) {
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}

	if k == 2 {
		twoSum(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))

				copy(next, cur)
				next = append(next, nums[i])

				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}
}

func twoSum(nums []int, left, right int, target int, cur []int, res *[][]int) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))

			copy(temp, cur)

			*res = append(*res, temp)

			cur = cur[:len(cur)-2]
			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}

			for left < right && nums[right] == nums[right+1] {
				right--
			}

		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}
