package main

func maxSubArray(nums []int) int {
	maxEndingHere := nums[0]
	maxSofFar := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxSofFar = max(maxEndingHere, maxSofFar)
	}
	return maxSofFar
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
