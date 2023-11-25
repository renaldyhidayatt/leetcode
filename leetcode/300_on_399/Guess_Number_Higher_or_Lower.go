package on399

import "sort"

func GuessNumber(n int) int {
	return sort.Search(n, func(i int) bool {
		return guess(i) <= 0
	})
}

func guess(num int) int {
	return 0
}
