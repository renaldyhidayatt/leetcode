package main

func strStr(haystack, needle string) int {
	lsh, lsn := len(haystack), len(needle)
	if lsn == 0 {
		return 0
	}
	next := makeNext(needle)
	i, j := 0, 0
	for i < lsh {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
			if j == lsn {
				return i - lsn
			}
		}
		if i < lsh && haystack[i] != needle[j] {
			j = next[j]
		}
	}
	return -1
}

func makeNext(needle string) []int {
	ls := len(needle)
	next := make([]int, ls)
	next[0] = -1
	var i, j = 0, -1
	for i < ls-1 {
		if j == -1 || needle[i] == needle[j] {
			next[i+1] = j + 1
			if needle[i+1] == needle[j+1] {
				next[i+1] = next[j+1]
			}
			i++
			j++
		}
		if needle[i] != needle[j] {
			j = next[j]
		}
	}
	return next

}
