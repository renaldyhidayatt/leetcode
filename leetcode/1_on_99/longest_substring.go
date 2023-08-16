package on99

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitset [256]bool
	result, left, right := 0, 0, 0

	for left < len(s) {
		if bitset[s[right]] {
			bitset[s[left]] = false
			left++
		} else {
			bitset[s[right]] = true
			right++
		}

		if result < right-left {
			result = right - left
		}

		if left+result >= len(s) || right >= len(s) {
			break
		}

	}

	return result
}
