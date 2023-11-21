package on199

func NumDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)

	for i, curT := range t {
		pre := 0

		for j, curS := range s {
			if i == 0 {
				pre = 1
			}
			newDP := dp[j+1]
			if curT == curS {
				dp[j+1] = dp[j] + pre
			} else {
				dp[j+1] = dp[j]
			}

			pre = newDP
		}
	}

	return dp[len(s)]
}
