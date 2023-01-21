package twosum

import "fmt"

func TwoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := myMap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		myMap[nums[i]] = i
	}

	result := []int{-1, -1}
	return result
}

func RunningTwoSum() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(TwoSum(nums, 9))
}
