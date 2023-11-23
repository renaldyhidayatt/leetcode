package on199

func FindPeakElement(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)>>1

		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
