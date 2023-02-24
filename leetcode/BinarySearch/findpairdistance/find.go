package findpairdistance

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	low, high := 0, nums[len(nums)-1]-nums[0]

	for low < high {
		mid := low + (high-low)>>1
		tmp := findDistanceCount(nums, mid)

		if tmp >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func findDistanceCount(nums []int, num int) int {
	count, i := 0, 0
	for j := 1; j < len(nums); j++ {
		for nums[j]-nums[i] > num && i < j {
			i++
		}

		count += (j - i)
	}
	return count
}
