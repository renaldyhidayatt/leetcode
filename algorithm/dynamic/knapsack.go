package dynamic

import "math"

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Knapsack(maxWeight int, weights, values []int) int {
	n := len(weights)
	m := maxWeight

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < len(weights); i++ {
		for j := 0; j <= maxWeight; j++ {
			if weights[i] > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = Max(dp[i][j-weights[i]]+values[i], dp[i][j])
			}
		}
	}

	return dp[n][m]
}
