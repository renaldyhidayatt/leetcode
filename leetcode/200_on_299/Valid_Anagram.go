package on299

func IsAnagram1(s string, t string) bool {
	hash := map[rune]int{}

	for _, value := range s {
		hash[value]++
	}

	for _, value := range t {
		hash[value]--
	}

	for _, value := range hash {
		if value != 0 {
			return false
		}
	}

	return true
}
