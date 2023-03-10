package subsets

import "sort"

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)

	for i := range nums {
		for _, org := range res {
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}
