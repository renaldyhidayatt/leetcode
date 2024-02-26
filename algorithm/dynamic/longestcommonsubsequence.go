package dynamic

func strToRuneSlice(s string) (r []rune, size int) {
	r = []rune(s)

	return r, len(r)
}

func LonggestCommonSubsequence(a string, b string) int {
	aRunes, aLen := strToRuneSlice(a)
	bRunes, bLen := strToRuneSlice(b)

	lcs := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lcs[i] = make([]int, bLen+1)
	}

	for i := 0; i <= aLen; i++ {
		for j := 0; j <= bLen; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if aRunes[i-1] == bRunes[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = Max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[aLen][bLen]
}
