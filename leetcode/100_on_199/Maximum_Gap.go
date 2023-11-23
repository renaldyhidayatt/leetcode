package on199

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	quickSort164(nums, 0, len(nums)-1)

	res := 0

	for i := 0; i < len(nums)-1; i++ {
		if (nums[i+1] - nums[i]) > res {
			res = nums[i+1] - nums[i]
		}
	}

	return res
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]

	i := lo - 1

	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}

	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)

}
