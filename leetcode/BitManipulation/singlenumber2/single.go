package singlenumber2

func singleNumberII(nums []int) int {
	ones, twos := 0, 0

	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
