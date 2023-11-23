package on199

func Rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))

	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}

	copy(nums, newNums)
}
