package dynamic

import "algoritmAndDs/math/min"

func EditDistanceRecursive(first string, second string, pointerFirst int, pointerSecond int) int {
	if pointerFirst == 0 {
		return pointerSecond
	}

	if pointerSecond == 0 {
		return pointerFirst
	}

	if first[pointerFirst-1] == second[pointerSecond-1] {
		return EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1)
	}

	return 1 + min.Int(EditDistanceRecursive(first, second, pointerFirst, pointerSecond-1),
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond),
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1))
}

func EditDistanceDP(first string, second string) int {
	m := len(first)
	n := len(second)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {

			if i == 0 {
				dp[i][j] = j
				continue
			}

			if j == 0 {
				dp[i][j] = i
				continue
			}

			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = 1 + min.Int(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
		}
	}

	return dp[m][n]
}
