package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestAltitude(gain []int) int {
	result := 0
	sum := 0
	for _, n := range gain {
		sum += n
		result = max(result, sum)
	}
	return result
}
