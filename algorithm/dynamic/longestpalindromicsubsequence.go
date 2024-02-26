package dynamic

func lpcRec(word string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}

	if word[i] == word[j] {
		return 2 + lpcRec(word, i+1, j-1)
	}

	return Max(lpcRec(word, i+1, j), lpcRec(word, i, j-1))
}

func LpsRec(word string) int {
	return lpcRec(word, 0, len(word)-1)
}

func LongestPalindromicSubsequence(word string) int {
	N := len(word)
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}

	for l := 2; l <= N; l++ {
		for i := 0; i < N-l+1; i++ {
			j := i + l - 1
			if word[i] == word[j] {
				if l == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = 2 + dp[i+1][j-1]
				}
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][N-1]
}
