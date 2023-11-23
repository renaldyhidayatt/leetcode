package binary

func XorSearchMissingNumber(a []int) int {
	n := len(a)
	result := len(a)
	for i := 0; i < n; i++ {
		result ^= i ^ a[i]
	}
	return result
}
