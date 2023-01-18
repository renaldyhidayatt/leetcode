package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLengthBetweenEqualCharacters(s string) int {
	pos := make([]int, 26)

	longest := -1
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if pos[idx] == 0 {
			pos[idx] = i + 1
		} else {
			longest = max(longest, i-pos[idx])
		}
	}
	return longest

}
