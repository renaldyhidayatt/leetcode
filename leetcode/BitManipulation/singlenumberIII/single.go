package singlenumberiii

func singleNumberIII(nums []int) []int {
	diff := 0

	for _, num := range nums {
		diff ^= num
	}

	diff &= -diff
	res := []int{0, 0}
	for _, num := range nums {
		if (num & diff) == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}
