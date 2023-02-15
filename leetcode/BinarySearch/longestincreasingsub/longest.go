package longestincreasingsub

func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0

	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := i; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])

			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
