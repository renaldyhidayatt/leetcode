package on399

func KthSmalesst378(matrix [][]int, k int) int {
	m, n, low := len(matrix), len(matrix[0]), matrix[0][0]

	high := matrix[m-1][n-1] + 1

	for low < high {
		mid := low + (high-low)>>1

		if counterKthSmall(m, n, mid, matrix) >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func counterKthSmall(m, n, mid int, matrix [][]int) int {
	count, j := 0, n-1

	for i := 0; i < m; i++ {
		for j >= 0 && mid < matrix[i][j] {
			j--
		}

		count += j + 1
	}

	return count
}
