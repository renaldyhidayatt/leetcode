package on299

func ContainsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))

	for _, n := range nums {
		if _, found := record[n]; found {
			return true
		}

		record[n] = true
	}

	return false
}
