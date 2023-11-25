package on399

import "sort"

func KSmallesPairs(nums1 []int, nums2 []int, k int) [][]int {
	size1, size2, res := len(nums1), len(nums2), [][]int{}

	if size1 == 0 || size2 == 0 || k < 0 {
		return nil
	}

	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			res = append(res, []int{nums1[i], nums2[j]})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})

	if len(res) >= k {
		return res[:k]
	}

	return res
}
