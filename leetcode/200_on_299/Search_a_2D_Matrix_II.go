package on299

func SearchMatrix240(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, row := range matrix {
		low, high := 0, len(matrix[0])-1
		for low <= high {
			mid := low + (high-low)>>1
			if row[mid] > target {
				high = mid - 1
			} else if row[mid] < target {
				low = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}
