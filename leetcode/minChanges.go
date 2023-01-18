package main

func addOneIfGreaterThan(cond bool) int {
	if cond {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minChanges(nums []int, k int) int {
	maxValue := 1024
	n := len(nums)
	freq := make([][]int, k)
	for i := 0; i < k; i++ {
		freq[i] = make([]int, maxValue)
	}
	for i := 0; i < n; i++ {
		freq[i%k][nums[i]]++
	}
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, maxValue)
	}
	for i := 0; i < k; i++ {
		for j := 0; j < maxValue; j++ {
			dp[i][j] = n + 1
		}
	}
	minChanges := n + 1
	for v := 0; v < maxValue; v++ {
		cntOfPos := n/k + addOneIfGreaterThan(n%k > v)

		dp[0][v] = cntOfPos - freq[0][v]
		minChanges = min(minChanges, dp[0][v])
	}
	for i := 1; i < k; i++ {
		cntOfPos := n/k + addOneIfGreaterThan(n%k > i)

		m := n + 1
		for v := 0; v < maxValue; v++ {
			for j := i; j < n; j += k {
				x := v ^ nums[j]
				dp[i][v] = min(dp[i][v], dp[i-1][x]+cntOfPos-freq[i][nums[j]])
			}
			dp[i][v] = min(dp[i][v], minChanges+cntOfPos)
			m = min(m, dp[i][v])
		}
		minChanges = m
	}
	return dp[k-1][0]
}
