package containerwithmostwater

import "fmt"

func MaxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RunningMaxArea() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(MaxArea(height))
}
