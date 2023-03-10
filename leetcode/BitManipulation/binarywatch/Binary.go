package binarywatch

import "strconv"

func readBinaryWatch(num int) []string {
	memo := make([]int, 60)
	count := func(n int) int {
		if memo[n] != 0 {
			return memo[n]
		}
		originN, res := n, 0
		for n != 0 {
			n = n & (n - 1)
			res++
		}
		memo[originN] = res
		return res
	}
	fmtMinute := func(m int) string {
		if m < 10 {
			return "0" + strconv.Itoa(m)
		}
		return strconv.Itoa(m)
	}

	var res []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count(i)+count(j) == num {
				res = append(res, strconv.Itoa(i)+":"+fmtMinute(j))
			}
		}
	}
	return res
}
