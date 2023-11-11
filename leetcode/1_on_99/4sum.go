package on99

import (
	"fmt"
	"sort"
)

func kSum[T int | int32 | int64](nums []T, target T) [][]T {
	res, cur := make([][]T, 0), make([]T, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	kSumHelper(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSumHelper[T int | int32 | int64](nums []T, left, right int, target T, k int, cur []T, res *[][]T) {
	if right-left+1 < k || k < 2 || target < T(k)*nums[left] || target > T(k)*nums[right] {
		return
	}

	if k == 2 {
		twoSum(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]T, len(cur))
				copy(next, cur)
				next = append(next, nums[i])

				kSumHelper(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}
}

func twoSum[T int | int32 | int64](nums []T, left, right int, target T, cur []T, res *[][]T) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]T, len(cur))
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

func Four_sum() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	quadruplets := kSum(nums, target)
	for _, quadruplet := range quadruplets {
		fmt.Println(quadruplet)
	}
}
