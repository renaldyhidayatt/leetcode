package on99

import "sort"

func combinationSum[T int | int32 | int64](candidates []T, target T) [][]T {
	if len(candidates) == 0 {
		return [][]T{}
	}

	c, res := []T{}, [][]T{}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum[T int | int32 | int64](nums []T, target T, index int, c []T, res *[][]T) {
	if target <= 0 {
		if target == 0 {
			b := make([]T, len(c))
			copy(b, c)

			*res = append(*res, b)
		}

		return
	}

	if nums[index] > target {
		return
	}

	for i := index; i < len(nums); i++ {

		if target-nums[i] <= 0 {
			break
		}

		c = append(c, nums[i])

		findCombinationSum(nums, target-nums[i], i, c, res)

		c = c[:len(c)-1]
	}
}
