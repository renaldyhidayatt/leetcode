package main

import "unicode"

func isPalingdrome(s string) bool {
	alnums := []rune{}

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			alnums = append(alnums, unicode.ToLower(r))
		}
	}

	ls := len(alnums)
	if ls <= 1 {
		return true
	}
	mid := ls / 2
	for i := 0; i < mid; i++ {
		if alnums[i] != alnums[ls-1-i] {
			return false
		}
	}

	return true
}
