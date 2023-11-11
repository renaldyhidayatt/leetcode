package on99

import "sort"

func combinationSum2[T int | int32 | int64](candidates []T, target T) [][]T {
	if len(candidates) == 0 {
		return [][]T{}
	}

	c, res := []T{}, [][]T{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	findCombinationSum2(candidates, target, 0, c, &res)

	return res
}

func findCombinationSum2[T int | int32 | int64](nums []T, target T, index int, c []T, res *[][]T) {
	if target == 0 {
		b := make([]T, len(c))
		copy(b, c)

		*res = append(*res, b)
		return
	}

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		if target >= nums[i] {
			c = append(c, nums[i])
			findCombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
