package main

func LongestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := ExpandAroundCenter(s, i, i)
		len2 := ExpandAroundCenter(s, i, i+1)
		len := Maxs(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func ExpandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func Maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}