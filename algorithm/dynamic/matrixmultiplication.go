package dynamic

import "algoritmAndDs/math/min"

func MatrixChainRec(D []int, i, j int) int {
	if i == j {
		return 0
	}

	q := 1 << 32

	for k := 1; k < j; k++ {
		prod := MatrixChainRec(D, i, k) + MatrixChainRec(D, k+1, j) + D[i-1]*D[k]*D[j]
		q = min.Int[int](prod, q)
	}

	return q
}

func MatrixChainDp(D []int) int {
	N := len(D)

	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 0
	}

	for l := 2; l < N; l++ {
		for i := 1; i < N-l+1; i++ {
			j := i + l - 1
			dp[i][j] = 1 << 31

			for k := i; k < j; k++ {
				prod := dp[i][k] + dp[k+1][j] + D[i-1]*D[k]*D[j]

				dp[i][j] = min.Int(prod, dp[i][j])
			}
		}
	}

	return dp[1][N-1]
}
