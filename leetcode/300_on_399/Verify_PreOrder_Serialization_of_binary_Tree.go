package on399

import "strings"

func IsValidSerialization(preoder string) bool {
	nodes, diff := strings.Split(preoder, ","), 1

	for _, node := range nodes {
		diff--

		if diff < 0 {
			return false
		}

		if node != "#" {
			diff += 2
		}
	}

	return diff == 0
}
