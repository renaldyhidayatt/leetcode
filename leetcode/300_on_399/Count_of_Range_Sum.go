package on399

func CountRangeSum(nums []int, lower int, upper int) int {
	res, n := 0, len(nums)

	for i := 0; i < n; i++ {
		tmp := 0

		for j := i; j < n; j++ {
			if i == j {
				tmp = nums[i]
			} else {
				tmp += nums[j]
			}

			if tmp <= upper && tmp >= lower {
				res++
			}
		}
	}

	return res
}
