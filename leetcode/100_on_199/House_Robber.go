package on199

func Rob198_1(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	curMax, preMax := 0, 0

	for i := 0; i < n; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)

		preMax = tmp
	}

	return curMax
}
