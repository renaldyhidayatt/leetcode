package main

func threeConsecutiveOdds(arr []int) bool {
	cnt := 0
	for _, n := range arr {
		if n%2 == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt >= 3 {
			return true
		}
	}
	return false
}
