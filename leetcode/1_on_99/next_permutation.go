package on99

func NextPermutation(nums []int) {
	var n = len(nums)
	var pIndex = checkPermutationPossibility(nums)

	if pIndex == -1 {
		reversee(&nums, 0, n-1)

		return
	}

	var rp = len(nums) - 1

	for rp > 0 {
		if nums[rp] > nums[pIndex] {
			swap(&nums, pIndex, rp)
			break
		} else {
			rp--
		}
	}

	reversee(&nums, pIndex+1, n-1)
}

func reversee(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)

		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func checkPermutationPossibility(nums []int) (idx int) {
	var rp = len(nums) - 1

	for rp > 0 {
		if nums[rp-1] < nums[rp] {
			idx = rp - 1

			return idx
		}

		rp--
	}

	return -1
}
