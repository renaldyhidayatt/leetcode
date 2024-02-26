package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	for _, v := range candidates {
		if anagram(subject, v) {
			anagrams = append(anagrams, v)
		}

	}

	return anagrams
}

func anagram(a_word, word string) bool {
	if len(a_word) != len(word) || identical(a_word, word) == 0 {
		return false
	}

	noramalize := func(a string) []string {
		x := strings.Split(strings.ToLower(a), "")

		sort.Strings(x)

		return x
	}

	a := noramalize(a_word)
	b := noramalize(word)

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func identical(a_word, word string) int {
	return strings.Compare(strings.ToLower(a_word), strings.ToLower(word))
}
