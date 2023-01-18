package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalingdrome(word1, word2 string) int {
	s := word1 + word2

	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	result := 0
	for start := n - 1; start >= 0; start-- {
		for end := start; end < n; end++ {
			if start == end {
				dp[start][end] = 1
				continue
			}
			if s[start] == s[end] {
				dp[start][end] = dp[start+1][end-1] + 2
				if start < len(word1) && end >= len(word1) {
					result = max(result, dp[start][end])
				}
			} else {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			}
		}
	}
	return result

}
