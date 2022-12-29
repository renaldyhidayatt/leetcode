package main

func permute(nums []int) [][]int {
	// DPS with swapping
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	getPermute(&res, nums, 0)
	return res
}

func getPermute(res *[][]int, nums []int, index int) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		// s(n) = 1 + s(n-1)
		getPermute(res, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
