package totalhammingdistance

func totalHammingDistance(nums []int) int {
	total, n := 0, len(nums)
	for i := 0; i < 32; i++ {
		bitCount := 0
		for j := 0; j < n; j++ {
			bitCount += (nums[j] >> uint(i)) & 1
		}
		total += bitCount * (n - bitCount)
	}
	return total
}
