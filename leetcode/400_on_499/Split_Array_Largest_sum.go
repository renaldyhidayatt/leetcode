package on499

func SplitArray(nums []int, m int) int {
	maxNum, sum := 0, 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if m == 1 {
		return sum
	}
	low, high := maxNum, sum
	for low < high {
		mid := low + (high-low)>>1
		if calSum(mid, m, nums) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func calSum(mid, m int, nums []int) bool {
	sum, count := 0, 0

	for _, v := range nums {
		sum += v

		if sum > mid {
			sum = v
			count++

			if count > m-1 {
				return false
			}
		}
	}

	return true
}
