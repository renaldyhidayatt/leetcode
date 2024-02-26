package strings

func CountChars(text string) map[rune]int {
	charMap := make(map[rune]int, 0)

	for _, c := range text {
		if _, ok := charMap[c]; !ok {
			charMap[c] = 0
		}

		charMap[c]++
	}

	return charMap
}
