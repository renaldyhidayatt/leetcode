package on299

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob213(nums []int, start, end int) int {
	preMax := nums[start]
	curMax := max(preMax, nums[start+1])
	for i := start + 2; i <= end; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}
