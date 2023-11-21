package on199

func MinimumTotal(triangel [][]int) int {
	if triangel == nil {
		return 0
	}

	for row := len(triangel) - 2; row >= 0; row-- {
		for col := 0; col < len(triangel[row]); col++ {
			triangel[row][col] += min(triangel[row+1][col], triangel[row+1][col+1])
		}
	}

	return triangel[0][0]
}
