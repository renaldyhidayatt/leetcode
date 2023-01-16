package main

import "math"

func maxDistance(nums1, nums2 []int) int {
	return maxDistance2(nums1, nums2)
	return maxDistance1(nums1, nums2)
}

func binarySearch(nums []int, start, target int) int {
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end
}

func maxDistance1(nums1, nums2 []int) int {
	maxDist := 0
	for i := 0; i < len(nums1); i++ {
		j := binarySearch(nums2, i, nums1[i])
		maxDist = int(math.Max(float64(maxDist), float64(j-i)))
	}
	return maxDist
}

func maxDistance2(nums1, nums2 []int) int {
	i, j, dist := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			i++
		} else {
			dist = int(math.Max(float64(dist), float64(j-i)))
			j++
		}
	}
	return dist
}
