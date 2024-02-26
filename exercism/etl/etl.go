package etl

import "strings"

func Transform(lps map[int][]string) map[string]int {
	spl := map[string]int{}

	for key, value := range lps {
		for _, v := range value {
			spl[strings.ToLower(v)] = key
		}
	}

	return spl
}
