package on399

import "sort"

func LengthOfLIS(nums []int) int {
	dp := []int{}

	for _, num := range nums {
		i := sort.SearchInts(dp, num)

		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}

	return len(dp)
}
