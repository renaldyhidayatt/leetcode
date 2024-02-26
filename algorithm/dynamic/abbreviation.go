package dynamic

import "strings"

func Abbreviation(a string, b string) bool {
	dp := make([][]bool, len(a)+1)

	for i := range dp {
		dp[i] = make([]bool, len(b)+1)
	}

	dp[0][0] = true

	for i := 0; i < len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if dp[i][j] {
				if j < len(b) && strings.ToUpper(string(a[i])) == string(b[j]) {
					dp[i+1][j+1] = true
				}
				if string(a[i]) == strings.ToLower(string(a[i])) {
					dp[i+1][j] = true
				}
			}
		}
	}

	return dp[len(a)][len(b)]
}
