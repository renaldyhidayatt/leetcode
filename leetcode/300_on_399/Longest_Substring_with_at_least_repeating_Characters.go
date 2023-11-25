package on399

import "strings"

func LongestSubString(s string, k int) int {
	if s == "" {
		return 0
	}

	freq, split, res := [26]int{}, byte(0), 0

	for _, ch := range s {
		freq[ch-'a']++
	}

	for i, c := range freq[:] {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		res = max(res, LongestSubString(subStr, k))
	}

	return res

}
