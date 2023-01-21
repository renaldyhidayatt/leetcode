package medianoftwosortedarray

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2, pos := 0, 0, 0
	ls1, ls2 := len(nums1), len(nums2)
	allNums := make([]int, ls1+ls2)
	var median float64
	for p1 < ls1 && p2 < ls2 {
		if nums1[p1] <= nums2[p2] {
			allNums[pos] = nums1[p1]
			p1++
		} else {
			allNums[pos] = nums2[p2]
			p2++
		}
		pos++
	}
	for p1 < ls1 {
		allNums[pos] = nums1[p1]
		p1++
		pos++
	}
	for p2 < ls2 {
		allNums[pos] = nums2[p2]
		p2++
		pos++
	}
	if (ls1+ls2)%2 == 1 {
		median = float64(allNums[(ls1+ls2)/2])
	} else {
		median = float64(allNums[(ls1+ls2)/2]+allNums[(ls1+ls2)/2-1]) / 2.0
	}
	return median
}
